package timer

import (
	"fmt"
	"time"

	"github.com/th3shadowbroker/hymetric/internal/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Frame struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Text        string `json:"text"`
	Icon        string `json:"icon"`
}

func NewFrame(timer *config.Timer, res *Response) Frame {
	var text string
	if res.Estimate != 0 {
		estimate := time.UnixMilli(res.Estimate)
		text = fmtTimeUntil(estimate)
	} else {
		text = cases.Title(language.English).String(res.Message)
	}

	return Frame{
		Name:        timer.Name,
		DisplayName: timer.DisplayName,
		Text:        text,
		Icon:        timer.Icon,
	}
}

func fmtTimeUntil(t time.Time) string {
	duration := time.Since(t).Abs()

	dur := duration.Round(time.Minute)
	if dur <= 0 {
		return "Now"
	}

	oneDay := time.Hour * 24
	d := dur / oneDay
	dur -= d * oneDay
	h := dur / time.Hour
	dur -= h * time.Hour
	m := dur / time.Minute

	return fmt.Sprintf("%02d:%02d:%02d", d, h, m)
}
