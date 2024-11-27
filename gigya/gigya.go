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

func (g *Gigya) SetApiKey(apiKey string) {
	g.apiKey = apiKey
	g.AccountsAPI = accounts.NewAccountsAPI(apiKey, g.userKey, g.secretKey, g.apiDomain)
}

func (g *Gigya) SetUserKey(userKey string) {
	g.userKey = userKey
	g.AccountsAPI = accounts.NewAccountsAPI(g.apiKey, userKey, g.secretKey, g.apiDomain)
}

func (g *Gigya) SetSecretKey(secretKey string) {
	g.secretKey = secretKey
	g.AccountsAPI = accounts.NewAccountsAPI(g.apiKey, g.userKey, secretKey, g.apiDomain)
}

func (g *Gigya) SetApiDomain(apiDomain string) {
	g.apiDomain = apiDomain
	g.AccountsAPI = accounts.NewAccountsAPI(g.apiKey, g.userKey, g.secretKey, apiDomain)
}
