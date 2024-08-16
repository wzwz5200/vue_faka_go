package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/GetKey"
	"github.com/flipped-aurora/gin-vue-admin/server/service/GetPay"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Goods"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Goods_info"
	"github.com/flipped-aurora/gin-vue-admin/server/service/KeyUse"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Key_use"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Name"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	GoodsServiceGroup      Goods.ServiceGroup
	NameServiceGroup       Name.ServiceGroup
	GetPayServiceGroup     GetPay.ServiceGroup
	KeyUseServiceGroup     KeyUse.ServiceGroup
	Key_useServiceGroup    Key_use.ServiceGroup
	GetKeyServiceGroup     GetKey.ServiceGroup
	Goods_infoServiceGroup Goods_info.ServiceGroup
}
