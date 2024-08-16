import service from '@/utils/request'

// @Tags Key
// @Summary 创建卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "创建卡密"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /KEYS/createKey [post]
export const createKey = (data) => {
  return service({
    url: '/KEYS/createKey',
    method: 'post',
    data
  })
}

// @Tags Key
// @Summary 删除卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "删除卡密"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /KEYS/deleteKey [delete]
export const deleteKey = (params) => {
  return service({
    url: '/KEYS/deleteKey',
    method: 'delete',
    params
  })
}

// @Tags Key
// @Summary 批量删除卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除卡密"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /KEYS/deleteKey [delete]
export const deleteKeyByIds = (params) => {
  return service({
    url: '/KEYS/deleteKeyByIds',
    method: 'delete',
    params
  })
}

// @Tags Key
// @Summary 更新卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "更新卡密"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /KEYS/updateKey [put]
export const updateKey = (data) => {
  return service({
    url: '/KEYS/updateKey',
    method: 'put',
    data
  })
}

// @Tags Key
// @Summary 用id查询卡密
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Key true "用id查询卡密"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /KEYS/findKey [get]
export const findKey = (params) => {
  return service({
    url: '/KEYS/findKey',
    method: 'get',
    params
  })
}

// @Tags Key
// @Summary 分页获取卡密列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取卡密列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /KEYS/getKeyList [get]
export const getKeyList = (params) => {
  return service({
    url: '/KEYS/getKeyList',
    method: 'get',
    params
  })
}
