package middlewares

import (
	"fmt"
	"portfolio/server/jwt_utils"
	logger "portfolio/server/logs"
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
		logger.Logger.Println("Error parsing public key: ", err)
		responses.Code500(c)
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomAccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil || !token.Valid {
		logger.Logger.Println("Error parsing token", err)
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	subj, err := token.Claims.GetSubject()

	if err != nil {
		logger.Logger.Println("Error fetching Subject: ", err)
		responses.Code500(c)
		c.Abort()
		return
	}

	if subj != "Access" {
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	claims := token.Claims.(*models.CustomAccessClaims)

	ref, err := claims.GetReferenceID()

	if err != nil {
		logger.Logger.Println("Error fetching reference id: ", err)
		responses.Code500(c)
		c.Abort()
		return
	}

	c.Set("Permission", ref)

	c.Next()
}
