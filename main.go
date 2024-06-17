package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	username string
	loggedIn bool
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// логика ввода пароля
	logWord := "Авторизация успешна"
	fmt.Fprint(w, logWord)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := User{
		username: "admin",
		loggedIn: false,
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]string{"title": "Главная страница", "description": "Некая информация"})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", map[string]string{"title": "Страница входа", "description": "Некая информация"})
	})

	router.Run(":8080")
}
