package extensions

import (
	"encoding/json"
	"gigya-module-go/accounts"

	"github.com/golang-jwt/jwt"
)

// mapClaimsToStruct convierte jwt.MapClaims a una estructura Go
func MapClaimsToStruct(claims jwt.MapClaims, target interface{}) error {
	// Marshal a JSON y unmarshal a la estructura destino
	jsonData, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, target)
}

/*
jwt.Claims(github.com/golang-jwt/jwt.MapClaims) ["apiKey": "4_fok0Wsjf2RMSy-Oksfktjw", "callID": "3aa65ea84e7c422f8238e08ad64b81e6", "extensionPoint": "OnBeforeAccountsLogin", "data": map[string]interface {} ["params": *(*interface {})(0x14000000628), "accountInfo": *(*interface {})(0x14000000638), "context": *(*interface {})(0x14000000648), ], ]
*/
type ExtensionClaims struct {
	ApiKey         string `json:"apiKey,omitempty"`
	CallID         string `json:"callID,omitempty"`
	ExtensionPoint string `json:"extensionPoint,omitempty"`
	Data           Data   `json:"data,omitempty"`
}

type Data struct {
	Params      Params           `json:"params,omitempty"`
	AccountInfo accounts.Account `json:"accountInfo,omitempty"`
	Context     Context          `json:"context,omitempty"`
	// Preferences Preferences      `json:"preferences,omitempty"`
}

type Preferences struct {
	Livx struct {
		IsConsentGranted string `json:"isConsentGranted,omitempty"`
	} `json:"livx"`
}

type Params struct {
	LoginId     string       `json:"loginId,omitempty"`
	Password    string       `json:"password,omitempty"`
	Email       string       `json:"email,omitempty"`
	Locale      string       `json:"locale,omitempty"`
	FirstName   string       `json:"firstName,omitempty"`
	LastName    string       `json:"lastName,omitempty"`
	Country     string       `json:"country,omitempty"`
	PostalCode  string       `json:"postalcode,omitempty"`
	State       string       `json:"state,omitempty"`
	StateName   string       `json:"stateName,omitempty"`
	Data        DataFromFrom `json:"data,omitempty"`
	Preferences Preferences  `json:"preferences"`
}

type DataFromFrom struct {
	Fantasy Fantasy `json:"fantasy,omitempty"`
}

type Fantasy struct {
	TeamName string `json:"teamName,omitempty"`
}

type Context struct {
	ClientIP string `json:"clientIP,omitempty"`
}
