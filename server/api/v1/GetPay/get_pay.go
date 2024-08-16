package GetPay

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/GetPay"
	GetPayReq "github.com/flipped-aurora/gin-vue-admin/server/model/GetPay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Get_payApi struct{}

// CreateGet_pay 创建订单
// @Tags Get_pay
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetPay.Get_pay true "创建订单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Pay/createGet_pay [post]
func (PayApi *Get_payApi) CreateGet_pay(c *gin.Context) {
	var Pay GetPay.Get_pay
	err := c.ShouldBindJSON(&Pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PayService.CreateGet_pay(&Pay)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteGet_pay 删除订单
// @Tags Get_pay
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetPay.Get_pay true "删除订单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Pay/deleteGet_pay [delete]
func (PayApi *Get_payApi) DeleteGet_pay(c *gin.Context) {
	ID := c.Query("ID")
	err := PayService.DeleteGet_pay(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGet_payByIds 批量删除订单
// @Tags Get_pay
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Pay/deleteGet_payByIds [delete]
func (PayApi *Get_payApi) DeleteGet_payByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := PayService.DeleteGet_payByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGet_pay 更新订单
// @Tags Get_pay
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetPay.Get_pay true "更新订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Pay/updateGet_pay [put]
func (PayApi *Get_payApi) UpdateGet_pay(c *gin.Context) {
	var Pay GetPay.Get_pay
	err := c.ShouldBindJSON(&Pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PayService.UpdateGet_pay(Pay)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGet_pay 用id查询订单
// @Tags Get_pay
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query GetPay.Get_pay true "用id查询订单"
// @Success 200 {object} response.Response{data=GetPay.Get_pay,msg=string} "查询成功"
// @Router /Pay/findGet_pay [get]
func (PayApi *Get_payApi) FindGet_pay(c *gin.Context) {
	ID := c.Query("ID")
	rePay, err := PayService.GetGet_pay(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rePay, c)
}

// GetGet_payList 分页获取订单列表
// @Tags Get_pay
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "分页获取订单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Pay/getGet_payList [get]
func (PayApi *Get_payApi) GetGet_payList(c *gin.Context) {
	var pageInfo GetPayReq.Get_paySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PayService.GetGet_payInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetGet_payPublic 不需要鉴权的订单接口
// @Tags Get_pay
// @Summary 不需要鉴权的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "分页获取订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Pay/getGet_payPublic [get]
func (PayApi *Get_payApi) GetGet_payPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订单接口信息",
	}, "获取成功", c)
}

// GetOrder 等待开发的的订单接口
// @Tags Get_pay
// @Summary 等待开发的的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pay/order [GET]

type adder struct {
	U_adder string
}

func (PayApi *Get_payApi) GetOrder(c *gin.Context) {
	// 请添加自己的业务逻辑

	var form adder
	c.Bind(&form)
	fmt.Println(form.U_adder)
	if form.U_adder == "" {
		response.Fail(c)
		return

	}

	order := PayService.GetOrder(form.U_adder)

	response.OkWithData(order, c)
}

type Payz struct {
	Order string
	Price string
}

type P_t struct {
	Price string
}

// AlilPay 等待开发的的订单接口
// @Tags Get_pay
// @Summary 等待开发的的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pay/alilpay [POST]
func (PayApi *Get_payApi) AlilPay(c *gin.Context) {
	// 请添加自己的业务逻辑
	var form Payz
	var fz P_t

	c.Bind(&form)
	fmt.Println(form.Order)
	fmt.Println(form.Price)
	fmt.Println(fz.Price)
	int, err := strconv.ParseFloat(form.Price, 64)
	if err != nil {
		fmt.Println("转int错误", err)

	}

	// fmt.Println(int)
	url := PayService.AlilPay(form.Order, int)

	response.OkWithData(url, c)
}

// func (PayApi *Get_payApi) AlilPay(c *gin.Context) {
// 	// 请添加自己的业务逻辑
// 	var form Login
// 	c.ShouldBind(&form)
// 	fmt.Println(form.Order)
// 	url := PayService.AlilPay(form.Order)

//		response.OkWithData(url, c)
//	}
//
// Notify 等待开发的的订单接口
// @Tags Get_pay
// @Summary 等待开发的的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pay/notify [POST]
func (PayApi *Get_payApi) Notify(c *gin.Context) {
	// 请添加自己的业务逻辑
	err := PayService.Notify(c)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// Stock 等待开发的的订单接口
// @Tags Get_pay
// @Summary 等待开发的的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pay/stock [POST]

type Keys struct {
	Orade string
}

func (PayApi *Get_payApi) Stock(c *gin.Context) {
	// 请添加自己的业务逻辑

	var form Keys
	c.ShouldBind(&form)
	fmt.Println(form.Orade)
	stock := PayService.Stock(form.Orade)

	response.OkWithData(stock, c)
}

// Query_order 等待开发的的订单接口
// @Tags Get_pay
// @Summary 等待开发的的订单接口
// @accept application/json
// @Produce application/json
// @Param data query GetPayReq.Get_paySearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Pay/qorder [POST]

type adders struct {
	U_add string
}
type re struct {
	Key string

	vaa []string
}

func (PayApi *Get_payApi) Query_order(c *gin.Context) {
	// 请添加自己的业务逻辑

	var form adder
	c.ShouldBind(&form)

	key, err := PayService.Query_order(c, form.U_adder)

	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}

	response.OkWithData(key, c)
}
