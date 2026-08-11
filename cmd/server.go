package cmd

import (
	"github.com/graingo/maltose-quickstart/internal/controller/hello"
	_ "github.com/graingo/maltose-quickstart/internal/logic"
	"github.com/graingo/maltose/frame/m"
	"github.com/graingo/maltose/net/mhttp"
)

// HTTPServer builds the application's HTTP server and registers its routes.
func HTTPServer() *mhttp.Server {
	server := m.Server()

	server.Group("/api/v1", func(group *mhttp.RouterGroup) {
		group.Middleware(mhttp.MiddlewareResponse(), mhttp.MiddlewareLog())
		group.Bind(hello.NewV1())
	})
	server.Group("/api/v2", func(group *mhttp.RouterGroup) {
		group.Middleware(mhttp.MiddlewareResponse(), mhttp.MiddlewareLog())
		group.Bind(hello.NewV2())
	})

	return server
}
