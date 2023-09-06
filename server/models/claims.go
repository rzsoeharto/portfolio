package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ReferenceID string `json:"ref"`
	jwt.RegisteredClaims
}

func (c CustomClaims) GetReferenceID() (string, error) {
	return c.ReferenceID, nil
}
