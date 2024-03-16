package routes

import (
	"ecommerce/internal/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	d   database.DB
	err error
)

func init() {
	d, err = database.NewConn()
	if err != nil {
		panic(err)
	}

	// adding static fileServer
	r := mux.NewRouter()
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("web/static/js"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("web/static/css"))))


	r.HandleFunc("/", home)
	r.HandleFunc("/products", products)
	r.HandleFunc("/products/{id}", product)
	r.HandleFunc("/auth", auth).Methods("GET")
	r.HandleFunc("/signup", signupView).Methods("post")

	r.HandleFunc("/profile", profile)
	r.HandleFunc("/cart", cart)

	n := negroni.New(negroni.NewLogger())
	n.UseHandler(r)
	if err := http.ListenAndServe(":8888", n); err != nil {
		log.Fatalln(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/home.tmpl",
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
		log.Println(err)
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

func profile(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/base.tmpl",
		"./web/templates/partials/nav.tmpl",
		"./web/templates/pages/profile.tmpl",
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
