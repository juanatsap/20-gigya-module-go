package gigya

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func (g *Gigya) AccountsSearch(query string) (string, error) {

	// Add  parameters
	method := "accounts.search"
	params := map[string]string{"query": query, "apiKey": g.apiKey, "userKey": g.userKey, "secret": g.secretKey}

	// Prepare the request URL
	baseURL := fmt.Sprintf("https://%s/%s", g.apiDomain, method)
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}

	// Send the POST request
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func generateSignature(params map[string]string, secret string) string {
	// Sort the parameters alphabetically by key
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Concatenate the sorted parameters
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(params[key])
	}
	baseString := buf.String()

	// Decode the secret key
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return ""
	}

	// Compute the HMAC-SHA1 signature
	h := hmac.New(sha1.New, decodedSecret)
	h.Write([]byte(baseString))
	signature := h.Sum(nil)

	// Encode the signature in base64
	return base64.StdEncoding.EncodeToString(signature)
}
