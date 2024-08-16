// 自动生成模板Get_pay
package GetPay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 订单 结构体  Get_pay
type Get_pay struct {
	global.GVA_MODEL
	Out_trade_no string `json:"out_trade_no" form:"out_trade_no" gorm:"column:out_trade_no;comment:;"` //订单号
	Status       string `json:"status" form:"status" gorm:"column:status;comment:;"`                   //订单状态
	Key          string `json:"key" form:"key" gorm:"column:key;comment:;"`                            //卡密
	U_add        string `json:"u_add" form:"u_add" gorm:"column:u_add;comment:;"`
}

// TableName 订单 Get_pay自定义表名 get_pay
func (Get_pay) TableName() string {
	return "get_pay"
}
