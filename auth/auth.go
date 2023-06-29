package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: APIKey [some_value]
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Incorrect authentication header")
	}

	if vals[0] != "APIKey" {
		return "", errors.New("Incorrect first part of authentication header")
	}
	
	return vals[1], nil
}
