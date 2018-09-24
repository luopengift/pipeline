package api

import (
	"github.com/luopengift/gohttp"
	"github.com/luopengift/pipeline"
)

var hello = func(ctx *gohttp.Context) {
	ctx.Output("hello")
}
type pipeLineHandler struct {
	gohttp.APIHandler
}
func (p *pipeLineHandler) GET() {
	p.Output("get pipeline")
}

func Init() {
	app := gohttp.Init()
	app.RouteFunCtx("/hello", hello)
	app.Route("/pipelines", &pipeline.PipeLine{})
	go app.Run()
}