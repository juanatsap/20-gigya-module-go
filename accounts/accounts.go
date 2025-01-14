package accounts

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
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
	RegSource            string                  `json:"regSource,omitempty"`
}
type Accounts []Account

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
func (a Account) PrintLine(index int) {
	// Only in one line
	fmt.Printf("%d. UID: %s, Email: %s\n", index, a.UID, a.Profile.Email)

}
func (a Account) PrintLineWith(index int, field string) {

	switch field {
	case "uid":
		fmt.Printf("%d. UID: %s\n", index, a.UID)
	case "email":
		fmt.Printf("%d. Email: %s\n", index, a.Profile.Email)
	case "competition":
		fmt.Printf("%d. Competition Name: %s\n", index, a.Data.Competition.Name)
		fmt.Printf("%d. Competition When: %s\n", index, a.Data.Competition.When)
	case "favoriteTeam":
		fmt.Printf("%d. Favorite Team Name: %s\n", index, a.Data.FavoriteTeam.Name)
		fmt.Printf("%d. Favorite Team Since: %s\n", index, a.Data.FavoriteTeam.Since)
	case "visited":
		fmt.Printf("%d. Visited: %s\n", index, a.Data.Visited)
	case "events":
		fmt.Printf("%d. Events Name: %s\n", index, a.Data.Events.Name)
		fmt.Printf("%d. Events When: %s\n", index, a.Data.Events.When)
	default:
		fmt.Printf("%d. %s: %s\n", index, field, a.Profile.Email)
	}
}
func (a Account) PrintLineWithText(title string, index int, field string) {

	log.Debugf("%s - UID: %s, Email: %s\n", title, a.UID, a.Profile.Email)
	switch field {
	case "uid":
		fmt.Printf("%d. UID: %s\n", index, a.UID)
	case "email":
		fmt.Printf("%d. Email: %s\n", index, a.Profile.Email)
	case "competition":
		fmt.Printf("%d. Competition Name: %s\n", index, a.Data.Competition.Name)
		fmt.Printf("%d. Competition When: %s\n", index, a.Data.Competition.When)
	case "favoriteTeam":
		fmt.Printf("%d. Favorite Team Name: %s\n", index, a.Data.FavoriteTeam.Name)
		fmt.Printf("%d. Favorite Team Since: %s\n", index, a.Data.FavoriteTeam.Since)
	case "visited":
		fmt.Printf("%d. Visited: %s\n", index, a.Data.Visited)
	case "events":
		fmt.Printf("%d. Events Name: %s\n", index, a.Data.Events.Name)
		fmt.Printf("%d. Events When: %s\n", index, a.Data.Events.When)
	default:
		fmt.Printf("%d. %s: %s\n", index, field, a.Profile.Email)
	}
}

func (a Account) AsJSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a Profile) AsJSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (accounts Accounts) Table() {
}
func (a *Account) FixCompetition() {

	if a.Data.Competition != nil {

		if a.Data.Competition.Name == "" {
			a.Data.Competition.Name = ""
			a.Data.Competition.When = ""
		}
	}
}
func (a *Account) FixFavoriteTeam() {

	if a.Data.FavoriteTeam != nil {

		if a.Data.FavoriteTeam.Name == "" {
			a.Data.FavoriteTeam.Name = ""
			a.Data.FavoriteTeam.Since = ""
		}
	}
}

func (a *Account) FixVisited() {
	if a.Data.Visited == "" {
		a.Data.Visited = ""
	}
}
