package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func adminProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["action"] == "delete" {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Redirect(w, r, "/admin/products", http.StatusFound)
		}
		if ok, err := d.DeleteProduct(uint(id)); err != nil || !ok{
			fmt.Println(err)
			http.Redirect(w, r, "/admin/products", http.StatusFound)
		}
	}
	products, err := d.GetProducts()
	if err != nil {
		panic(err)
	}
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/admin.tmpl",
		"./web/templates/pages/admin/products.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	
	err = t.ExecuteTemplate(w, "base", products)
	if err != nil {
		log.Println(err)
	}
}

