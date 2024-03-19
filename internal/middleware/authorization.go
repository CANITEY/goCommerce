package middleware

import (
	"ecommerce/api/models"
	"net/http"

	"github.com/gorilla/sessions"
)

type Store struct {
	S sessions.Store
}

func (s *Store)EnsureLoggedIn(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.S.Get(r, "session")
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}

		if  len(session.Values) == 0 {
			http.Redirect(w, r, "/auth", http.StatusFound)
			return
		}
		if value, ok := session.Values["user"].(*models.User); ok && value != nil{
			next.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, "/auth", http.StatusFound)
	})
}

func (s *Store)AuthBlock(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.S.Get(r, "session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if len(session.Values) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		if value, ok := session.Values["user"].(*models.User); ok && value != nil {
			http.Redirect(w, r, "/profile", http.StatusFound)
		}
	})
}
