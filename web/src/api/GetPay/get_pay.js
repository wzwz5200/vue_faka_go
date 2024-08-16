import service from '@/utils/request'

// @Tags Get_pay
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Get_pay true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Pay/createGet_pay [post]
export const createGet_pay = (data) => {
  return service({
    url: '/Pay/createGet_pay',
    method: 'post',
    data
  })
}

// @Tags Get_pay
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Get_pay true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pay/deleteGet_pay [delete]
export const deleteGet_pay = (params) => {
  return service({
    url: '/Pay/deleteGet_pay',
    method: 'delete',
    params
  })
}

// @Tags Get_pay
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pay/deleteGet_pay [delete]
export const deleteGet_payByIds = (params) => {
  return service({
    url: '/Pay/deleteGet_payByIds',
    method: 'delete',
    params
  })
}

// @Tags Get_pay
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Get_pay true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Pay/updateGet_pay [put]
export const updateGet_pay = (data) => {
  return service({
    url: '/Pay/updateGet_pay',
    method: 'put',
    data
  })
}

// @Tags Get_pay
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Get_pay true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pay/findGet_pay [get]
export const findGet_pay = (params) => {
  return service({
    url: '/Pay/findGet_pay',
    method: 'get',
    params
  })
}

// @Tags Get_pay
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pay/getGet_payList [get]
export const getGet_payList = (params) => {
  return service({
    url: '/Pay/getGet_payList',
    method: 'get',
    params
  })
}
