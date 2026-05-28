package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong!"))
		})
	}
}
