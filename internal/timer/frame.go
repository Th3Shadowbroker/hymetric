package timer

import (
	"fmt"
	"time"

	"github.com/th3shadowbroker/hymetric/internal/config"
)

type Frame struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Text        string `json:"text"`
	Icon        string `json:"icon"`
}

func NewFrame(timer *config.Timer, res *Response) Frame {
	estimate := time.UnixMilli(res.Estimate)

	return Frame{
		Name:        timer.Name,
		DisplayName: timer.DisplayName,
		Text:        fmtTimeUntil(estimate),
		Icon:        timer.Icon,
	}
}

func fmtTimeUntil(t time.Time) string {
	duration := time.Since(t).Abs()

	d := duration.Round(time.Minute)
	if d <= 0 {
		return "Now"
	}

	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	return fmt.Sprintf("%02d:%02d", h, m)
}
