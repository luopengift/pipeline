package pipeline

import (
	"context"
	"encoding/json"
	"path"

	"github.com/luopengift/gohttp"
	"github.com/luopengift/golibs/uuid"
	"github.com/luopengift/log"
)

// PipeLine pipeline instance
type PipeLine struct {
	ID                     string            `json:"id"`
	Name                   string            `json:"name"`
	Events                 string            `json:"events"`
	Labels                 map[string]string `json:"labels"`
	Env                    map[string]string `json:"env"`
	Stages                 []Stager          `json:"stages"`
	gohttp.BaseHTTPHandler `json:"-"`
}

// NewPipeLine new pipeline
func NewPipeLine() *PipeLine {
	return &PipeLine{
		ID: uuid.Rand().Hex(),
	}
}

// Copy get a copy of pipeline
func (p *PipeLine) Copy() *PipeLine {
	return nil
}

// Dump format pipeline
func (p *PipeLine) Dump() {
}

// Run run
func (p *PipeLine) Run(ctx context.Context) error {
	for _, stage := range p.Stages {
		stage.Execute(ctx)
	}
	return nil
}

// AddStage add stage
func (p *PipeLine) AddStage(s Stager) {
	p.Stages = append(p.Stages, s)
}

// GET method
func (p *PipeLine) GET() {
	//query := p.GetQueryArgs()
}

// POST method, New PipeLine createby http
func (p *PipeLine) POST() {
	p.ID = uuid.Rand().Hex()
	if err := json.Unmarshal(p.GetBodyArgs(), p); err != nil {
		log.Error("%v", err)
		return
	}
	if err := store.Put(path.Join("pipeline", p.ID), p); err != nil {
		log.Error("%v", err)
		return
	}
	log.Info("%s", p.ID)
}
