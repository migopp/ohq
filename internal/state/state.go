package state

import (
	"errors"
	"fmt"

	"github.com/migopp/ohq/internal/users"
)

type State struct {
	Queue []users.User
}

func (s *State) Offer(u users.User) {
	s.Queue = append(s.Queue, u)
}

func (s *State) Poll() (users.User, error) {
	if len(s.Queue) == 0 {
		var u users.User
		return u, errors.New("Attempted `Poll` on empty queue")
	}
	u, r := s.Queue[0], s.Queue[1:]
	s.Queue = r
	return u, nil
}

func (s *State) Debug() {
	for idx, u := range s.Queue {
		fmt.Printf("idx: %d, u.id: %s\n", idx, u.ID)
	}
}

var GlobalState State
