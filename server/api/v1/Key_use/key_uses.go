package Key_use

import (
	"io"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Key_use"
	Key_useReq "github.com/flipped-aurora/gin-vue-admin/server/model/Key_use/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Key_usesApi struct{}

// CreateKey_uses 创建key表
// @Tags Key_uses
// @Summary 创建key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Key_use.Key_uses true "创建key表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Key_usess/createKey_uses [post]
func (Key_usessApi *Key_usesApi) CreateKey_uses(c *gin.Context) {
	var Key_usess Key_use.Key_uses
	err := c.ShouldBindJSON(&Key_usess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = Key_usessService.CreateKey_uses(&Key_usess)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteKey_uses 删除key表
// @Tags Key_uses
// @Summary 删除key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Key_use.Key_uses true "删除key表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Key_usess/deleteKey_uses [delete]
func (Key_usessApi *Key_usesApi) DeleteKey_uses(c *gin.Context) {
	id := c.Query("id")
	err := Key_usessService.DeleteKey_uses(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteKey_usesByIds 批量删除key表
// @Tags Key_uses
// @Summary 批量删除key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Key_usess/deleteKey_usesByIds [delete]
func (Key_usessApi *Key_usesApi) DeleteKey_usesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := Key_usessService.DeleteKey_usesByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateKey_uses 更新key表
// @Tags Key_uses
// @Summary 更新key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Key_use.Key_uses true "更新key表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Key_usess/updateKey_uses [put]
func (Key_usessApi *Key_usesApi) UpdateKey_uses(c *gin.Context) {
	var Key_usess Key_use.Key_uses
	err := c.ShouldBindJSON(&Key_usess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = Key_usessService.UpdateKey_uses(Key_usess)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindKey_uses 用id查询key表
// @Tags Key_uses
// @Summary 用id查询key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Key_use.Key_uses true "用id查询key表"
// @Success 200 {object} response.Response{data=Key_use.Key_uses,msg=string} "查询成功"
// @Router /Key_usess/findKey_uses [get]
func (Key_usessApi *Key_usesApi) FindKey_uses(c *gin.Context) {
	id := c.Query("id")
	reKey_usess, err := Key_usessService.GetKey_uses(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reKey_usess, c)
}

// GetKey_usesList 分页获取key表列表
// @Tags Key_uses
// @Summary 分页获取key表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Key_useReq.Key_usesSearch true "分页获取key表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Key_usess/getKey_usesList [get]
func (Key_usessApi *Key_usesApi) GetKey_usesList(c *gin.Context) {
	var pageInfo Key_useReq.Key_usesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := Key_usessService.GetKey_usesInfoList(pageInfo)
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

// GetKey_usesPublic 不需要鉴权的key表接口
// @Tags Key_uses
// @Summary 不需要鉴权的key表接口
// @accept application/json
// @Produce application/json
// @Param data query Key_useReq.Key_usesSearch true "分页获取key表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Key_usess/getKey_usesPublic [get]
func (Key_usessApi *Key_usesApi) GetKey_usesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的key表接口信息",
	}, "获取成功", c)
}

// Getkeys 等待开发的的key表接口
// @Tags Key_uses
// @Summary 等待开发的的key表接口
// @accept application/json
// @Produce application/json
// @Param data query Key_useReq.Key_usesSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /Key_usess/Getkey_p [POST]
func (Key_usessApi *Key_usesApi) Getkeys(c *gin.Context) {
	// 请添加自己的业务逻辑
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	words := string(body)
	wz := strings.Fields(words)
	Key_usessService.Getkeys(wz)

	response.OkWithData(wz, c)
}
