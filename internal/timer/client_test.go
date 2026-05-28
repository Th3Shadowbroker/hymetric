package timer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/th3shadowbroker/hymetric/internal/timer"
)

func TestClient(t *testing.T) {
	client := timer.NewClient()
	timers := []timer.Definition{
		{
			Name: "darkauction",
			Url:  "https://hypixel-api.inventivetalent.org/api/skyblock/darkauction/estimate",
		},
	}

	for _, timer := range timers {
		t.Run(timer.Name, func(t *testing.T) {
			_, err := client.GetTimer(timer)
			assert.NoError(t, err)
		})
	}
}
