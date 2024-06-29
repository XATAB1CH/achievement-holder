package routes

import (
	"net/http"

	author "github.com/XATAB1CH/achievement-holder/models/author"
	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine, author *author.Author) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", author)
	})

	router.GET("/registration", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.html", nil)
	})

	router.GET("/logining", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logining.html", nil)
	})

}
