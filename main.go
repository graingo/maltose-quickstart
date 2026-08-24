package main

import (
	"log"

	"github.com/graingo/maltose-quickstart/cmd"
	"github.com/graingo/maltose/frame/m"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}

func run() error {
	logger := m.Log()
	server := cmd.HTTPServer()

	app := m.NewApp(
		m.WithLogger(logger),
		m.WithServer(server),
	)
	return app.Run()
}
