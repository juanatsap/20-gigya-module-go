package accounts

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type AccountsAPI struct {
	apiKey    string
	userKey   string
	secretKey string
	apiDomain string
}

func NewAccountsAPI(apiKey, userKey, secretKey, apiDomain string) *AccountsAPI {

	return &AccountsAPI{
		apiKey:    apiKey,
		userKey:   userKey,
		secretKey: secretKey,
		apiDomain: apiDomain,
	}
}

/* ╭──────────────────────────────────────────╮ */
/* │            ACCOUNT API CALLS             │ */
/* ╰──────────────────────────────────────────╯ */

// SearchWithCursor performs a search with pagination support using cursor
// Parameters:
// - query: The search query to execute
// - limit: The maximum number of results to return per request (max 100 recommended)
// - cursor: The cursor string from a previous search result, empty string for first page
// Returns:
// - accounts: The list of accounts matching the query
// - totalCount: The total number of accounts matching the query
// - nextCursor: Cursor string for retrieving the next batch, empty if no more results
// - error: Any error that occurred during the search
func (a *AccountsAPI) SearchWithCursor(query string, limit int, cursor string) (Accounts, int, string, error) {
	if limit < 1 {
		limit = 1
	}

	// // Limit should not exceed 100 for reliable results with cursor pagination
	// if limit > 100 {
	// 	limit = 100
	// }

	// Add the limit to the query if this is the first call (no cursor provided)
	if cursor == "" && limit > 0 {
		// Only add limit to the query for the first call
		query = fmt.Sprintf("%s limit %d", query, limit)
	}

	// Prepare the API request parameters
	method := "accounts.search"
	params := map[string]string{
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
	}

	// If this is the first call, use query and openCursor=true
	if cursor == "" {
		params["query"] = query
		params["openCursor"] = "true"
	} else {
		// For subsequent calls, use the cursor ID from the previous call
		params["cursorId"] = cursor
	}

	// Build the request URL
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Send the POST request
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, 0, "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, "", err
	}

	// Parse the JSON response
	var response SearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Error parsing JSON: %v", err)
		
		// Try to identify the problematic record
		var debugResponse map[string]interface{}
		if debugErr := json.Unmarshal(body, &debugResponse); debugErr == nil {
			if results, ok := debugResponse["results"].([]interface{}); ok {
				log.Errorf("Found %d results in response", len(results))
				
				// Check each result for markedForDeletion field
				for i, result := range results {
					if account, ok := result.(map[string]interface{}); ok {
						uid, _ := account["UID"].(string)
						
						// Check for the data.account.markedForDeletion field
						if data, ok := account["data"].(map[string]interface{}); ok {
							if accountData, ok := data["account"].(map[string]interface{}); ok {
								if markedForDeletion, exists := accountData["markedForDeletion"]; exists {
									markedType := fmt.Sprintf("%T", markedForDeletion)
									log.Errorf("Record %d (UID: %s) has markedForDeletion of type %s with value: %v", 
										i, uid, markedType, markedForDeletion)
									
									// Log the entire account data for this problematic record
									if accountJSON, _ := json.MarshalIndent(account, "", "  "); len(accountJSON) > 0 {
										log.Errorf("Full record data:\n%s", string(accountJSON))
									}
								}
							}
						}
					}
				}
			}
		}
		
		return nil, 0, "", err
	}

	// Check for API errors
	if response.ErrorCode != 0 {
		return nil, 0, "", fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return response.Results, response.TotalCount, response.NextCursor, nil
}

// SearchAll retrieves all accounts matching the specified query by making multiple paginated requests
// This method will automatically handle the pagination using cursors
// Parameters:
// - query: The search query to execute
// - batchSize: The number of records to retrieve per request (max 100 recommended)
// - progressCallback: Optional callback function to report progress (can be nil)
// Returns:
// - accounts: All accounts matching the query
// - totalCount: The total number of accounts matching the query
// - error: Any error that occurred during the search
func (a *AccountsAPI) SearchAll(query string, batchSize int, progressCallback func(fetched, total int)) (Accounts, int, error) {
	if batchSize < 1 {
		batchSize = 100 // Default batch size
	}

	// Cap the batch size to a reasonable value
	if batchSize > 100 {
		batchSize = 100
	}

	// Initial search to get the first batch and total count
	accounts, totalCount, nextCursor, err := a.SearchWithCursor(query, batchSize, "")
	if err != nil {
		return nil, 0, err
	}

	// Call the progress callback if provided
	if progressCallback != nil {
		progressCallback(len(accounts), totalCount)
	}

	// If there are more results, continue fetching
	for nextCursor != "" {
		// Fetch the next batch
		nextBatch, _, nextCursorVal, err := a.SearchWithCursor(query, batchSize, nextCursor)
		if err != nil {
			// Return what we've got so far along with the error
			return accounts, totalCount, fmt.Errorf("error fetching batch with cursor %s: %w", nextCursor, err)
		}

		// Append the results
		accounts = append(accounts, nextBatch...)

		// Update the cursor for the next iteration
		nextCursor = nextCursorVal

		// Call the progress callback if provided
		if progressCallback != nil {
			progressCallback(len(accounts), totalCount)
		}
	}

	return accounts, totalCount, nil
}

