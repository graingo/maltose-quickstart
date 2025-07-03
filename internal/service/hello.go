// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package service

import (
	"context"

	"github.com/graingo/maltose-quickstart/internal/model"
)

type IHello interface {
	// TODO: Define your service interface methods here.
	Hello(ctx context.Context, req *model.HelloInput) (res *model.HelloOutput, err error)
}

var localHello IHello

// Hello returns the registered implementation of IHello.
// It panics if no implementation is registered.
func Hello() IHello {
	if localHello == nil {
		panic("implement not found for interface IHello, forgot register?")
	}
	return localHello
}

// RegisterHello registers an implementation for the IHello interface.
func RegisterHello(i IHello) {
	localHello = i
}
