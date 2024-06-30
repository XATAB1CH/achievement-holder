package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/registration", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.html", nil)
	})

	router.GET("/logining", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logining.html", nil)
	})

}
