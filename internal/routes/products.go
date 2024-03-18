package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func products(w http.ResponseWriter, r *http.Request) {
	products, err := d.GetProducts()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "500 Internal Error")
	}
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/products.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "base", products)
	if err != nil {
		panic(err)
	}
}

func product(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Enter correct id")
	}
	product, err := d.GetProduct(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Product not found")
		return
	}

	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/product.tmpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "base", product)
	if err != nil {
		panic(err)
	}
}
