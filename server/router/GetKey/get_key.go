package GetKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GetKeysRouter struct{}

func (s *GetKeysRouter) InitGetKeysRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	GKeyRouter := Router.Group("GKey").Use(middleware.OperationRecord())
	GKeyRouterWithoutRecord := Router.Group("GKey")
	GKeyRouterWithoutAuth := PublicRouter.Group("GKey")
	{
		GKeyRouter.POST("createGetKeys", GKeyApi.CreateGetKeys)
		GKeyRouter.DELETE("deleteGetKeys", GKeyApi.DeleteGetKeys)
		GKeyRouter.DELETE("deleteGetKeysByIds", GKeyApi.DeleteGetKeysByIds)
		GKeyRouter.PUT("updateGetKeys", GKeyApi.UpdateGetKeys)
	}
	{
		GKeyRouterWithoutRecord.GET("findGetKeys", GKeyApi.FindGetKeys)
		GKeyRouterWithoutRecord.GET("getGetKeysList", GKeyApi.GetGetKeysList)
	}
	{
		GKeyRouterWithoutAuth.GET("getGetKeysPublic", GKeyApi.GetGetKeysPublic)
		GKeyRouterWithoutAuth.
			POST("key", GKeyApi.GetKey_psot)

	}
}
