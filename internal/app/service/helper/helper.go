package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSession() (string, error) {
	buf := make([]byte, 32)

	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}
