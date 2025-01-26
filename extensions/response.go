package extensions

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CDCResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func HandleExtensionsRequest(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Prepare the response
	response := CDCResponse{Status: "OK"}

	extensionPoint := requestBody["extensionPoint"].(string)
	data := requestBody["data"].(map[string]interface{})
	params := data["params"].(map[string]interface{})

	// OnBeforeAccountsRegister
	if extensionPoint == "OnBeforeAccountsRegister" {
		email := params["email"].(string)
		if !strings.HasSuffix(email, "@xyz.com") {
			customMessage := "Email should belong to domain 'xyz.com'"
			if lang, ok := params["lang"].(string); ok && lang == "he" {
				customMessage = "אימייל צריך להיות בדומיין איקס ווי זד"
			}
			response.Status = "FAIL"
			response.Data = map[string]interface{}{
				"validationErrors": []map[string]string{
					{"fieldName": "email", "message": customMessage},
				},
			}
		}
	}

	// OnBeforeAccountsLogin
	if extensionPoint == "OnBeforeAccountsLogin" {
		accountInfo, hasAccountInfo := data["accountInfo"].(map[string]interface{})
		if hasAccountInfo {
			profile, hasProfile := accountInfo["profile"].(map[string]interface{})
			if hasProfile && profile["firstName"] == "block" && profile["lastName"] == "me" {
				customMessage := "Your account is temporarily blocked"
				if lang, ok := params["lang"].(string); ok && lang == "he" {
					customMessage = "חשבונך חסום באופן זמני"
				}
				response.Status = "FAIL"
				response.Data = map[string]interface{}{
					"userFacingErrorMessage": customMessage,
				}
			}
		}
	}

	// OnBeforeSetAccountInfo
	if extensionPoint == "OnBeforeSetAccountInfo" {
		if profile, ok := params["profile"].(map[string]interface{}); ok {
			if firstName, ok := profile["firstName"].(string); ok {
				if strings.Contains(firstName, "fail") {
					customMessage := "Invalid name - contains a word with a negative meaning"
					if lang, ok := params["lang"].(string); ok && lang == "he" {
						customMessage = "שם לא חוקי - מכיל מילה עם משמעות שלילית"
					}
					response.Status = "FAIL"
					response.Data = map[string]interface{}{
						"validationErrors": []map[string]string{
							{"fieldName": "profile.firstName", "message": customMessage},
						},
					}
				} else if firstLetter := firstName[0:1]; strings.ToLower(firstLetter) == firstLetter {
					response.Status = "ENRICH"
					response.Data = map[string]interface{}{
						"profile": map[string]string{
							"firstName": strings.ToUpper(firstLetter) + firstName[1:],
						},
					}
				}
			}
		}
	}

	// Send the response
	c.JSON(http.StatusOK, response)
}
