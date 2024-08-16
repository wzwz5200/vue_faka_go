package Goods_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Goods_infoRouter struct{}

func (s *Goods_infoRouter) InitGoods_infoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	GinfoRouter := Router.Group("Ginfo").Use(middleware.OperationRecord())
	GinfoRouterWithoutRecord := Router.Group("Ginfo")
	GinfoRouterWithoutAuth := PublicRouter.Group("Ginfo")
	{
		GinfoRouter.POST("createGoods_info", GinfoApi.CreateGoods_info)
		GinfoRouter.DELETE("deleteGoods_info", GinfoApi.DeleteGoods_info)
		GinfoRouter.DELETE("deleteGoods_infoByIds", GinfoApi.DeleteGoods_infoByIds)
		GinfoRouter.PUT("updateGoods_info", GinfoApi.UpdateGoods_info)
	}
	{
		GinfoRouterWithoutRecord.GET("findGoods_info", GinfoApi.FindGoods_info)
		GinfoRouterWithoutRecord.GET("getGoods_infoList", GinfoApi.GetGoods_infoList)
	}
	{
		GinfoRouterWithoutAuth.GET("getGoods_infoPublic", GinfoApi.GetGoods_infoPublic)
		GinfoRouterWithoutAuth.
			GET("goodsinfo",
				GinfoApi.GetGoods_info)
		GinfoRouterWithoutAuth.
			POST("lnt",
				GinfoApi.Goods_Int)

	}
}
