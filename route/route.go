package route

import (
	"github.com/graingo/maltose/net/mhttp"
)

func Build(s *mhttp.Server) {
	s.Use(mhttp.MiddlewareResponse())
	s.Use(mhttp.MiddlewareLog())

}
