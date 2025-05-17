package cmd

import (
	"github.com/graingo/maltose-quickstart/route"

	"github.com/graingo/maltose/frame/m"
)

func Server() {
	s := m.Server()
	route.Build(s)
	s.Run()
}
