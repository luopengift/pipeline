package main

import (
	"github.com/luopengift/pipeline"
	"github.com/luopengift/pipeline/api"
	"github.com/luopengift/log"

)

func main() {
	pipeline.InitStore("./data")
	log.Info("pipeline start....")
	api.Init()
	select {}
}