package Key_use

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Key_usesRouter struct{}

func (s *Key_usesRouter) InitKey_usesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	Key_usessRouter := Router.Group("Key_usess").Use(middleware.OperationRecord())
	Key_usessRouterWithoutRecord := Router.Group("Key_usess")
	Key_usessRouterWithoutAuth := PublicRouter.Group("Key_usess")
	{
		Key_usessRouter.POST("createKey_uses", Key_usessApi.CreateKey_uses)
		Key_usessRouter.DELETE("deleteKey_uses", Key_usessApi.DeleteKey_uses)
		Key_usessRouter.DELETE("deleteKey_usesByIds", Key_usessApi.DeleteKey_usesByIds)
		Key_usessRouter.PUT("updateKey_uses", Key_usessApi.UpdateKey_uses)
	}
	{
		Key_usessRouterWithoutRecord.GET("findKey_uses", Key_usessApi.FindKey_uses)
		Key_usessRouterWithoutRecord.GET("getKey_usesList", Key_usessApi.GetKey_usesList)
	}
	{
		Key_usessRouterWithoutAuth.GET("getKey_usesPublic", Key_usessApi.GetKey_usesPublic)
		Key_usessRouterWithoutAuth.
			POST("Getkey_p", Key_usessApi.Getkeys)

	}
}
