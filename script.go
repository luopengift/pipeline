package pipeline

import (
	"context"

	"github.com/luopengift/log"
)

// Script script
type Script struct {
	Cmd string `json:"cmd"`
}

// New new
func New() Stager {
	return &Script{
		Cmd: "init cmd",
	}
}

// Execute exec
func (s *Script) Execute(ctx context.Context) {
	log.Info("execute script, %s", s.Cmd)
}
