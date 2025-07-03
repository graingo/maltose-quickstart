package v2

import "github.com/graingo/maltose/frame/m"

type ByeReq struct {
	m.Meta `method:"POST" path:"/bye" summary:"Bye请求" tag:"公共服务" dc:"BBBBBBCCCCCCC"`
	Name   string `json:"name" dc:"姓名" binding:"required"`
}

type ByeRes struct {
	Name string `json:"name" dc:"姓名"`
}
