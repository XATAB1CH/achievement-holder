package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func loginHandler(c *gin.Context) {
// 	// логика ввода пароля
// 	logWord := "Авторизация успешна"

// 	username := c.PostForm("username")

// 	c.HTML(http.StatusOK, "index.html", map[string]string{"title": logWord, "description": fmt.Sprintf("Добро пожаловать %s", username)})
// }

func main() { // все мтеоды get post

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")
	router.Static("styles",  "./assets/styles")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]string{"title": "Главная страница", "description": "Некая информация"})
	})

	router.GET("/signup", func(c  *gin.Context)  {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	router.GET("/signin", func(c  *gin.Context)  {
		c.HTML(http.StatusOK, "signin.html", nil)
	})

	// router.POST("/submit", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", map[string]string{"title": "Главная страница", "description": "Некая информация"})
	// })

	router.Run(":8080")
}
