package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// The headers of an HTTP request
// Exemple :
// Authorization: ApiKey {insert apikey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentification info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed forst part of auth header")
	}
	return vals[1], nil
}
