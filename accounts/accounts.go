package accounts

import "fmt"

// Account representa una cuenta individual en Gigya
type Account struct {
	LastUpdatedTimestamp int64                   `json:"lastUpdatedTimestamp"`
	Preferences          Preferences             `json:"preferences"`
	Subscriptions        map[string]Subscription `json:"subscriptions"`
	Data                 Data                    `json:"data"`
	Created              string                  `json:"created"`
	CreatedTimestamp     int64                   `json:"createdTimestamp"`
	Profile              Profile                 `json:"profile"`
	Channel              string                  `json:"channel"`
	Token                string                  `json:"token"`
	LastUpdated          string                  `json:"lastUpdated"`
	UID                  string                  `json:"UID"`
	HasLiteAccount       bool                    `json:"hasLiteAccount"`
	HasFullAccount       bool                    `json:"hasFullAccount"`
	Email                string                  `json:"email"`
}

func (a Account) Print() {
	fmt.Println("--------------------")
	fmt.Println("System:")
	fmt.Println("--------------------")
	fmt.Printf("Account UID: %s\n", a.UID)
	fmt.Printf("Created: %s\n", a.Created)
	fmt.Printf("Last Updated: %s\n", a.LastUpdated)
	fmt.Printf("UID: %s\n", a.UID)
	fmt.Printf("Has Lite Account: %t\n", a.HasLiteAccount)
	fmt.Printf("Has Full Account: %t\n", a.HasFullAccount)
	fmt.Printf("Last Updated Timestamp: %d\n", a.LastUpdatedTimestamp)
	fmt.Printf("Created Timestamp: %d\n", a.CreatedTimestamp)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Profile:")
	fmt.Println("--------------------")
	fmt.Printf("  First Name: %s\n", a.Profile.FirstName)
	fmt.Printf("  Last Name: %s\n", a.Profile.LastName)
	fmt.Printf("  Country: %s\n", a.Profile.Country)
	fmt.Printf("  Email: %s\n", a.Profile.Email)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Data:")
	fmt.Println("--------------------")
	fmt.Printf("  LIVX Member: %s\n", a.Data.LIVXMember)
	fmt.Printf("  Visited: %s\n", a.Data.Visited)
	fmt.Printf("  Competition: %s\n", a.Data.Competition)
	fmt.Printf("  Favorite Team: %s\n", a.Data.FavoriteTeam)
	fmt.Printf("  Data Source: %s\n", a.Data.DataSource)
	fmt.Printf("  Events: %s\n", a.Data.Events)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Preferences:")
	fmt.Println("--------------------")
	fmt.Printf("  Marketing: %v\n", a.Preferences.Marketing)
	fmt.Printf("  Terms: %v\n", a.Preferences.Terms)
	fmt.Printf("  Privacy: %v\n", a.Preferences.Privacy)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Subscriptions:")
	fmt.Println("--------------------")
	for channel, subscription := range a.Subscriptions {
		fmt.Printf("  %s: %v\n", channel, subscription)
	}
}

// Preferences representa las preferencias de la cuenta
type Preferences struct {
	Marketing map[string]ConsentDetail `json:"marketing"`
	Terms     map[string]ConsentDetail `json:"terms"`
	Privacy   map[string]ConsentDetail `json:"privacy"`
}

// ConsentDetail representa los detalles de consentimiento
type ConsentDetail struct {
	Entitlements        []string                `json:"entitlements"`
	Locales             map[string]LocaleDetail `json:"locales"`
	IsConsentGranted    bool                    `json:"isConsentGranted"`
	ActionTimestamp     string                  `json:"actionTimestamp"`
	CustomData          []string                `json:"customData"`
	Language            string                  `json:"language"`
	LastConsentModified string                  `json:"lastConsentModified"`
	DocVersion          float64                 `json:"docVersion,omitempty"`
	DocDate             string                  `json:"docDate,omitempty"`
	Tags                []string                `json:"tags"`
}

// LocaleDetail representa los detalles específicos de una localidad
type LocaleDetail struct {
	DocVersion float64 `json:"docVersion,omitempty"`
	DocDate    string  `json:"docDate,omitempty"`
}

// Subscription representa las suscripciones de la cuenta
type Subscription struct {
	Email SubscriptionChannel `json:"email"`
}

// SubscriptionChannel representa los detalles de una suscripción por canal
type SubscriptionChannel struct {
	IsSubscribed                 bool              `json:"isSubscribed"`
	LastUpdatedSubscriptionState string            `json:"lastUpdatedSubscriptionState"`
	DoubleOptIn                  DoubleOptInDetail `json:"doubleOptIn"`
}

// DoubleOptInDetail representa los detalles de doble opt-in
type DoubleOptInDetail struct {
	Status string `json:"status"`
}

// Data representa los datos personalizados de la cuenta
type Data struct {
	LIVXMember   string    `json:"LIVX_Member"`
	Visited      string    `json:"visited"`
	Competition  NameWhen  `json:"competition"`
	FavoriteTeam NameSince `json:"favoriteTeam"`
	DataSource   string    `json:"dataSource"`
	Events       NameWhen  `json:"events"`
}

// NameWhen representa una estructura con nombre y fecha
type NameWhen struct {
	Name string `json:"name"`
	When string `json:"when"`
}

// NameSince representa una estructura con nombre y fecha desde
type NameSince struct {
	Name  string `json:"name"`
	Since string `json:"since"`
}

// Profile representa el perfil de la cuenta
type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Email     string `json:"email"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

func table(accounts Accounts) {

	for _, account := range accounts.Accounts {
		fmt.Println(account.UID)
	}

}
