package gigya

type Gigya struct {
	apiKey    string
	userKey   string
	secretKey string
	apiDomain string
}

func NewGigya(apiKey, userKey, secretKey, apiDomain string) *Gigya {
	return &Gigya{
		apiKey:    apiKey,
		userKey:   userKey,
		secretKey: secretKey,
		apiDomain: apiDomain,
	}
}
