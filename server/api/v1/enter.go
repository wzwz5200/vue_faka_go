package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/GetKey"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/GetPay"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Goods"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Goods_info"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/KeyUse"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Key_use"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Name"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	GoodsApiGroup      Goods.ApiGroup
	NameApiGroup       Name.ApiGroup
	GetPayApiGroup     GetPay.ApiGroup
	KeyUseApiGroup     KeyUse.ApiGroup
	Key_useApiGroup    Key_use.ApiGroup
	GetKeyApiGroup     GetKey.ApiGroup
	Goods_infoApiGroup Goods_info.ApiGroup
}
