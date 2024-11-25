package accounts

import (
	"encoding/json"
	"fmt"
	"strconv"

	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	if err != nil {
		return nil, total, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		total = response.TotalCount
		return nil, total, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return response.Results, total, nil
}
func (a *AccountsAPI) SearchGrouped(query string) (GroupedLIVGolfItems, int, error) {

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
	var response SearchGroupedResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, total, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return nil, total, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}
	total = response.TotalCount

	return response.Results, total, nil
}
func (a *AccountsAPI) SetAccountInfoLIV(account Account, isLite bool) (Account, error) {

	dataAsJSON := account.Data.AsJSON()
	if dataAsJSON == "{\"competition\":{},\"favoriteTeam\":{}}" {
		dataAsJSON = `{"competition":{"name":null,"when":null},"favoriteTeam":{}}`
	}

	// Añadir parámetros
	method := "accounts.setAccountInfo"
	params := map[string]string{
		"UID":     account.UID,
		"apiKey":  a.apiKey,
		"userKey": a.userKey,
		"secret":  a.secretKey,
		"profile": account.Profile.AsJSON(),
		"data":    dataAsJSON,
		// "data":   `{"favoriteTeam":{"name":null,"since":null}}`,
		"isLite": strconv.FormatBool(isLite),
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

	return Account{UID: response.UID}, nil
}
func (a *AccountsAPI) SetAccountInfo(account Account) (Account, error) {

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
func (a *AccountsAPI) GetAccountInfo(UID string) (Account, error) {

	// Añadir parámetros
	method := "accounts.getAccountInfo"
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
	var response GetAccountInfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Account{}, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return Account{}, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return Account{UID: response.UID}, nil
}

// SearchResponse representa la respuesta completa de accounts.search
type SearchResponse struct {
	CallID       string   `json:"callId"`
	ErrorCode    int      `json:"errorCode"`
	APIVersion   int      `json:"apiVersion"`
	StatusCode   int      `json:"statusCode"`
	StatusReason string   `json:"statusReason"`
	Time         string   `json:"time"`
	Results      Accounts `json:"results"`
	ObjectsCount int      `json:"objectsCount"`
	TotalCount   int      `json:"totalCount"`
}
type SearchGroupedResponse struct {
	CallID       string              `json:"callId"`
	ErrorCode    int                 `json:"errorCode"`
	APIVersion   int                 `json:"apiVersion"`
	StatusCode   int                 `json:"statusCode"`
	StatusReason string              `json:"statusReason"`
	Time         string              `json:"time"`
	Results      GroupedLIVGolfItems `json:"results"`
	ObjectsCount int                 `json:"objectsCount"`
	TotalCount   int                 `json:"totalCount"`
}
type GroupedLIVGolfItem struct {
	Count           int    `json:"count(*)"`
	Name            string `json:"data.favoriteTeam.name,omitempty"`
	Visited         string `json:"data.visited,omitempty"`
	CompetitionName string `json:"data.competition.name,omitempty"`
	EventsName      string `json:"data.events.name,omitempty"`
}
type GroupedVisited struct {
	Count int    `json:"count(*)"`
	Name  string `json:"data.visited"`
}
type SetAccountInfoResponse struct {
	CallID       string `json:"callId"`
	ErrorCode    int    `json:"errorCode"`
	APIVersion   int    `json:"apiVersion"`
	StatusCode   int    `json:"statusCode"`
	StatusReason string `json:"statusReason"`
	Time         string `json:"time"`
	UID          string `json:"UID"`
}
type GetAccountInfoResponse struct {
	CallID       string      `json:"callId"`
	ErrorCode    int         `json:"errorCode"`
	APIVersion   int         `json:"apiVersion"`
	StatusCode   int         `json:"statusCode"`
	StatusReason string      `json:"statusReason"`
	Time         string      `json:"time"`
	UID          string      `json:"UID"`
	Profile      Profile     `json:"profile"`
	Data         Data        `json:"data"`
	Preferences  Preferences `json:"preferences"`
}
type GroupedLIVGolfItems []GroupedLIVGolfItem
