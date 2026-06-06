package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/th3shadowbroker/hymetric/internal/timer"
)

var timerClient *timer.Client

func NewRouter() func(r chi.Router) {
	timerClient = timer.NewClient()

	return func(r chi.Router) {
		r.Get("/readyz", getHealthCheck())
		r.Get("/livez", getHealthCheck())

		r.Get("/timer", getTimer)
	}
}
