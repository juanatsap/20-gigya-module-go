# Examples

This document provides practical examples of how to use the gigya-module-go package for common Gigya CDC operations.

## Table of Contents

- [Initialization](#initialization)
- [Account Search](#account-search)
- [Creating Users](#creating-users)
- [JWT Generation](#jwt-generation)
- [Deleting Accounts](#deleting-accounts)
- [Advanced Search Queries](#advanced-search-queries)
- [Error Handling](#error-handling)

## Initialization

Before performing any operations, you need to initialize the Gigya client:

```go
package main

import (
    "fmt"
    "log"
    
    "gigya-module-go/gigya"
)

func main() {
    // Initialize the Gigya client with your credentials
    gigyaClient := gigya.NewGigya(
        "your-api-key",
        "your-user-key",
        "your-secret-key",
        "us1.gigya.com", // or the appropriate datacenter domain
    )
    
    // Now you can use gigyaClient for various operations
}
```

## Account Search

### Basic Account Search

Search for accounts with a specific email pattern:

```go
func searchAccounts(gigyaClient *gigya.Gigya) {
    query := "email: \"*@example.com\""
    limit := 10
    
    accounts, total, err := gigyaClient.AccountsAPI.Search(query, limit)
    if err != nil {
        log.Fatalf("Error searching accounts: %v", err)
    }
    
    fmt.Printf("Found %d accounts out of %d total\n", len(accounts), total)
    for i, account := range accounts {
        fmt.Printf("%d. UID: %s, Email: %s\n", i+1, account.UID, account.Profile.Email)
    }
}
```

### Filtering by Registration Date

Search for accounts registered after a specific date:

```go
func searchRecentAccounts(gigyaClient *gigya.Gigya) {
    query := "created > \"2023-01-01 00:00:00\""
    limit := 20
    
    accounts, total, err := gigyaClient.AccountsAPI.Search(query, limit)
    if err != nil {
        log.Fatalf("Error searching accounts: %v", err)
    }
    
    fmt.Printf("Found %d accounts out of %d total\n", len(accounts), total)
}
```

## Creating Users

Register a new user account:

```go
func registerUser(gigyaClient *gigya.Gigya) {
    params := map[string]string{
        "email": "newuser@example.com",
        "password": "SecurePassword123!",
        "profile.firstName": "John",
        "profile.lastName": "Doe",
    }
    
    response, err := gigyaClient.AccountsAPI.Register(params)
    if err != nil {
        log.Fatalf("Error registering user: %v", err)
    }
    
    fmt.Printf("User registered successfully. UID: %s\n", response.UID)
}
```

## JWT Generation

Generate a JWT token for a user:

```go
func generateJWT(gigyaClient *gigya.Gigya, uid string) {
    params := map[string]string{
        "targetUID": uid,
        "expires": "3600",  // expires in 1 hour
    }
    
    jwtResponse, err := gigyaClient.AccountsAPI.GetJWT(params)
    if err != nil {
        log.Fatalf("Error generating JWT: %v", err)
    }
    
    fmt.Printf("JWT token: %s\n", jwtResponse.ID_Token)
}
```

## Deleting Accounts

### Delete a Single Account

Delete an account by its UID:

```go
func deleteAccount(gigyaClient *gigya.Gigya, uid string) {
    response, err := gigyaClient.AccountsAPI.DeleteAccount(uid)
    if err != nil {
        log.Fatalf("Error deleting account: %v", err)
    }
    
    fmt.Println("Account deleted successfully")
}
```

### Delete Multiple Accounts

Delete all accounts with a specific import ID:

```go
func deleteAccountsByImportId(gigyaClient *gigya.Gigya, importId string) {
    deletedAccounts, err := gigyaClient.AccountsAPI.DeleteAccountsForIdxImportId(importId)
    if err != nil {
        log.Fatalf("Error deleting accounts: %v", err)
    }
    
    fmt.Printf("Deleted %d accounts\n", len(deletedAccounts))
    for i, account := range deletedAccounts {
        fmt.Printf("%d. Deleted UID: %s, Email: %s\n", i+1, account.UID, account.Profile.Email)
    }
}
```

## Advanced Search Queries

Gigya's search API supports complex SQL-like queries:

```go
func advancedSearch(gigyaClient *gigya.Gigya) {
    // Find accounts that have a verified email and have logged in recently
    query := "emails.verified: \"*@example.com\" AND lastLogin > \"2023-06-01\""
    limit := 50
    
    accounts, total, err := gigyaClient.AccountsAPI.Search(query, limit)
    if err != nil {
        log.Fatalf("Error searching accounts: %v", err)
    }
    
    fmt.Printf("Found %d matching accounts\n", total)
}
```

## Error Handling

Proper error handling for Gigya API responses:

```go
func errorHandlingExample(gigyaClient *gigya.Gigya) {
    // Try to get an account with an invalid UID
    _, err := gigyaClient.AccountsAPI.GetAccount("invalid-uid")
    if err != nil {
        // Check if it's a specific error type
        fmt.Printf("Error occurred: %v\n", err)
        
        // You could parse the error message to check for specific error codes
        if contains(err.Error(), "400102") {
            fmt.Println("This is a 'UID not found' error")
        } else if contains(err.Error(), "403") {
            fmt.Println("This is a permission error")
        }
    }
}

func contains(s, substr string) bool {
    return strings.Contains(s, substr)
}
```