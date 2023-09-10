package models

import "github.com/golang-jwt/jwt/v5"

type CustomAccessClaims struct {
	ReferenceID string `json:"ref"`
	jwt.RegisteredClaims
}

func (c CustomAccessClaims) GetReferenceID() (string, error) {
	return c.ReferenceID, nil
}

type CustomRefreshClaims struct {
	jwt.RegisteredClaims
	SessionID string `json:"sid"`
}

func (c CustomRefreshClaims) GetSessionID() (string, error) {
	return c.SessionID, nil
}
