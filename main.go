package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", loginHandler)

	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
