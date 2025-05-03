# API Reference

## Table of Contents

- [Gigya Client](#gigya-client)
- [Accounts API](#accounts-api)
  - [Search](#search)
  - [Get Account](#get-account)
  - [Get Account Info](#get-account-info)
  - [Set Account Info](#set-account-info)
  - [Delete Account](#delete-account)
  - [Search Accounts For IdxImportId](#search-accounts-for-idximportid)
  - [Delete Accounts For IdxImportId](#delete-accounts-for-idximportid)
- [JWT Functions](#jwt-functions)
  - [Get JWT Public Key](#get-jwt-public-key)

## Gigya Client

### NewGigya

Creates a new Gigya client instance.

```go
func NewGigya(apiKey, userKey, secretKey, apiDomain string) *Gigya
```

**Parameters:**
- `apiKey` - Your Gigya API key
- `userKey` - Your Gigya user key
- `secretKey` - Your Gigya secret key
- `apiDomain` - The API domain to use (e.g., "us1.gigya.com")

**Returns:**
- A pointer to a new Gigya instance

### Configuration Methods

```go
func (g *Gigya) SetApiKey(apiKey string)
func (g *Gigya) SetUserKey(userKey string)
func (g *Gigya) SetSecretKey(secretKey string)
func (g *Gigya) SetApiDomain(apiDomain string)
```

## Accounts API

The AccountsAPI provides access to Gigya's account management functionality.

### Search

Searches for accounts using a SQL-like query.

```go
func (a *AccountsAPI) Search(query string, limit int) (Accounts, int, error)
```

**Parameters:**
- `query` - SQL-like query string (e.g., "email: \"*@example.com\"")
- `limit` - Maximum number of results to return

**Returns:**
- A list of accounts matching the query
- The total count of matching accounts
- Any error that occurred

**Example:**
```go
accounts, total, err := gigyaClient.AccountsAPI.Search("email: \"*@example.com\"", 10)
```

### Get Account

Retrieves account information for a specific UID.

```go
func (a *AccountsAPI) GetAccount(uid string) (Account, error)
```

**Parameters:**
- `uid` - The Gigya UID of the account to retrieve

**Returns:**
- The account object
- Any error that occurred

**Example:**
```go
account, err := gigyaClient.AccountsAPI.GetAccount("1234567890")
```

### Get Account Info

Retrieves specific account information based on provided parameters.

```go
func (a *AccountsAPI) GetAccountInfo(uid string, params map[string]string) (AccountInfo, error)
```

**Parameters:**
- `uid` - The Gigya UID of the account
- `params` - Additional parameters to pass to the API

**Returns:**
- Account information response
- Any error that occurred

### Set Account Info

Updates account information.

```go
func (a *AccountsAPI) SetAccountInfo(uid string, params map[string]string) (SetAccountInfoResponse, error)
```

**Parameters:**
- `uid` - The Gigya UID of the account to update
- `params` - The parameters to update

**Returns:**
- Response from the setAccountInfo API call
- Any error that occurred

### Delete Account

Deletes an account by UID.

```go
func (a *AccountsAPI) DeleteAccount(uid string) (DeleteAccountResponse, error)
```

**Parameters:**
- `uid` - The Gigya UID of the account to delete

**Returns:**
- Response from the deleteAccount API call
- Any error that occurred

### Search Accounts For IdxImportId

Searches for accounts with a specific idxImportId.

```go
func (a *AccountsAPI) SearchAccountsForIdxImportId(idxImportId string) ([]Account, error)
```

**Parameters:**
- `idxImportId` - The idxImportId to search for

**Returns:**
- List of matching accounts
- Any error that occurred

### Delete Accounts For IdxImportId

Deletes all accounts with a specific idxImportId.

```go
func (a *AccountsAPI) DeleteAccountsForIdxImportId(idxImportId string) ([]Account, error)
```

**Parameters:**
- `idxImportId` - The idxImportId to search for and delete

**Returns:**
- List of deleted accounts
- Any error that occurred

## JWT Functions

### Get JWT Public Key

Retrieves the JWT public key from Gigya.

```go
func (a *AccountsAPI) GetJWTPublicKey() (GetJWTPublicKeyResponse, error)
```

**Returns:**
- The JWT public key response
- Any error that occurred