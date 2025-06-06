// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package service

type sA struct{}

var localA = NewA()

// NewA creates a new service instance.
func NewA() *sA {
	return &sA{}
}

// A returns the default service instance.
func A() *sA {
	return localA
}
