package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/GetKey"
	"github.com/flipped-aurora/gin-vue-admin/server/router/GetPay"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Goods"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Goods_info"
	"github.com/flipped-aurora/gin-vue-admin/server/router/KeyUse"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Key_use"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Name"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Goods      Goods.RouterGroup
	Name       Name.RouterGroup
	GetPay     GetPay.RouterGroup
	KeyUse     KeyUse.RouterGroup
	Key_use    Key_use.RouterGroup
	GetKey     GetKey.RouterGroup
	Goods_info Goods_info.RouterGroup
}
