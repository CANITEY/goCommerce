package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Store struct {
	S sessions.Store
}

func (s *Store)ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, err := s.S.Get(r, "session")
	if err != nil {
		http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
		return
	}


	if err != nil {
		http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
		return
	}

	next(w, r)
}

func IsNotAuthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}
