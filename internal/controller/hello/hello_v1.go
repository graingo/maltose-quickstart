// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package hello

import (
	"context"

	"github.com/graingo/maltose-quickstart/api/hello/v1"
	"github.com/graingo/maltose-quickstart/internal/model"
	"github.com/graingo/maltose-quickstart/internal/service"
)

// Hello is the handler for the Hello API.
func (c *HelloV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	output, err := service.Hello().Hello(ctx, &model.HelloInput{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloRes{Name: output.Greeting, Tags: []string{}}, nil
}

// Bye is the handler for the Bye API.
func (c *HelloV1) Bye(_ context.Context, req *v1.ByeReq) (res *v1.ByeRes, err error) {
	return &v1.ByeRes{Name: "Goodbye, " + req.Name + "!"}, nil
}
