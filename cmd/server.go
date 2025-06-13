package cmd

import (
	"github.com/graingo/maltose-quickstart/internal/controller/hello"
	"github.com/graingo/maltose/frame/m"
	"github.com/graingo/maltose/net/mhttp"
)

func HttpServer() *mhttp.Server {
	s := m.Server()

	// middleware
	s.Group("/api/v1", func(group *mhttp.RouterGroup) {
		group.Middleware(mhttp.MiddlewareResponse(), mhttp.MiddlewareLog())
		group.Bind(
			hello.NewV1(),
		)
	})

	return s
}
