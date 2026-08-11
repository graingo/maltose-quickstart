// =================================================================================
// Code generated and maintained by Maltose tool. You can edit this file as you like.
// =================================================================================
package hello

import (
	"context"

	"github.com/graingo/maltose-quickstart/api/hello/v2"
)

// Bye is the handler for the Bye API.
func (c *HelloV2) Bye(_ context.Context, req *v2.ByeReq) (res *v2.ByeRes, err error) {
	return &v2.ByeRes{Name: "See you again, " + req.Name + "!"}, nil
}
