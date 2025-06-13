package cmd

import (
	"github.com/graingo/maltose/net/mhttp"
)

func route(s *mhttp.Server) {
	s.Use(mhttp.MiddlewareResponse())
	s.Use(mhttp.MiddlewareLog())

}
