package KeyUse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KeyRouter struct {}

// InitKeyRouter 初始化 卡密 路由信息
func (s *KeyRouter) InitKeyRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	KEYSRouter := Router.Group("KEYS").Use(middleware.OperationRecord())
	KEYSRouterWithoutRecord := Router.Group("KEYS")
	KEYSRouterWithoutAuth := PublicRouter.Group("KEYS")
	{
		KEYSRouter.POST("createKey", KEYSApi.CreateKey)   // 新建卡密
		KEYSRouter.DELETE("deleteKey", KEYSApi.DeleteKey) // 删除卡密
		KEYSRouter.DELETE("deleteKeyByIds", KEYSApi.DeleteKeyByIds) // 批量删除卡密
		KEYSRouter.PUT("updateKey", KEYSApi.UpdateKey)    // 更新卡密
	}
	{
		KEYSRouterWithoutRecord.GET("findKey", KEYSApi.FindKey)        // 根据ID获取卡密
		KEYSRouterWithoutRecord.GET("getKeyList", KEYSApi.GetKeyList)  // 获取卡密列表
	}
	{
	    KEYSRouterWithoutAuth.GET("getKeyPublic", KEYSApi.GetKeyPublic)  // 获取卡密列表
	}
}
