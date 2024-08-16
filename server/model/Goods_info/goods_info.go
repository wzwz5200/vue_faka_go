// 自动生成模板Goods_info
package Goods_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 获取商品信息 结构体  Goods_info
type Goods_info struct {
    global.GVA_MODEL
    Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`  //商品名字 
    Price  string `json:"price" form:"price" gorm:"column:price;comment:;"`  //价格 
    Bio  string `json:"bio" form:"bio" gorm:"column:bio;comment:;"`  //商品介绍 
    Goods_id  string `json:"goods_id" form:"goods_id" gorm:"column:goods_id;comment:;"`  //商品id 
    Img  string `json:"img" form:"img" gorm:"column:img;comment:;"`  //商品图 
}


// TableName 获取商品信息 Goods_info自定义表名 goods_info
func (Goods_info) TableName() string {
    return "goods_info"
}