// Search maintains backward compatibility with existing code
// Deprecated: Use SearchWithCursor or SearchAll instead
func (a *AccountsAPI) Search(query string, limit int) (Accounts, int, error) {
	accounts, totalCount, _, err := a.SearchWithCursor(query, limit, "")
	return accounts, totalCount, err
}
func (a *AccountsAPI) GetAccountInfo(UID string) (Account, error) {

	// Añadir parámetros
	method := "accounts.getAccountInfo"
	params := map[string]string{
		"UID":     UID,
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
		"include": "profile, data, preferences, emails, loginIDs",
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response GetAccountInfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Error parsing JSON for UID %s: %v", UID, err)
		return Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return Account{}, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	// Create a new Account object
	account := Account{
		UID:     response.UID,
		Profile: response.Profile,
		Data:    response.Data,
		Preferences: Preferences{
			Marketing: Marketing{
				Email: ConsentDetail{
					IsConsentGranted: response.Preferences.Marketing.Email.IsConsentGranted,
				},
			},
			Terms: Terms{
				ToS: ConsentDetail{
					IsConsentGranted: response.Preferences.Terms.ToS.IsConsentGranted,
				},
			},
			Privacy: Privacy{
				Livgolf: ConsentDetail{
					IsConsentGranted: response.Preferences.Privacy.Livgolf.IsConsentGranted,
				},
			},
		},
		Created: response.Created,
		Emails:  response.Emails,
		LoginIDs: LoginIDs{
			Emails: response.LoginIDs.Emails,
		},
		IsVerified:     response.IsVerified,
		IsRegistered:   response.IsRegistered,
		Password:       response.Password,
		RegSource:      response.RegSource,
		HasLiteAccount: response.HasLiteAccount,
		HasFullAccount: response.HasFullAccount,
		IsActive:       response.IsActive,
	}

	return account, nil
}
func (a *AccountsAPI) SetAccountInfo(account Account, isLite bool) (Account, error) {

	// Añadir parámetros
	method := "accounts.setAccountInfo"
	params := map[string]string{
		"UID":     account.UID,
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
		// "profile": account.Profile.AsJSON(),
		"data": account.Data.AsJSON(),
	}

	// Add isLite: true if isLite is true
	if isLite {
		params["isLite"] = "true"
	}

	if account.Data.Competition != nil {
		// Check if both competition name and competition when are empty. if yes, set this fixed data section: {data: {competition: null}}
		if account.Data.Competition.Name == "" && account.Data.Competition.When == "" {
			params["data"] = `{"competition":null}`
		}
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response SetAccountInfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return Account{}, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return Account{UID: account.UID}, nil
}
func (a *AccountsAPI) ImportFullAccount(account Account) (Account, error) {
	// Añadir parámetros
	method := "accounts.importFullAccount"
	params := map[string]string{
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
		// "importPolicy": "upsert",
		"importPolicy": "insert",
		"account":      account.AsJSON(),
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response ImportFullAccountResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return Account{}, fmt.Errorf("API error %d: %s, %s\n\nDetails: %s", response.ErrorCode, response.StatusReason, response.ErrorMessage, response.ErrorDetails)
	}

	return Account{UID: response.UID}, nil
}
func (a *AccountsAPI) DeleteAccount(UID string) (Account, error) {

	// Añadir parámetros
	method := "accounts.setAccountInfo"
	params := map[string]string{
		"UID":     UID,
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response DeleteAccountResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return Account{}, fmt.Errorf("API error %d: %s, %s\n\nDetails: %s", response.ErrorCode, response.StatusReason, response.ErrorMessage, response.ErrorDetails)
	}

	return Account{UID: response.UID}, nil
}

/* ╭──────────────────────────────────────────╮ */
/* │         IDXIMPORT ID  API CALLS          │ */
/* ╰──────────────────────────────────────────╯ */
func (a *AccountsAPI) SearchAccountsForIdxImportId(idxImportId string) ([]Account, error) {

	query := fmt.Sprintf("Select * from accounts where idxImportId='%s'", idxImportId)
	// Añadir parámetros
	method := "accounts.search"
	params := map[string]string{
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
		"query":   query,
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return []Account{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Account{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response SearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return []Account{}, fmt.Errorf("API error %d: %s\n\nDetails: %s", response.ErrorCode, response.StatusReason, response.ErrorDetails)
	}

	return response.Results, nil
}
func (a *AccountsAPI) DeleteAccountsForIdxImportId(idxImportId string) ([]Account, error) {
	accounts, err := a.SearchAccountsForIdxImportId(idxImportId)
	if err != nil {
		return []Account{}, err
	}
	if len(accounts) == 0 {
		// ui.CreateBox(fmt.Sprintf("No accounts found for idxImportId: %s", idxImportId), "Highly Customized Terminal Box Maker", 0)
		log.Debugf("No accounts found for idxImportId: %s\n", idxImportId)
		return []Account{}, nil
	}
	for i, account := range accounts {
		_, err := a.DeleteAccount(account.UID)

		if err != nil {
			log.Errorf("i: %d - Error deleting account: %s\n", i, err)
			return []Account{}, err
		} else {
			log.Printf("i: %d - Account deleted: %s\n", i, account.Profile.Email)
		}
	}
	return accounts, nil
}
func (a *AccountsAPI) GetJWTPublicKey() (GetJWTPublicKeyResponse, error) {
	// Añadir parámetros
	method := "accounts.getJWTPublicKey"
	params := map[string]string{
		"apiKey": a.apiKey,
	}

	// Preparar la URL de la solicitud
	baseURL := fmt.Sprintf("https://%s/%s", a.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return GetJWTPublicKeyResponse{}, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetJWTPublicKeyResponse{}, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response GetJWTPublicKeyResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return GetJWTPublicKeyResponse{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return GetJWTPublicKeyResponse{}, fmt.Errorf("API error %d: %s\n\nDetails: %s", response.ErrorCode, response.StatusReason, response.ErrorDetails)
	}

	return response, nil
}
