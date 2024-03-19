package routes

import (
	"ecommerce/internal/database"
	"ecommerce/internal/middleware"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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

	// authentication
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.HandleFunc("/auth", auth)	
	authRouter.HandleFunc("/signup", signupController).Methods("POST")
	authRouter.HandleFunc("/login", loginController).Methods("POST")
	authRouter.Use(sessionStore.AuthBlock)

	// protected endpoints
	protectedRouter := r.PathPrefix("/").Subrouter()
	protectedRouter.HandleFunc("/profile", profile)
	protectedRouter.HandleFunc("/cart", cart)
	protectedRouter.Use(sessionStore.EnsureLoggedIn)

	logged := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8888", logged)

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


