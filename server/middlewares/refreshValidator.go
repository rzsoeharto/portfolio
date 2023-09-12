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

func RefreshValidator(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		responses.Code401(c, "Missing authorization header")
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == "" {
		fmt.Println("Trim")
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	fmt.Println(tokenString)

	key, err := jwt_utils.ParsePublicKey()

	if err != nil {
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
		fmt.Println(err)
		responses.Code401(c, "Token is invalid or expired")
		c.Abort()
		return
	}

	claims := token.Claims.(*models.CustomRefreshClaims)

	sub, subErr := claims.GetSubject()
	sid, sesErr := claims.GetSessionID()

	if subErr != nil && sesErr != nil {
		fmt.Println("sub", subErr)
		fmt.Println("ses", sesErr)
		responses.Code500(c)
		c.Abort()
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
