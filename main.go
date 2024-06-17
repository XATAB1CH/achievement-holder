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
	testUser := User{
		username: "admin",
		loggedIn: false,
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": testUser.username,
		})
	})

	http.ListenAndServe(":8080", router)

}
