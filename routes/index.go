package routes

import (
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine, user *user.User) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", user)
	})

	router.GET("/registration", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.html", nil)
	})

	router.GET("/logining", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logining.html", nil)
	})

}
