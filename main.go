package main

import (
	"log"

	"github.com/graingo/maltose-quickstart/cmd"
	"github.com/graingo/maltose/frame/m"
)

func main() {
	// Build the HTTP server before handing its lifecycle to the application.
	httpServer := cmd.HTTPServer()

	if err := m.NewApp(m.WithServer(httpServer)).Run(); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
