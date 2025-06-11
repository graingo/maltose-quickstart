package cmd

import (
	"github.com/graingo/maltose/frame/m"
)

func Server() {
	s := m.Server()
	route(s)
	s.Run()
}
