package accounts

import (
	"encoding/json"
	"fmt"

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
func (a *AccountsAPI) Search(query string, limit int) (Accounts, error) {

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

	// Enviar la solicitud POST
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Deserializar la respuesta JSON en SearchResponse
	var response SearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Verificar si hubo un error en la respuesta
	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("API error %d: %s", response.ErrorCode, response.StatusReason)
	}

	return response.Results, nil
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
