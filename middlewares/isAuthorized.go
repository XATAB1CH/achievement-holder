package middlewares

import (
	"github.com/XATAB1CH/achievement-holder/utils"
	"github.com/gin-gonic/gin"
)

func IsUserized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unUserized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unUserized"})
			c.Abort()
			return
		}

		c.Set("regular", claims.Role)
		c.Next()
	}
}
