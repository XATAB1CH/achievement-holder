package routes

import (
	"net/http"

	"github.com/XATAB1CH/achievement-holder/handlers"
	mw "github.com/XATAB1CH/achievement-holder/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	auth := router.Group("/auth")
	{
		auth.Static("assets", "./assets")
		auth.Static("styles", "./assets/styles")

		auth.POST("/login", handlers.Login)
		auth.POST("/signup", handlers.Signup)
		auth.GET("/logout", handlers.Logout)
	}

	unknown := router.Group("/")
	{
		unknown.Static("assets", "./assets")
		unknown.Static("styles", "./assets/styles")

		unknown.GET("/", mw.IsAuthorized(), func(c *gin.Context) {
			claims, _ := c.Get("claims")
			if claims != nil {
				c.HTML(http.StatusOK, "home.html", claims)
				return
			}

			c.HTML(http.StatusOK, "index.html", nil)
		})

		unknown.GET("/registration", func(c *gin.Context) {
			c.HTML(http.StatusOK, "registration.html", nil)
		})
		unknown.GET("/logining", func(c *gin.Context) {
			c.HTML(http.StatusOK, "logining.html", nil)
		})
	}

	achievement := router.Group("/achievement")
	{
		achievement.Static("assets", "./assets")
		achievement.Static("styles", "./assets/styles")

		achievement.GET("/creation", func(c *gin.Context) {
			c.HTML(http.StatusOK, "creation.html", nil)
		})
	}

	logined := router.Group("/logined")
	{
		logined.Static("assets", "./assets")
		logined.Static("styles", "./assets/styles")

	}

	return router
}
