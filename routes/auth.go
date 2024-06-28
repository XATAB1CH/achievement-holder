package routes

import (
	"github.com/XATAB1CH/achievement-holder/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", handlers.Login)
	r.POST("/signup", handlers.Signup)
	// r.GET("/home", handlers.Home)
	// r.GET("/premium", handlers.Premium)
	// r.GET("/logout", handlers.Logout)
}
