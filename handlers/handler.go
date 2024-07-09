package handlers

import (
	"net/http"

	mw "github.com/XATAB1CH/achievement-holder/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")
	router.Static("styles", "./assets/styles")

	auth := router.Group("/auth")
	{
		auth.POST("/login", Login)
		auth.POST("/signup", Signup)
	}

	api := router.Group("/")
	{
		api.GET("/", mw.IsAuthorized(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		api.GET("/registration", func(c *gin.Context) {
			c.HTML(http.StatusOK, "registration.html", nil)
		})

		api.GET("/logining", func(c *gin.Context) {
			c.HTML(http.StatusOK, "logining.html", nil)
		})
	}

	return router
}
