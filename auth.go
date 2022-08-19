package main

import (
	"encoding/base64"
	"errors"
	"strings"
)

func authenticate(auth_header string) (bool, error) {
	if !strings.Contains(auth_header, "Basic ") {
		return false, errors.New("Authentication Header was not \"Basic BASE64_ENCODED_CREDENTIALS\"")
	}

	creds_encoded := strings.Split(auth_header, "Basic ")[1]
	creds_ascii, creds_ascii_err := base64.StdEncoding.DecodeString(creds_encoded)
	if creds_ascii_err != nil {
		return false, creds_ascii_err
	}
	creds := strings.Split(string(creds_ascii), ":")
	if len(creds) != 2 {
		return false, errors.New("Credentials could not be parsed. Use schema base64(\"username:password\")")
	}
	username := creds[0]
	password := strings.Split(creds[1], "\n")[0]

	if AUTH_USERS[username] == password {
		return true, nil
	} else {
		return false, errors.New("Invalid credentials")
	}
}
