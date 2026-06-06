package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/th3shadowbroker/hymetric/internal/config"
	"github.com/th3shadowbroker/hymetric/internal/timer"
)

func getTimer(w http.ResponseWriter, r *http.Request) {
	timers := make([]timer.Frame, 0)

	for _, t := range config.Active.Timers {
		if chi.URLParam(r, t.Name) == "false" {
			continue
		}

		res, err := timerClient.GetTimer(t)
		if err != nil {
			log.Printf("Failed to get timer %s: %s", t.Name, err)
		}

		timers = append(timers, timer.NewFrame(&t, res))
	}

	body, _ := json.Marshal(map[string][]timer.Frame{
		"frames": timers,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
