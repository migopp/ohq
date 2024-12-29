package state

import (
	"errors"
	"fmt"

	"github.com/migopp/ohq/internal/students"
)

type State struct {
	Queue []students.Student
}

func (s *State) Offer(u students.Student) {
	s.Queue = append(s.Queue, u)
}

func (s *State) Poll() (students.Student, error) {
	if len(s.Queue) == 0 {
		var u students.Student
		return u, errors.New("Attempted `Poll` on empty queue")
	}
	u, r := s.Queue[0], s.Queue[1:]
	s.Queue = r
	return u, nil
}

func (s *State) Debug() {
	for idx, u := range s.Queue {
		fmt.Printf("idx: %d, u.id: %s\n", idx, u.CSID)
	}
}

var GlobalState State
