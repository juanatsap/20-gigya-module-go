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
type LoginClaims struct {
	ApiKey         string `json:"apiKey"`
	CallID         string `json:"callID"`
	ExtensionPoint string `json:"extensionPoint"`
	Data           Data   `json:"data"`
}

type Data struct {
	Params      Params           `json:"params"`
	AccountInfo accounts.Account `json:"accountInfo"`
	Context     Context          `json:"context"`
}

type Params struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
}

type Context struct {
	ClientIP string `json:"clientIP"`
}
