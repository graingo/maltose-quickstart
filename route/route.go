package route

import (
	"github.com/graingo/maltose-quickstart/internal/controller"
	"github.com/graingo/maltose/net/mhttp"
)

func Build(s *mhttp.Server) {
	s.Use(mhttp.MiddlewareResponse())
	s.Use(mhttp.MiddlewareLog())

	hello := controller.NewHelloController()

	s.Group("api/v1", func(rg *mhttp.RouterGroup) {
		rg.BindObject(hello)
	})
}
