package templates

import (
	"github.com/migopp/ohq/internal/users"
)

type QueueContent struct {
	Users []users.User
}

type ErrContent struct {
	Err error
}
