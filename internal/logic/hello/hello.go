// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package hello

import (
	"context"

	"github.com/graingo/maltose-quickstart/internal/model"
	"github.com/graingo/maltose-quickstart/internal/service"
)

func init() {
	service.RegisterHello(New())
}

type sHello struct{}

// New creates a new service logic implementation.
func New() service.IHello {
	return &sHello{}
}

func (s *sHello) Hello(_ context.Context, input *model.HelloInput) (output *model.HelloOutput, err error) {
	return &model.HelloOutput{Greeting: "Hello, " + input.Name + "!"}, nil
}
