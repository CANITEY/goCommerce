package routes

import (
	"ecommerce/internal/database"
	"ecommerce/internal/middleware"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
)

var (
	d   database.DB
	err error
	Store = sessions.NewCookieStore([]byte("VERYSECUREKEY"))
)

func init() {
	d, err = database.NewConn()
	if err != nil {
		panic(err)
	}

	sessionStore := middleware.Store{
		S: Store,
	}
	r := mux.NewRouter()

	// adding static fileServer
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("web/static/js"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("web/static/css"))))

	// public endpoints
	r.HandleFunc("/", home)
	r.HandleFunc("/products", products)
	r.HandleFunc("/products/{id}", product)
	r.HandleFunc("/auth", auth).Methods("GET")
	r.HandleFunc("/signup", signupController).Methods("POST")
	r.HandleFunc("/login", loginController).Methods("POST")

	// protected endpoints
	protectedRouter := mux.NewRouter()
	protectedRouter.HandleFunc("/profile", profile)
	protectedRouter.HandleFunc("/cart", cart)
	protectedMiddle := negroni.New()
	protectedMiddle.Use(&sessionStore)
	protectedMiddle.UseHandler(protectedRouter)


	r.PathPrefix("/").Handler(protectedMiddle)


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
