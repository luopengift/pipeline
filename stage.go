package pipeline

import (
	"context"
	"github.com/luopengift/log"
)
// Stager 
type Stager interface {
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
