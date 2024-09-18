package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extract key from header
// Authorization: ApiKey {apikey}
func GetApiKey(headers http.Header) (string, error) {

	authHeader := headers.Get("Authorization")
	authHeaderParts := strings.Fields(authHeader)

	if len(authHeaderParts) == 0 {
		return "", errors.New("ApiKey not found")
	}

	if len(authHeaderParts) != 2 || authHeaderParts[0] != "ApiKey" {
		return "", errors.New("Malformed authorization header")
	}

	return authHeaderParts[1], nil
}
