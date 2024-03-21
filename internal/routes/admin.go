package routes

import (
	"html/template"
	"log"
	"net/http"
)

func adminPanel(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		panic(err)
	}
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/admin.tmpl",
		"./web/templates/pages/admin/home.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	
	data := session.Values["user"]
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

