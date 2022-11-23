package timertrack

import (
	log "github.com/chris-dot-exe/AwesomeLog"
	"time"
)

type Timetracker struct {
	start time.Time
}

func NewTimetracker() *Timetracker {
	return &Timetracker{
		start: time.Now(),
	}
}

func (tt *Timetracker) Track(info string) {
	endTime := time.Now()
	log.Println(info, ": ", endTime.Sub(tt.start))
}
