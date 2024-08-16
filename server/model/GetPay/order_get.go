// 自动生成模板Order
package GetPay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 订单 结构体  Order
type Order struct {
    global.GVA_MODEL
    out_trade_no  string `json:"out_trade_no" form:"out_trade_no" gorm:"column:out_trade_no;comment:;"`  //订单 
    Status  string `json:"status" form:"status" gorm:"column:status;comment:;"`  //订单状态 
    Key  string `json:"key" form:"key" gorm:"column:key;comment:;"`  //卡密 
}


// TableName 订单 Order自定义表名 order
func (Order) TableName() string {
    return "order"
}

