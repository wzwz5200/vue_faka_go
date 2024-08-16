import service from '@/utils/request'

// @Tags Goods_info
// @Summary 创建获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods_info true "创建获取商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Ginfo/createGoods_info [post]
export const createGoods_info = (data) => {
  return service({
    url: '/Ginfo/createGoods_info',
    method: 'post',
    data
  })
}

// @Tags Goods_info
// @Summary 删除获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods_info true "删除获取商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ginfo/deleteGoods_info [delete]
export const deleteGoods_info = (params) => {
  return service({
    url: '/Ginfo/deleteGoods_info',
    method: 'delete',
    params
  })
}

// @Tags Goods_info
// @Summary 批量删除获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除获取商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ginfo/deleteGoods_info [delete]
export const deleteGoods_infoByIds = (params) => {
  return service({
    url: '/Ginfo/deleteGoods_infoByIds',
    method: 'delete',
    params
  })
}

// @Tags Goods_info
// @Summary 更新获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods_info true "更新获取商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Ginfo/updateGoods_info [put]
export const updateGoods_info = (data) => {
  return service({
    url: '/Ginfo/updateGoods_info',
    method: 'put',
    data
  })
}

// @Tags Goods_info
// @Summary 用id查询获取商品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Goods_info true "用id查询获取商品信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Ginfo/findGoods_info [get]
export const findGoods_info = (params) => {
  return service({
    url: '/Ginfo/findGoods_info',
    method: 'get',
    params
  })
}

// @Tags Goods_info
// @Summary 分页获取获取商品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取获取商品信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ginfo/getGoods_infoList [get]
export const getGoods_infoList = (params) => {
  return service({
    url: '/Ginfo/getGoods_infoList',
    method: 'get',
    params
  })
}
