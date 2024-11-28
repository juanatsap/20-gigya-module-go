package accounts

import (
	"encoding/json"
	"fmt"
)

// Account representa una cuenta individual en Gigya
type Account struct {
	UID                  string                  `json:"UID,omitempty"`
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
}
type Password struct {
	Created        string       `json:"created,omitempty"`
	HashedPassword string       `json:"hashedPassword,omitempty"`
	HashSettings   HashSettings `json:"hashSettings,omitempty"`
}
type HashSettings struct {
	Algorithm string `json:"algorithm,omitempty"`
	Rounds    int    `json:"rounds,omitempty"`
	Salt      string `json:"salt,omitempty"`
}
type Emails struct {
	Verified   []string `json:"verified,omitempty"`
	Unverified []string `json:"unverified,omitempty"`
}

type LoginIDs struct {
	Emails []string `json:"emails,omitempty"`
}

func (a Account) Print() {
	fmt.Println("--------------------")
	fmt.Println("System:")
	fmt.Println("--------------------")
	fmt.Printf("Account UID: %s\n", a.UID)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Profile:")
	fmt.Println("--------------------")
	fmt.Printf("  First Name: %s\n", a.Profile.FirstName)
	fmt.Printf("  Last Name: %s\n", a.Profile.LastName)
	fmt.Printf("  Country: %s\n", a.Profile.Country)
	fmt.Printf("  Email: %s\n", a.Profile.Email)
	fmt.Println("")
	if a.Data.Competition != nil || a.Data.FavoriteTeam != nil || a.Data.Visited != "" || a.Data.Events != nil {

		fmt.Println("--------------------")
		fmt.Println("Data:")
		fmt.Println("--------------------")
		// fmt.Printf("  LIVX Member: %s\n", a.Data.LIVXMember)
		// fmt.Printf("  Visited: %s\n", a.Data.Visited)
		// fmt.Printf("  Competition: %s\n", *a.Data.Competition.Name)
		// fmt.Printf("  Favorite Team: %s\n", *a.Data.FavoriteTeam.Name)
		// fmt.Printf("  Data Source: %s\n", a.Data.DataSource)
		fmt.Printf("  Events: %s\n", a.Data.Events)
		fmt.Println("")
	}
	fmt.Println("--------------------")
	fmt.Println("Preferences:")
	fmt.Println("--------------------")
	fmt.Printf("  Marketing: %v\n", a.Preferences.Marketing.Email.IsConsentGranted)
	fmt.Printf("  Terms: %v\n", a.Preferences.Terms.ToS.IsConsentGranted)
	fmt.Printf("  Privacy: %v\n", a.Preferences.Privacy.Livgolf.IsConsentGranted)
	fmt.Println("")
	fmt.Printf("Created: %s\n", a.Created)
	fmt.Printf("Last Updated: %s\n", a.LastUpdated)
	fmt.Printf("Has Lite Account: %t\n", a.HasLiteAccount)
	fmt.Printf("Has Full Account: %t\n", a.HasFullAccount)
	fmt.Printf("Last Updated Timestamp: %d\n", a.LastUpdatedTimestamp)
	fmt.Printf("Created Timestamp: %d\n", a.CreatedTimestamp)
}

func (a Account) PrintShort() {
	fmt.Println("--------------------")
	fmt.Println("Profile:")
	fmt.Println("--------------------")
	fmt.Printf("  Account UID: %s\n", a.UID)
	fmt.Printf("  First Name: %s\n", a.Profile.FirstName)
	fmt.Printf("  Last Name: %s\n", a.Profile.LastName)
	fmt.Printf("  Country: %s\n", a.Profile.Country)
	fmt.Printf("  Email: %s\n", a.Profile.Email)
	fmt.Printf("  Marketing (OPT-IN): %v\n", a.Preferences.Marketing.Email.IsConsentGranted)
	fmt.Printf("  Terms: %v\n", a.Preferences.Terms.ToS.IsConsentGranted)
	fmt.Printf("  Privacy: %v\n", a.Preferences.Privacy.Livgolf.IsConsentGranted)
	fmt.Printf("  Created: %s\n", a.Created)
	fmt.Printf("  IdxImportID: %s\n", a.Data.IdxImportId)
	fmt.Println("---------------------------------------------------")
}

// Preferences representa las preferencias de la cuenta
type Preferences struct {
	Marketing Marketing `json:"marketing,omitempty"`
	Terms     ToS       `json:"terms,omitempty"`
	Privacy   Livgolf   `json:"privacy,omitempty"`
}

// Livgolf representa las preferencias de la cuenta
type Livgolf struct {
	Livgolf ConsentDetail `json:"livgolf,omitempty"`
}
type ToS struct {
	ToS ConsentDetail `json:"ToS,omitempty"`
}
type Marketing struct {
	Email ConsentDetail `json:"email,omitempty"`
}

// ConsentDetail representa los detalles de consentimiento
type ConsentDetail struct {
	Entitlements        []string                `json:"entitlements,omitempty"`
	Locales             map[string]LocaleDetail `json:"locales,omitempty"`
	IsConsentGranted    bool                    `json:"isConsentGranted,omitempty"`
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

// Data representa los datos personalizados de la cuenta
type Data struct {
	IdxImportId string `json:"idxImportId,omitempty"`

	/* ╭──────────────────────────────────────────╮ */
	/* │                 LIVGOLF                  │ */
	/* ╰──────────────────────────────────────────╯ */
	// LIVXMember string `json:"LIVX_Member,omitempty"`
	Visited      string     `json:"visited,omitempty"`
	Competition  *NameWhen  `json:"competition,omitempty"`
	FavoriteTeam *NameSince `json:"favoriteTeam,omitempty"`
	// DataSource string `json:"dataSource,omitempty"`
	Events   *Event    `json:"events,omitempty"`
	RipperGC *RipperGC `json:"rippergc,omitempty"`

	/* ╭──────────────────────────────────────────╮ */
	/* │                 OLYMPICS                 │ */
	/* ╰──────────────────────────────────────────╯ */
	// Add data.utility.isAthlete
	Utility *Utility `json:"utility,omitempty"`
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

// NameSince representa una estructura que puede ser un string, un array o un objeto vacío
type NameSince struct {
	Name  string `json:"name,omitempty"`
	Since string `json:"since,omitempty"`
}

// Profile representa el perfil de la cuenta
type Profile struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
}

func (a Account) AsJSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a Profile) AsJSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}

type Accounts []Account

func (accounts Accounts) Table() {

}
