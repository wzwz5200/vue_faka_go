package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
	}
	{
		GetPayRouter := router.RouterGroupApp.GetPay
		GetPayRouter.InitGet_payRouter(privateGroup, publicGroup)
	}
	{
		KeyUseRouter := router.RouterGroupApp.KeyUse
		KeyUseRouter.InitKeyRouter(privateGroup, publicGroup)
	}
	{
		Key_useRouter := router.RouterGroupApp.Key_use
		Key_useRouter.InitKey_usesRouter(privateGroup, publicGroup)
	}
	{
		GetKeyRouter := router.RouterGroupApp.GetKey
		GetKeyRouter.InitGetKeysRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		Goods_infoRouter := router.RouterGroupApp.Goods_info
		Goods_infoRouter.InitGoods_infoRouter(privateGroup, publicGroup)
	}
}
