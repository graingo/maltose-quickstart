// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package v1

import (
	"context"
	"github.com/graingo/maltose-quickstart/api/v1"
)

type cA struct{}

// NewA creates a new controller.
func NewA() *cA {
	return &cA{}
}

// A is the handler for the A API.
func (c *cA) A(ctx context.Context, req *v1.AReq) (res *v1.ARes, err error) {
	// TODO: Implement the business logic here.
	panic("implement me")
}

// D is the handler for the D API.
func (c *cA) D(ctx context.Context, req *v1.DReq) (res *v1.DRes, err error) {
	// TODO: Implement the business logic here.
	panic("implement me")
}
