# Gigya Module Go

## Overview

This Go module provides a comprehensive wrapper around Gigya's Customer Data Cloud (CDC) APIs, offering simplified access to Gigya's account management and authentication services through a clean Go interface. The module is designed to be easily integrated into any Go application requiring Gigya CDC integration.

## Features

- **Complete API Coverage**: Supports essential Gigya CDC API endpoints
- **Account Management**: Create, read, update, delete user accounts
- **Search Capabilities**: Advanced search functionality for user accounts
- **JWT Support**: Generate and validate JWT tokens
- **Error Handling**: Comprehensive error handling and reporting
- **Type-Safe Responses**: All API responses are properly typed

## Installation

```bash
go get github.com/yourusername/gigya-module-go
```

Or add it to your Go module dependencies:

```bash
go mod edit -require github.com/yourusername/gigya-module-go@latest
go mod tidy
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/yourusername/gigya-module-go/gigya"
)

func main() {
    // Initialize the Gigya client
    gigyaClient := gigya.NewGigya(
        "your-api-key",
        "your-user-key",
        "your-secret-key",
        "your-api-domain",
    )
    
    // Search for accounts
    accounts, total, err := gigyaClient.AccountsAPI.Search("email: \"*@example.com\"", 10)
    if err != nil {
        log.Fatalf("Error searching accounts: %v", err)
    }
    
    fmt.Printf("Found %d accounts out of %d total\n", len(accounts), total)
    for i, account := range accounts {
        fmt.Printf("%d. UID: %s, Email: %s\n", i+1, account.UID, account.Profile.Email)
    }
}
```

## Module Structure

- **gigya**: Main package providing the Gigya client
- **accounts**: Package handling account-related operations
- **jwt**: Package for JWT token operations
- **extensions**: Additional functionality extending the core capabilities
- **helpers**: Utility functions supporting the module's operations

## API Documentation

For detailed API documentation, see the [API Reference](./doc/api-reference.md).

## Examples

The module includes several examples demonstrating common use cases:

- [Account Search](./doc/examples.md#account-search)
- [Creating Users](./doc/examples.md#creating-users)
- [JWT Generation](./doc/examples.md#jwt-generation)
- [Deleting Accounts](./doc/examples.md#deleting-accounts)

## Error Handling

The module provides detailed error responses from the Gigya API. All errors include the original error code, reason, and details when available:

```go
accounts, _, err := gigyaClient.AccountsAPI.Search("invalid:query", 10)
if err != nil {
    // err contains the full error details from Gigya
    fmt.Printf("Error: %v\n", err)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For questions about this module, contact:
- Juan Andrés Moreno Rubio, Developer