package jwt_utils

import (
	"fmt"
	"log"
	"portfolio/server/models"
	"portfolio/server/responses"
	"portfolio/server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(c *gin.Context) (string, string) {
	key, err := ParsePrivateKey()

	if err != nil {
		fmt.Println("error parsing private key: %w", err)
		return "", ""
	}

	refid, err := utils.GenerateReferenceID()

	if err != nil {
		log.Fatal()
	}

	claims := models.CustomAccessClaims{
		ReferenceID: refid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Service Backend",
			Subject:   "Access",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	acc, err := token.SignedString(key)

	if err != nil {
		responses.Code500(c)
		log.Fatal(err)
	}

	return acc, refid
}

func GenerateRefreshToken(c *gin.Context) (string, string) {
	key, err := ParsePrivateKey()

	if err != nil {
		fmt.Println("error parsing private key: %w", err)
		return "", ""
	}

	sid, err := utils.GenerateReferenceID()

	if err != nil {
		log.Fatal(err)
	}

	claims := models.CustomRefreshClaims{
		SessionID: sid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Service Backend",
			Subject:   "Refresh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	ref, err := token.SignedString(key)

	if err != nil {
		responses.Code500(c)
		log.Fatal(err)
	}

	return ref, sid
}
