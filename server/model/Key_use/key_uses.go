// 自动生成模板Key_uses
package Key_use

import (
	"time"
)

// key表 结构体  Key_uses
type Key_uses struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:64;"`    //id字段
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"` //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"` //updatedAt字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"` //deletedAt字段
	Keys      string     `json:"keys" form:"keys" gorm:"column:keys;comment:;"`                 //keys字段
	Use       *bool      `json:"use" form:"use" gorm:"default:false;column:use;comment:;"`      //是否使用
}

type Key_z struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:64;"`    //id字段
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"` //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"` //updatedAt字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"` //deletedAt字段
	Keys      string
	Use       *bool `json:"use" form:"use" gorm:"default:false;column:use;comment:;"` //是否使用
}

// TableName key表 Key_uses自定义表名 key
func (Key_uses) TableName() string {
	return "key"
}
