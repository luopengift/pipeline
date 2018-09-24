package pipeline

import (
	"context"
	"github.com/luopengift/log"
)
// Stager 
type Stager interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
	Execute(context.Context)
}


type Script struct {
	Cmd	string	`json:"cmd"`
}

func New() Stager {
	return &Script{
		Cmd: "init cmd",
	}
}

func (s *Script) Execute(ctx context.Context) {
	log.Info("execute script, %s", s.Cmd)
}

func (s *Script) MarshalJSON() ([]byte, error) {
	return []byte("testing..."), nil
}
func (s *Script) UnmarshalJSON([]byte) error {
	return nil
}
