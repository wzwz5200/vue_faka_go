package GetKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/GetKey"
    GetKeyReq "github.com/flipped-aurora/gin-vue-admin/server/model/GetKey/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GetKeysApi struct {}

// CreateGetKeys 创建获取key
// @Tags GetKeys
// @Summary 创建获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetKey.GetKeys true "创建获取key"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /GKey/createGetKeys [post]
func (GKeyApi *GetKeysApi) CreateGetKeys(c *gin.Context) {
	var GKey GetKey.GetKeys
	err := c.ShouldBindJSON(&GKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = GKeyService.CreateGetKeys(&GKey)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGetKeys 删除获取key
// @Tags GetKeys
// @Summary 删除获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetKey.GetKeys true "删除获取key"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /GKey/deleteGetKeys [delete]
func (GKeyApi *GetKeysApi) DeleteGetKeys(c *gin.Context) {
	ID := c.Query("ID")
	err := GKeyService.DeleteGetKeys(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGetKeysByIds 批量删除获取key
// @Tags GetKeys
// @Summary 批量删除获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /GKey/deleteGetKeysByIds [delete]
func (GKeyApi *GetKeysApi) DeleteGetKeysByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := GKeyService.DeleteGetKeysByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGetKeys 更新获取key
// @Tags GetKeys
// @Summary 更新获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body GetKey.GetKeys true "更新获取key"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /GKey/updateGetKeys [put]
func (GKeyApi *GetKeysApi) UpdateGetKeys(c *gin.Context) {
	var GKey GetKey.GetKeys
	err := c.ShouldBindJSON(&GKey)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = GKeyService.UpdateGetKeys(GKey)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGetKeys 用id查询获取key
// @Tags GetKeys
// @Summary 用id查询获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query GetKey.GetKeys true "用id查询获取key"
// @Success 200 {object} response.Response{data=GetKey.GetKeys,msg=string} "查询成功"
// @Router /GKey/findGetKeys [get]
func (GKeyApi *GetKeysApi) FindGetKeys(c *gin.Context) {
	ID := c.Query("ID")
	reGKey, err := GKeyService.GetGetKeys(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reGKey, c)
}

// GetGetKeysList 分页获取获取key列表
// @Tags GetKeys
// @Summary 分页获取获取key列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query GetKeyReq.GetKeysSearch true "分页获取获取key列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /GKey/getGetKeysList [get]
func (GKeyApi *GetKeysApi) GetGetKeysList(c *gin.Context) {
	var pageInfo GetKeyReq.GetKeysSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := GKeyService.GetGetKeysInfoList(pageInfo)
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

// GetGetKeysPublic 不需要鉴权的获取key接口
// @Tags GetKeys
// @Summary 不需要鉴权的获取key接口
// @accept application/json
// @Produce application/json
// @Param data query GetKeyReq.GetKeysSearch true "分页获取获取key列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /GKey/getGetKeysPublic [get]
func (GKeyApi *GetKeysApi) GetGetKeysPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的获取key接口信息",
    }, "获取成功", c)
}
// GetKey_psot 等待开发的的获取key接口
// @Tags GetKeys
// @Summary 等待开发的的获取key接口
// @accept application/json
// @Produce application/json
// @Param data query GetKeyReq.GetKeysSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /GKey/key [POST]
func (GKeyApi *GetKeysApi)GetKey_psot(c *gin.Context) {
    // 请添加自己的业务逻辑
    err := GKeyService.GetKey_psot()
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}

