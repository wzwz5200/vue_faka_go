// 自动生成模板Key
package KeyUse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 卡密 结构体  Key
type Key struct {
    global.GVA_MODEL
    Keys  string `json:"keys" form:"keys" gorm:"column:keys;comment:;"`  //卡密 
}


// TableName 卡密 Key自定义表名 key
func (Key) TableName() string {
    return "key"
}

