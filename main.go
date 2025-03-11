package main

import (
	"github.com/graingo/maltose-quickstart/route"

	"github.com/graingo/maltose/frame/m"
)

func main() {
	s := m.Server()
	route.Build(s)
	s.Run()
}
