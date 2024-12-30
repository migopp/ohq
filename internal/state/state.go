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

func (s *State) Debug() {
	for idx, se := range s.Queue {
		fmt.Printf("idx: %d, se.CSID: %s\n", idx, se.CSID)
	}
}

var GlobalState State
