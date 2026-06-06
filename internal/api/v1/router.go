package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/th3shadowbroker/hymetric/internal/timer"
)

var timerClient *timer.Client

func NewRouter() func(r chi.Router) {
	timerClient = timer.NewClient()

	return func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong!"))
		})

		r.Get("/timers", getTimers)
	}
}
