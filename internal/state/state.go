package state

import (
	"errors"
	"fmt"
	"time"
)

type State struct {
	Queue []Session
}

func (s *State) Offer(se Session) {
	if len(s.Queue) == 0 {
		se.StartTime = time.Now()
	}
	s.Queue = append(s.Queue, se)
}

func (s *State) Poll() (Session, error) {
	if len(s.Queue) == 0 {
		var se Session
		return se, errors.New("Attempted `Poll` on empty queue")
	}
	se, r := s.Queue[0], s.Queue[1:]
	s.Queue = r
	if len(s.Queue) != 0 {
		s.Queue[0].StartTime = time.Now()
	}
	return se, nil
}

func (s *State) OnQueue(se Session) bool {
	for _, qse := range s.Queue {
		if se.CSID == qse.CSID {
			return true
		}
	}
	return false
}

func (s *State) TopTime() string {
	if len(s.Queue) == 0 {
		return ""
	}
	rsecs := uint64(time.Since(s.Queue[0].StartTime).Seconds())
	mins := rsecs / 60
	secs := rsecs % 60
	if mins >= 1 {
		return fmt.Sprintf("%dm %ds", mins, secs)
	} else {
		return fmt.Sprintf("%ds", secs)
	}
}

func (s *State) Debug() {
	for idx, se := range s.Queue {
		fmt.Printf("idx: %d, se.CSID: %s\n", idx, se.CSID)
	}
}

var GlobalState State
