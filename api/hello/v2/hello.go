package v2

import "github.com/graingo/maltose/frame/m"

// ByeReq contains the input for the version 2 goodbye endpoint.
type ByeReq struct {
	m.Meta `group:"/api/v2" envelope:"maltose" method:"POST" path:"/bye" summary:"告别示例 V2" tag:"公共服务"`
	Name   string `json:"name" dc:"姓名" binding:"required"`
}

// ByeRes contains the version 2 goodbye message.
type ByeRes struct {
	Name string `json:"name" dc:"姓名"`
}
