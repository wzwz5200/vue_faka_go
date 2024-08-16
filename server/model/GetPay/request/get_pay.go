package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Get_paySearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Out_trade_no  string `json:"out_trade_no" form:"out_trade_no" `
    Key  string `json:"key" form:"key" `
    request.PageInfo
}
