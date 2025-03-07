package route

import (
	"github.com/graingo/maltose-quickstart/controller"
	"github.com/graingo/maltose/net/mhttp"
)

func Build(s *mhttp.Server) {
	s.Use(mhttp.MiddlewareResponse())
	s.Use(mhttp.MiddlewareLog())

	hello := controller.NewHelloController()

	g := s.Group("api/v1")
	g.Bind(hello)
}
