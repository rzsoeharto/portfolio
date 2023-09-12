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
	SessionID string `json:"sid"`
	jwt.RegisteredClaims
}

func (c CustomRefreshClaims) GetSessionID() (string, error) {
	return c.SessionID, nil
}
