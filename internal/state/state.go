package state

import (
	"fmt"

	"github.com/migopp/ohq/internal/users"
)

type State struct {
	Queue []users.User
}

func (s *State) Offer(u users.User) {
	s.Queue = append(s.Queue, u)
}

func (s *State) Debug() {
	for idx, u := range s.Queue {
		fmt.Printf("idx: %d, u.id: %s\n", idx, u.ID)
	}
}

var GlobalState State
