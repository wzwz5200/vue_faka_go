package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrderSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    out_trade_no  string `json:"out_trade_no" form:"out_trade_no" `
    Status  string `json:"status" form:"status" `
    request.PageInfo
}
