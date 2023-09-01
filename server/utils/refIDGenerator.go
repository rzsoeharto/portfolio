package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateReferenceID() (string, error) {
	randBytes := make([]byte, 16)

	_, err := rand.Read(randBytes)

	if err != nil {
		return "", err
	}

	refId := base64.URLEncoding.EncodeToString(randBytes)

	return refId, nil
}
