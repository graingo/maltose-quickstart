package v1

import "github.com/graingo/maltose/frame/m"

// HelloReq contains the input for the hello endpoint.
type HelloReq struct {
	m.Meta `group:"/api/v1" envelope:"maltose" method:"GET" path:"/hello" summary:"问候示例" tag:"公共服务"`
	Name   string `query:"name" dc:"姓名" binding:"required"`
}

// HelloRes contains the greeting returned by the hello endpoint.
type HelloRes struct {
	Name string   `json:"name" dc:"姓名"`
	Tags []string `json:"tags" field:"required" dc:"标签，空集合返回空数组"`
}

// ByeReq contains the input for the goodbye endpoint.
type ByeReq struct {
	m.Meta `group:"/api/v1" envelope:"maltose" method:"POST" path:"/bye" summary:"告别示例" tag:"公共服务"`
	Name   string `json:"name" dc:"姓名" binding:"required"`
}

// ByeRes contains the message returned by the goodbye endpoint.
type ByeRes struct {
	Name string `json:"name" dc:"姓名"`
}
