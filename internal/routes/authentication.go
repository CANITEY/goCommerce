package routes

import (
	"ecommerce/api/models"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"

)

func init() {
	gob.Register(&models.User{})
}

func signupController(w http.ResponseWriter, r *http.Request) {
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
		if err.Error() == "UNIQUE constraint failed: users.email" {
			http.Redirect(w, r, fmt.Sprintf("/auth?message=%v", "email used before"), http.StatusFound)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/auth?message=%v", err), http.StatusFound)
		return
	}

	http.Redirect(w, r, "/auth?message=success", http.StatusFound)
}

func loginController(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/auth?message=form data maliformed", http.StatusMovedPermanently)
		return
	}
	data := r.Form
	if ok := d.CheckUser(data.Get("email"), data.Get("password")); !ok {
		http.Redirect(w, r, "/auth?message=invalid email or password", http.StatusFound)
		return
	}

	user, err := d.GetUser(data.Get("email"))
	if err != nil {
		http.Redirect(w, r, "/auth?message=an error happened in backend system", http.StatusFound)
		fmt.Println(err)
		return
	}

	session, err := Store.Get(r, "session")
	if err != nil {
		http.Redirect(w, r, "/auth?message=an error happened in backend system", http.StatusFound)
		return
	}
	session.Values["user"] = user
	fmt.Println(session.Values["user"])
	if err := session.Save(r, w); err != nil {
		panic(err)
	}
	fmt.Println("SUCESS")
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func auth(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, "session")

	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/auth.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	message := r.URL.Query().Get("message")
	err = t.ExecuteTemplate(w, "base", message)
	if err != nil {
		log.Println(err)
	}
}
