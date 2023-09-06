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

func GenerateAccessToken(c *gin.Context, data *models.User) (string, string) {
	key, err := ParsePrivateKey()

	if err != nil {
		fmt.Println("error parsing private key: %w", err)
		return "", ""
	}

	refid, err := utils.GenerateReferenceID()

	if err != nil {
		log.Fatal()
	}

	claims := models.CustomClaims{
		ReferenceID: refid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Service Backend",
			Subject:   data.Username,
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

func GenerateRefreshToken(c *gin.Context) string {
	key, err := ParsePrivateKey()

	if err != nil {
		fmt.Println("error parsing private key: %w", err)
		return ""
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "Service Backend",
		Subject:   "Refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	ref, err := token.SignedString(key)

	if err != nil {
		responses.Code500(c)
		log.Fatal(err)
	}

	return ref
}
