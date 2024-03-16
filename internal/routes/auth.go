package routes

import (
	"ecommerce/api/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)


func signupView(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		fmt.Sprintln("Error parsing form parameter")
	}
	form := r.PostForm

	user, err := models.NewSignUpUser(form)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/auth?message=%v", err), http.StatusFound)
		return
	}

	if err := d.AddUser(*user); err != nil {
		http.Redirect(w, r, fmt.Sprintf("/auth?message=%v", err), http.StatusFound)
		return
	}

	http.Redirect(w, r, "/auth?message=success", http.StatusFound)
}

func auth(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/auth.tmpl",
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
