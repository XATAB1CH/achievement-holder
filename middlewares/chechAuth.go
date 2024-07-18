package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		cookie, err := c.Request.Cookie("token")

		if err != nil {
			c.Set("auth", "false")
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

		claims, ok := token.Claims.(*models.Claims)
		if ok && token.Valid {
			claims.Achievements, err = postgresql.GetAchievementsByUserID(conn, claims.UserID)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			c.Set("auth", "true")
		} else {
			c.Set("auth", "false")
			return
		}

	}
}
