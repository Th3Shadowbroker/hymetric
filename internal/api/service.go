package api

import (
	"net/http"
	"time"

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

func createServer(address string, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.Protocols = new(http.Protocols)
	server.Protocols.SetHTTP1(true)
	server.Protocols.SetUnencryptedHTTP2(true)

	return server
}

func Listen(address string) error {
	return createServer(address, createRootRouter()).ListenAndServe()
}
