// 自动生成模板GetKeys
package GetKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 获取key 结构体  GetKeys
type GetKeys struct {
    global.GVA_MODEL
    GetKey  string `json:"getKey" form:"getKey" gorm:"column:get_key;comment:;"`  //获取key 
}


// TableName 获取key GetKeys自定义表名 get_keys
func (GetKeys) TableName() string {
    return "get_keys"
}

