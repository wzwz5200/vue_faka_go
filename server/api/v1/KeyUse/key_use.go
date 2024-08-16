package KeyUse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/KeyUse"
    KeyUseReq "github.com/flipped-aurora/gin-vue-admin/server/model/KeyUse/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type KeyApi struct {}

// CreateKey 创建卡密
// @Tags Key
// @Summary 创建卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body KeyUse.Key true "创建卡密"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /KEYS/createKey [post]
func (KEYSApi *KeyApi) CreateKey(c *gin.Context) {
	var KEYS KeyUse.Key
	err := c.ShouldBindJSON(&KEYS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = KEYSService.CreateKey(&KEYS)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteKey 删除卡密
// @Tags Key
// @Summary 删除卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body KeyUse.Key true "删除卡密"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /KEYS/deleteKey [delete]
func (KEYSApi *KeyApi) DeleteKey(c *gin.Context) {
	ID := c.Query("ID")
	err := KEYSService.DeleteKey(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteKeyByIds 批量删除卡密
// @Tags Key
// @Summary 批量删除卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /KEYS/deleteKeyByIds [delete]
func (KEYSApi *KeyApi) DeleteKeyByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := KEYSService.DeleteKeyByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateKey 更新卡密
// @Tags Key
// @Summary 更新卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body KeyUse.Key true "更新卡密"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /KEYS/updateKey [put]
func (KEYSApi *KeyApi) UpdateKey(c *gin.Context) {
	var KEYS KeyUse.Key
	err := c.ShouldBindJSON(&KEYS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = KEYSService.UpdateKey(KEYS)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindKey 用id查询卡密
// @Tags Key
// @Summary 用id查询卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query KeyUse.Key true "用id查询卡密"
// @Success 200 {object} response.Response{data=KeyUse.Key,msg=string} "查询成功"
// @Router /KEYS/findKey [get]
func (KEYSApi *KeyApi) FindKey(c *gin.Context) {
	ID := c.Query("ID")
	reKEYS, err := KEYSService.GetKey(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reKEYS, c)
}

// GetKeyList 分页获取卡密列表
// @Tags Key
// @Summary 分页获取卡密列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query KeyUseReq.KeySearch true "分页获取卡密列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /KEYS/getKeyList [get]
func (KEYSApi *KeyApi) GetKeyList(c *gin.Context) {
	var pageInfo KeyUseReq.KeySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := KEYSService.GetKeyInfoList(pageInfo)
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

// GetKeyPublic 不需要鉴权的卡密接口
// @Tags Key
// @Summary 不需要鉴权的卡密接口
// @accept application/json
// @Produce application/json
// @Param data query KeyUseReq.KeySearch true "分页获取卡密列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /KEYS/getKeyPublic [get]
func (KEYSApi *KeyApi) GetKeyPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的卡密接口信息",
    }, "获取成功", c)
}
