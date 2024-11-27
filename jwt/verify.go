package jwt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
)

// sanitizeAndBase64Decode replaces URL-safe base64 characters and decodes the string.
func sanitizeAndBase64Decode(str string) ([]byte, error) {
	// Replace URL-safe characters with standard base64 characters
	str = strings.ReplaceAll(str, "-", "+")
	str = strings.ReplaceAll(str, "_", "/")

	// Check if padding is needed for base64 decoding (should be a multiple of 4)
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	// Decode base64 string
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	return data, nil
}

func VerifyRSASignature(idToken, nStr, eStr string) (bool, error) {
	// Split JWT token
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid token format")
	}

	// Decode signature
	signature, err := sanitizeAndBase64Decode(parts[2])
	if err != nil {
		return false, err
	}

	// Decode modulus and exponent
	modulus, err := sanitizeAndBase64Decode(nStr)
	if err != nil {
		return false, err
	}
	exponent, err := sanitizeAndBase64Decode(eStr)
	if err != nil {
		return false, err
	}

	// Convert modulus and exponent to big.Int
	n := new(big.Int).SetBytes(modulus)
	eBytes := make([]byte, 8)
	copy(eBytes[8-len(exponent):], exponent)
	e := int(binary.BigEndian.Uint64(eBytes))

	// Create RSA public key
	pubKey := &rsa.PublicKey{N: n, E: e}

	// Concatenate header and payload
	tokenData := parts[0] + "." + parts[1]

	// Hash the data using SHA-256
	hashed := sha256.Sum256([]byte(tokenData))

	// Verify the signature using RSA
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signature)
	return err == nil, err
}

func VerifyToken() (bool, error) {
	// Example values (replace with actual values)
	idToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlJFUTBNVVE1TjBOQ1JUSkVNemszTTBVMVJrTkRRMFUwUTBNMVJFRkJSamhETWpkRU5VRkJRZyJ9.eyJpc3MiOiJodHRwczovL2ZpZG0uZ2lneWEuY29tL2p3dC80X2ZLN2Vud2JUUkdPakJUeDlHQ1E2WmcvIiwiYXBpS2V5IjoiNF9mSzdlbndiVFJHT2pCVHg5R0NRNlpnIiwiaWF0IjoxNzMxMjk3MzY5LCJleHAiOjE3MzEyOTc2NjksInN1YiI6IjRjN2MyMDk4OWY0NjQzZGM5N2VlZjlmZWE3YjAwODZlIn0.p--rDFJox1jqe-I4SbhO9V0uDJXW_sXnWI8x9_vLV7t0TYS5tDthls55HxQemw6_ZS5g2Mt7i15heXr-qG_dI9QLf2NaKLdbm88UmexXGgtnhkPdkkoJ93sHle-ZRdbGGdktsI9h4BEFZ4LqvfN4tA8x7k-NbNUKlvqWAdSgUF59MlHUkiWBVRaGSZEvdq-8g0fRZrDfWgcAutIjE7aZM2fn3rzHOgmIrOukejPhmyEU7yHg7Q4oSMn62nn15fqhvAdjiOu3NGERWO20joAiVDpX1xkmwSjq5T6_70QvON9e_IiEoevNpggPJnOPjDImFYWbxEXwBb5sbb1EVdXGhA"
	n := "qoQah4MFGYedrbWwFc3UkC1hpZlnB2_E922yRJfHqpq2tTHL_NvjYmssVdJBgSKi36cptKqUJ0Phui9Z_kk8zMPrPfV16h0ZfBzKsvIy6_d7cWnn163BMz46kAHtZXqXhNuj19IZRCDfNoqVVxxCIYvbsgInbzZM82CB86iYPAS7piijYn1S6hueVHGAzQorOetZevKIAvbH3kJXZ4KdY6Ffz5SFDJBxC3bycN4q2JM1qnyD53vcc0MitxyIUF7a06iJb5_xXBiA-3xnTI0FU5hw_k6x-sdB5Rglx13_2aNzdWBSBAnxs1XXtZUt9_2RAUxP1XORkrBGlPg9D7cBtQ"
	e := "AQAB"
	// Verify the token
	isValid, err := VerifyRSASignature(idToken, n, e)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
	} else if isValid {
		fmt.Println("=> Valid Signature!")
	} else {
		fmt.Println("=> Invalid Signature!")
	}

	return isValid, err
}
