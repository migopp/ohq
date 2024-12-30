package state

import (
	"errors"
	"fmt"
	"time"
)

type State struct {
	Queue []Entry
}

func (s *State) Offer(e Entry) {
	if len(s.Queue) == 0 {
		e.StartTime = time.Now()
	}
	s.Queue = append(s.Queue, e)
}

func (s *State) Poll() (Entry, error) {
	if len(s.Queue) == 0 {
		var e Entry
		return e, errors.New("Attempted `Poll` on empty queue")
	}
	e, r := s.Queue[0], s.Queue[1:]
	s.Queue = r
	if len(s.Queue) != 0 {
		s.Queue[0].StartTime = time.Now()
	}
	return e, nil
}

func (s *State) Debug() {
	for idx, e := range s.Queue {
		fmt.Printf("idx: %d, e.CSID: %s\n", idx, e.CSID)
	}
}

var GlobalState State
