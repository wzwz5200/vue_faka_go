package Goods_info

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Goods_info"
	Goods_infoReq "github.com/flipped-aurora/gin-vue-admin/server/model/Goods_info/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Goods_infoApi struct{}

// CreateGoods_info 创建获取商品信息
// @Tags Goods_info
// @Summary 创建获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Goods_info.Goods_info true "创建获取商品信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Ginfo/createGoods_info [post]
func (GinfoApi *Goods_infoApi) CreateGoods_info(c *gin.Context) {
	var Ginfo Goods_info.Goods_info
	err := c.ShouldBindJSON(&Ginfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = GinfoService.CreateGoods_info(&Ginfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteGoods_info 删除获取商品信息
// @Tags Goods_info
// @Summary 删除获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Goods_info.Goods_info true "删除获取商品信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Ginfo/deleteGoods_info [delete]
func (GinfoApi *Goods_infoApi) DeleteGoods_info(c *gin.Context) {
	ID := c.Query("ID")
	err := GinfoService.DeleteGoods_info(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGoods_infoByIds 批量删除获取商品信息
// @Tags Goods_info
// @Summary 批量删除获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Ginfo/deleteGoods_infoByIds [delete]
func (GinfoApi *Goods_infoApi) DeleteGoods_infoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := GinfoService.DeleteGoods_infoByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGoods_info 更新获取商品信息
// @Tags Goods_info
// @Summary 更新获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Goods_info.Goods_info true "更新获取商品信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Ginfo/updateGoods_info [put]
func (GinfoApi *Goods_infoApi) UpdateGoods_info(c *gin.Context) {
	var Ginfo Goods_info.Goods_info
	err := c.ShouldBindJSON(&Ginfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = GinfoService.UpdateGoods_info(Ginfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGoods_info 用id查询获取商品信息
// @Tags Goods_info
// @Summary 用id查询获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Goods_info.Goods_info true "用id查询获取商品信息"
// @Success 200 {object} response.Response{data=Goods_info.Goods_info,msg=string} "查询成功"
// @Router /Ginfo/findGoods_info [get]
func (GinfoApi *Goods_infoApi) FindGoods_info(c *gin.Context) {
	ID := c.Query("ID")
	reGinfo, err := GinfoService.GetGoods_info(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reGinfo, c)
}

// GetGoods_infoList 分页获取获取商品信息列表
// @Tags Goods_info
// @Summary 分页获取获取商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Goods_infoReq.Goods_infoSearch true "分页获取获取商品信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Ginfo/getGoods_infoList [get]
func (GinfoApi *Goods_infoApi) GetGoods_infoList(c *gin.Context) {
	var pageInfo Goods_infoReq.Goods_infoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := GinfoService.GetGoods_infoInfoList(pageInfo)
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

// GetGoods_infoPublic 不需要鉴权的获取商品信息接口
// @Tags Goods_info
// @Summary 不需要鉴权的获取商品信息接口
// @accept application/json
// @Produce application/json
// @Param data query Goods_infoReq.Goods_infoSearch true "分页获取获取商品信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Ginfo/getGoods_infoPublic [get]
func (GinfoApi *Goods_infoApi) GetGoods_infoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的获取商品信息接口信息",
	}, "获取成功", c)
}

// GetGoods_info 等待开发的的获取商品信息接口
// @Tags Goods_info
// @Summary 等待开发的的获取商品信息接口
// @accept application/json
// @Produce application/json
// @Param data query Goods_infoReq.Goods_infoSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Ginfo/goodsinfo [GET]
func (GinfoApi *Goods_infoApi) GetGoods_info(c *gin.Context) {
	// 请添加自己的业务逻辑
	Goods_info, err := GinfoService.GetGoods()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	c.JSON(http.StatusOK, Goods_info)

}

// Goods_Int 等待开发的的获取商品信息接口
// @Tags Goods_info
// @Summary 等待开发的的获取商品信息接口
// @accept application/json
// @Produce application/json
// @Param data query Goods_infoReq.Goods_infoSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Ginfo/lnt [POST]

type Goods_IDz struct {
	Goods_ID string
}

func (GinfoApi *Goods_infoApi) Goods_Int(c *gin.Context) {
	// 请添加自己的业务逻辑

	var form Goods_IDz
	c.ShouldBind(&form)
	newbio, err := GinfoService.Goods_Int(form.Goods_ID)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(newbio, c)
}
