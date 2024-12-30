package state

import "time"

type Session struct {
	CSID      string
	Admin     bool
	StartTime time.Time
}
