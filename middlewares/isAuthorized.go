package middlewares

import (
	"fmt"
	"net/http"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			// Если токен отсутствует или ошибка, отправляем ответ с кодом 401
			c.JSON(http.StatusUnauthorized, gin.H{"message": "cant scan token"})
			return
		}

		// Проверяем токен
		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Проверяем, что используется правильный алгоритм подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Возвращаем секретный ключ для проверки подписи
			return []byte("secret"), nil
		})
		if err != nil {
			// Если произошла ошибка при проверке токена, отправляем ответ с кодом 401
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token beda"})
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			// Токен действителен, продолжаем обработку запроса
			fmt.Println(claims)
			c.Next()
		} else {
			// Отправляем ответ с кодом 401
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token is not active"})
		}
	}
}
