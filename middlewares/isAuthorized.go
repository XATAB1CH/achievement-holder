package middlewares

import (
	"fmt"
	"net/http"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	jwtKey = []byte("secret")
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "cant scan token"})
			return
		}

		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "cant parse token"})
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			fmt.Println(claims.Name)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token is not active"})
		}
	}
}
