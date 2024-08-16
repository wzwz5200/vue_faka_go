import service from '@/utils/request'

// @Tags Key_uses
// @Summary 创建key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key_uses true "创建key表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Key_usess/createKey_uses [post]
export const createKey_uses = (data) => {
  return service({
    url: '/Key_usess/createKey_uses',
    method: 'post',
    data
  })
}

// @Tags Key_uses
// @Summary 删除key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key_uses true "删除key表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Key_usess/deleteKey_uses [delete]
export const deleteKey_uses = (params) => {
  return service({
    url: '/Key_usess/deleteKey_uses',
    method: 'delete',
    params
  })
}

// @Tags Key_uses
// @Summary 批量删除key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除key表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Key_usess/deleteKey_uses [delete]
export const deleteKey_usesByIds = (params) => {
  return service({
    url: '/Key_usess/deleteKey_usesByIds',
    method: 'delete',
    params
  })
}

// @Tags Key_uses
// @Summary 更新key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key_uses true "更新key表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Key_usess/updateKey_uses [put]
export const updateKey_uses = (data) => {
  return service({
    url: '/Key_usess/updateKey_uses',
    method: 'put',
    data
  })
}

// @Tags Key_uses
// @Summary 用id查询key表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Key_uses true "用id查询key表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Key_usess/findKey_uses [get]
export const findKey_uses = (params) => {
  return service({
    url: '/Key_usess/findKey_uses',
    method: 'get',
    params
  })
}

// @Tags Key_uses
// @Summary 分页获取key表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取key表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Key_usess/getKey_usesList [get]
export const getKey_usesList = (params) => {
  return service({
    url: '/Key_usess/getKey_usesList',
    method: 'get',
    params
  })
}
