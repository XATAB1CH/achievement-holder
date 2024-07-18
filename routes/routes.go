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

		auth.GET("/registration", func(c *gin.Context) {
			c.HTML(http.StatusOK, "registration.html", nil)
		})
		auth.GET("/logining", func(c *gin.Context) {
			c.HTML(http.StatusOK, "logining.html", nil)
		})
		auth.POST("/login", handlers.Login)
		auth.POST("/signup", handlers.Signup)
		auth.GET("/logout", handlers.Logout)
	}

	api := router.Group("/")
	{
		api.Static("assets", "./assets")
		api.Static("styles", "./assets/styles")

		api.GET("/", mw.IsAuthorized(), func(c *gin.Context) {
			claims, _ := c.Get("claims")
			if claims != nil {
				c.HTML(http.StatusOK, "home.html", claims)
				return
			}

			c.HTML(http.StatusOK, "index.html", nil)
		})
	}

	achievement := router.Group("/achievement")
	{
		achievement.Static("assets", "./assets")
		achievement.Static("styles", "./assets/styles")

		achievement.GET("/:id", mw.CheckAuth(), handlers.Information)
		achievement.GET("/creation", func(c *gin.Context) {
			c.HTML(http.StatusOK, "creation.html", nil)
		})
		achievement.POST("/create", mw.IsAuthorized(), handlers.Create)
	}

	demo := router.Group("/demo")
	{
		demo.Static("assets", "./assets")
		demo.Static("styles", "./assets/styles")

		api.POST("/search", handlers.Search)
		api.GET("/:id", handlers.Demo)
	}

	return router
}
