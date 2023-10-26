package middlewares

import (
	"fmt"
	"portfolio/server/jwt_utils"
	logger "portfolio/server/logs"
	"portfolio/server/models"
	"portfolio/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RefreshValidator(c *gin.Context) {
	refreshToken, cookieErr := c.Cookie("Refresh-Token")

	if cookieErr != nil {
		responses.Code401(c, "Missing refresh cookie")
		c.Abort()
		return
	}

	if refreshToken == "" {
		responses.Code401(c, "Missing authorization token")
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

	token, err := jwt.ParseWithClaims(refreshToken, &models.CustomRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			errorMsg := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			logger.Logger.Printf(errorMsg)
			return nil, fmt.Errorf(errorMsg)
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
	c.Set("Refresh Token", refreshToken)
	c.Set("Session ID", sid)

	c.Next()
}
