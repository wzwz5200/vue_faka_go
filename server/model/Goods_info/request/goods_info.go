package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Goods_infoSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Name  string `json:"name" form:"name" `
    Price  string `json:"price" form:"price" `
    Goods_id  string `json:"goods_id" form:"goods_id" `
    request.PageInfo
}
