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

func RefreshValidator(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		responses.Code401(c, "Missing authorization header")
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(header, "Refresh ")
	if tokenString == "" {
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	key, err := jwt_utils.ParsePublicKey()

	if err != nil {
		logger.Logger.Println("Error parsing public key: ", err)
		responses.Code500(c)
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	claims := token.Claims.(*models.CustomRefreshClaims)

	sub, subErr := claims.GetSubject()
	sid, sesErr := claims.GetSessionID()

	if subErr != nil && sesErr != nil {
		logger.Logger.Println("Error fetching subject: ", subErr, "Error fetching session: ", sesErr)
		c.Abort()
		responses.Code500(c)
		return
	}

	if sub != "Refresh" {
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}
	c.Set("Refresh Token", tokenString)
	c.Set("Session ID", sid)

	c.Next()
}
