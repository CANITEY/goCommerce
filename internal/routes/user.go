package routes

import (
	"html/template"
	"log"
	"net/http"
)

func profile(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		panic(err)
	}

	user := session.Values["user"]
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/profile.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "base", user)
	if err != nil {
		log.Println(err)
	}
}

func cart(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/cart.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "base", "Hello world!")
	if err != nil {
		log.Println(err)
	}
}
