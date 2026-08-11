// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package service

import (
	"context"

	"github.com/graingo/maltose-quickstart/internal/model"
)

type IHello interface {
	// Hello builds a greeting for the supplied name.
	Hello(ctx context.Context, req *model.HelloInput) (res *model.HelloOutput, err error)
}

var localHello IHello

// Hello returns the registered IHello implementation.
func Hello() IHello {
	if localHello == nil {
		panic("implement not found for interface IHello, forgot register?")
	}
	return localHello
}

// RegisterHello registers an IHello implementation.
func RegisterHello(i IHello) {
	localHello = i
}
