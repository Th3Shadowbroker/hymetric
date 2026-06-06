package v1

import (
	"net/http"
)

func getHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UP"))
	}
}
