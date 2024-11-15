package gigya

import "gigya-module-go/accounts"

type Gigya struct {
	apiKey      string
	userKey     string
	secretKey   string
	apiDomain   string
	AccountsAPI *accounts.AccountsAPI
}

func NewGigya(apiKey, userKey, secretKey, apiDomain string) *Gigya {

	return &Gigya{
		apiKey:      apiKey,
		userKey:     userKey,
		secretKey:   secretKey,
		apiDomain:   apiDomain,
		AccountsAPI: accounts.NewAccountsAPI(apiKey, userKey, secretKey, apiDomain),
	}
}
