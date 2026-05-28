package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/th3shadowbroker/hymetric/internal/api/v1"
)

func createRootRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	router.Route("/api/v1", v1.NewRouter())

	return router
}

func Listen(address string) error {
	return http.ListenAndServe(address, createRootRouter())
}
