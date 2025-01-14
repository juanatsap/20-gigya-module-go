package accounts

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Data representa los datos personalizados de la cuenta
type Data struct {
	IdxImportId string `json:"idxImportId,omitempty"`

	/* ╭──────────────────────────────────────────╮ */
	/* │                 LIVGOLF                  │ */
	/* ╰──────────────────────────────────────────╯ */
	// LIVXMember string `json:"LIVX_Member,omitempty"`
	Visited      string       `json:"visited,omitempty"`
	Competition  *Competition `json:"competition,omitempty"`
	FavoriteTeam *NameSince   `json:"favoriteTeam,omitempty"`
	// DataSource string `json:"dataSource,omitempty"`
	Events     *Event    `json:"events,omitempty"`
	RipperGC   *RipperGC `json:"rippergc,omitempty"`
	DataSource string    `json:"dataSource,omitempty"`

	/* ╭──────────────────────────────────────────╮ */
	/* │                 OLYMPICS                 │ */
	/* ╰──────────────────────────────────────────╯ */
	// Add data.utility.isAthlete
	Utility         *Utility         `json:"utility,omitempty"`
	Personalization *Personalization `json:"personalization,omitempty"`
}
type Personalization struct {
	SiteLanguageP24 string `json:"siteLanguageP24,omitempty"`
	SiteLanguage    string `json:"siteLanguage,omitempty"`
}
type Utility struct {
	IsAthlete bool `json:"isAthlete,omitempty"`
}
type RipperGC struct {
	Trivia *Trivia `json:"trivia,omitempty"`
}
type Trivia struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

func (a Data) AsJSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}

// NameWhen representa una estructura con nombre y fecha
type NameWhen struct {
	Name string `json:"name,omitempty"`
	When string `json:"when,omitempty"`
}
type Event NameWhen
type Competition NameWhen

// NameSince representa una estructura que puede ser un string, un array o un objeto vacío
type NameSince struct {
	Name  string `json:"name,omitempty"`
	Since string `json:"since,omitempty"`
}

// Preferences representa las preferencias de la cuenta
type Preferences struct {
	Marketing Marketing `json:"marketing,omitempty"`
	Terms     Terms     `json:"terms,omitempty"`
	Privacy   Privacy   `json:"privacy,omitempty"`
}

// Privacy representa las preferencias de la cuenta
type Privacy struct {
	Livgolf ConsentDetail `json:"livgolf,omitempty"`
	// Majesticks   ConsentDetail `json:"majesticks,omitempty"`
	// SportsBreaks ConsentDetail `json:"sportsBreaks,omitempty"`
}
type Terms struct {
	ToS ConsentDetail `json:"ToS,omitempty"`
}
type Marketing struct {
	Email        ConsentDetail `json:"email,omitempty"`
	SportsBreaks ConsentDetail `json:"sportsBreaks,omitempty"`
}

// ConsentDetail representa los detalles de consentimiento
type ConsentDetail struct {
	Entitlements        []string                `json:"entitlements,omitempty"`
	Locales             map[string]LocaleDetail `json:"locales,omitempty"`
	IsConsentGranted    bool                    `json:"isConsentGranted"`
	ActionTimestamp     string                  `json:"actionTimestamp,omitempty"`
	CustomData          []string                `json:"customData,omitempty"`
	Language            string                  `json:"language,omitempty"`
	LastConsentModified string                  `json:"lastConsentModified,omitempty"`
	DocVersion          float64                 `json:"docVersion,omitempty"`
	DocDate             string                  `json:"docDate,omitempty"`
	Tags                []string                `json:"tags,omitempty"`
}

// LocaleDetail representa los detalles específicos de una localidad
type LocaleDetail struct {
	DocVersion float64 `json:"docVersion,omitempty"`
	DocDate    string  `json:"docDate,omitempty"`
}

// Subscription representa las suscripciones de la cuenta
type Subscription struct {
	Email SubscriptionChannel `json:"email,omitempty"`
}

// SubscriptionChannel representa los detalles de una suscripción por canal
type SubscriptionChannel struct {
	IsSubscribed                 bool              `json:"isSubscribed,omitempty"`
	LastUpdatedSubscriptionState string            `json:"lastUpdatedSubscriptionState,omitempty"`
	DoubleOptIn                  DoubleOptInDetail `json:"doubleOptIn,omitempty"`
}

// DoubleOptInDetail representa los detalles de doble opt-in
type DoubleOptInDetail struct {
	Status string `json:"status,omitempty"`
}

type GroupedLIVGolfItems []GroupedLIVGolfItem

func (a GroupedLIVGolfItems) ToCSV() string {

	var b bytes.Buffer
	w := csv.NewWriter(&b)
	w.Write([]string{"Count", "Name"})
	for _, item := range a {

		var row []string
		name := item.Name

		if name == "" {
			name = item.DataIdxImportId
		}

		row = append(row, strconv.Itoa(item.Count))
		row = append(row, name)
		w.Write(row)
	}

	w.Flush()
	return b.String()
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
	DataIdxImportId string `json:"data.idxImportId,omitempty"`
}
type GroupedVisited struct {
	Count int    `json:"count(*)"`
	Name  string `json:"data.visited"`
}

/* ╭──────────────────────────────────────────╮ */
/* │               ARRAYS SHIT                │ */
/* ╰──────────────────────────────────────────╯ */
type AccountWithArrays struct {
	UID     string         `json:"UID,omitempty"`
	Profile Profile        `json:"profile,omitempty"`
	Data    DataWithArrays `json:"data,omitempty"`
}
type DataWithArrays struct {
	Events  EventsWithArrays `json:"events,omitempty"`
	Visited []string         `json:"visited,omitempty"`
}
type EventsWithArrays struct {
	Name []string `json:"name,omitempty"`
	When []string `json:"when,omitempty"`
}
type VisitedWithArrays struct {
	Name []string `json:"name,omitempty"`
}

func (account AccountWithArrays) GetGigyaAccount() Account {

	eventsAsArray := account.Data.Events
	visitedAsArray := account.Data.Visited

	eventsNameAsString := strings.Join(eventsAsArray.Name, ", ")
	eventsWhenAsString := strings.Join(eventsAsArray.When, ", ")

	visitedAsString := strings.Join(visitedAsArray, ", ")
	return Account{
		UID:     account.UID,
		Profile: account.Profile,
		Data: Data{
			Events: &Event{
				Name: eventsNameAsString,
				When: eventsWhenAsString,
			},
			Visited: visitedAsString,
		},
	}
}
