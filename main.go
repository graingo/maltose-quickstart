package main

import (
	"log"

	"github.com/graingo/maltose-quickstart/cmd"
	"github.com/graingo/maltose/frame/m"
)

func main() {
	// init http server
	httpSvr := cmd.HttpServer()

	// run app
	if err := m.AppRun(httpSvr); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
