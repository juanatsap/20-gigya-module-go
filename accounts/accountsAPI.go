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
func (a *AccountsAPI) Search(query string, limit int) (Accounts, int, error) {

	if limit < 1 {
		limit = 1
	}

	if limit > 0 {
		query = fmt.Sprintf("%s limit %d", query, limit)
	}

	// Añadir parámetros
	method := "accounts.search"
	params := map[string]string{
		"query":   query,
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
	total := 0

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, total, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, total, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response SearchResponse
	err = json.Unmarshal(body, &response)
	bodyAsString := string(body)
	if err != nil {
		log.Errorf("Error parsing JSON for body %v", bodyAsString)
		return nil, total, err
	}

	// Verificar si hubo un error en la respuesta
	total = response.TotalCount

	if response.ErrorCode != 0 {
		return nil, total, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return response.Results, total, nil
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
		"profile": account.Profile.AsJSON(),
		"data":    account.Data.AsJSON(),
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
	method := "accounts.deleteAccount"
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
