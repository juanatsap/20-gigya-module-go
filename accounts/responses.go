package accounts

type GetAccountInfoResponse struct {
	CallID               string                  `json:"callId"`
	ErrorCode            int                     `json:"errorCode"`
	APIVersion           int                     `json:"apiVersion"`
	StatusCode           int                     `json:"statusCode"`
	StatusReason         string                  `json:"statusReason"`
	Time                 string                  `json:"time"`
	UID                  string                  `json:"UID"`
	Profile              Profile                 `json:"profile,omitempty"`
	Data                 Data                    `json:"data,omitempty"`
	Preferences          Preferences             `json:"preferences,omitempty"`
	Subscriptions        map[string]Subscription `json:"subscriptions,omitempty"`
	Created              string                  `json:"created,omitempty"`
	CreatedTimestamp     int64                   `json:"createdTimestamp,omitempty"`
	LastUpdated          string                  `json:"lastUpdated,omitempty"`
	LastUpdatedTimestamp int64                   `json:"lastUpdatedTimestamp,omitempty"`
	HasLiteAccount       bool                    `json:"hasLiteAccount,omitempty"`
	HasFullAccount       bool                    `json:"hasFullAccount,omitempty"`
	Emails               Emails                  `json:"emails,omitempty"`
	LoginIDs             LoginIDs                `json:"loginIDs,omitempty"`
	IsVerified           bool                    `json:"isVerified,omitempty"`
	IsRegistered         bool                    `json:"isRegistered,omitempty"`
	Password             Password                `json:"password,omitempty"`
	RegSource            string                  `json:"regSource,omitempty"`
	Error                string                  `json:"error,omitempty"`
	IsActive             bool                    `json:"isActive,omitempty"`
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
type SearchResponse struct {
	CallID       string   `json:"callId"`
	ErrorCode    int      `json:"errorCode"`
	ErrorDetails string   `json:"errorDetails"`
	APIVersion   int      `json:"apiVersion"`
	StatusCode   int      `json:"statusCode"`
	StatusReason string   `json:"statusReason"`
	Time         string   `json:"time"`
	Results      Accounts `json:"results"`
	ObjectsCount int      `json:"objectsCount"`
	TotalCount   int      `json:"totalCount"`
	NextCursor   string   `json:"nextCursor,omitempty"` // Cursor for paginated results
}
type ImportFullAccountResponse struct {
	CallID       string `json:"callId"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetails string `json:"errorDetails"`
	APIVersion   int    `json:"apiVersion"`
	StatusCode   int    `json:"statusCode"`
	StatusReason string `json:"statusReason"`
	Time         string `json:"time"`
	UID          string `json:"UID"`
}
type DeleteAccountResponse struct {
	CallID       string `json:"callId"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetails string `json:"errorDetails"`
	APIVersion   int    `json:"apiVersion"`
	StatusCode   int    `json:"statusCode"`
	StatusReason string `json:"statusReason"`
	Time         string `json:"time"`
	UID          string `json:"UID"`
}

type GetJWTPublicKeyResponse struct {
	CallID       string `json:"callId"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetails string `json:"errorDetails"`
	APIVersion   int    `json:"apiVersion"`
	StatusCode   int    `json:"statusCode"`
	StatusReason string `json:"statusReason"`
	Time         string `json:"time"`
	Alg          string `json:"alg"`
	Kty          string `json:"kty"`
	N            string `json:"n"`
	E            string `json:"e"`
	Use          string `json:"use"`
	Kid          string `json:"kid"`
}
