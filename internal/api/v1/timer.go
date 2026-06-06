package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/th3shadowbroker/hymetric/internal/comparison"
	"github.com/th3shadowbroker/hymetric/internal/config"
	"github.com/th3shadowbroker/hymetric/internal/timer"
)

func getTimer(w http.ResponseWriter, r *http.Request) {
	timers := make([]timer.Frame, 0)
	query := r.URL.Query()

	defaults := config.Active.Defaults
	prefix := comparison.IfThenElse(query.Has("prefix"), query.Get("prefix") == "true", defaults.Prefix)
	activeText := comparison.IfThenElse(query.Has("activeText"), query.Get("activeText"), defaults.ActiveText)

	for _, t := range config.Active.Timers {
		if query.Get(t.Name) == "false" {
			continue
		}

		res, err := timerClient.GetTimer(t)
		if err != nil {
			log.Printf("Failed to get timer %s: %s", t.Name, err)
		}

		timers = append(timers, timer.NewFrame(&t, res, prefix, activeText))
	}

	body, _ := json.Marshal(map[string][]timer.Frame{
		"frames": timers,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
