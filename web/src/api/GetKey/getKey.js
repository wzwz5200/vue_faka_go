import service from '@/utils/request'

// @Tags GetKeys
// @Summary 创建获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GetKeys true "创建获取key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /GKey/createGetKeys [post]
export const createGetKeys = (data) => {
  return service({
    url: '/GKey/createGetKeys',
    method: 'post',
    data
  })
}

// @Tags GetKeys
// @Summary 删除获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GetKeys true "删除获取key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /GKey/deleteGetKeys [delete]
export const deleteGetKeys = (params) => {
  return service({
    url: '/GKey/deleteGetKeys',
    method: 'delete',
    params
  })
}

// @Tags GetKeys
// @Summary 批量删除获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除获取key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /GKey/deleteGetKeys [delete]
export const deleteGetKeysByIds = (params) => {
  return service({
    url: '/GKey/deleteGetKeysByIds',
    method: 'delete',
    params
  })
}

// @Tags GetKeys
// @Summary 更新获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GetKeys true "更新获取key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /GKey/updateGetKeys [put]
export const updateGetKeys = (data) => {
  return service({
    url: '/GKey/updateGetKeys',
    method: 'put',
    data
  })
}

// @Tags GetKeys
// @Summary 用id查询获取key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GetKeys true "用id查询获取key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /GKey/findGetKeys [get]
export const findGetKeys = (params) => {
  return service({
    url: '/GKey/findGetKeys',
    method: 'get',
    params
  })
}

// @Tags GetKeys
// @Summary 分页获取获取key列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取获取key列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /GKey/getGetKeysList [get]
export const getGetKeysList = (params) => {
  return service({
    url: '/GKey/getGetKeysList',
    method: 'get',
    params
  })
}
