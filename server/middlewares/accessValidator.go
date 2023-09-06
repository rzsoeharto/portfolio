package middlewares

import (
	"fmt"
	"portfolio/server/jwt_utils"
	"portfolio/server/models"
	"portfolio/server/responses"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AccessValidator(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		responses.Code401(c, "Missing authorization header")
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == "" {
		responses.Code401(c, "Token is invalid or expired")
		return
	}

	key, err := jwt_utils.ParsePublicKey()

	if err != nil {
		responses.Code500(c)
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil

	})

	if err != nil || !token.Valid {
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	if err != nil {
		responses.Code500(c)
		c.Abort()
		return
	}

	claims := token.Claims.(*models.CustomClaims)

	ref, err := claims.GetReferenceID()

	if err != nil {
		fmt.Println(err)
		c.Abort()
	}

	c.Set("Permission", ref)

	c.Next()
}