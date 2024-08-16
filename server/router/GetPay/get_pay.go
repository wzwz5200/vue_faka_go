package GetPay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Get_payRouter struct{}

func (s *Get_payRouter) InitGet_payRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PayRouter := Router.Group("Pay").Use(middleware.OperationRecord())
	PayRouterWithoutRecord := Router.Group("Pay")
	PayRouterWithoutAuth := PublicRouter.Group("Pay")
	{
		PayRouter.POST("createGet_pay", PayApi.CreateGet_pay)
		PayRouter.DELETE("deleteGet_pay", PayApi.DeleteGet_pay)
		PayRouter.DELETE("deleteGet_payByIds", PayApi.DeleteGet_payByIds)
		PayRouter.PUT("updateGet_pay", PayApi.UpdateGet_pay)
	}
	{
		PayRouterWithoutRecord.GET("findGet_pay", PayApi.FindGet_pay)
		PayRouterWithoutRecord.GET("getGet_payList", PayApi.GetGet_payList)
	}
	{
		PayRouterWithoutAuth.GET("getGet_payPublic", PayApi.GetGet_payPublic)
		PayRouterWithoutAuth.
			GET(
				"order", PayApi.GetOrder)
		PayRouterWithoutAuth.
			POST("alilpay", PayApi.AlilPay)
		PayRouterWithoutAuth.
			POST("notify", PayApi.Notify)
		PayRouterWithoutAuth.
			POST("stock", PayApi.Stock)
		PayRouterWithoutAuth.
			POST("qorder", PayApi.Query_order)

	}
}
