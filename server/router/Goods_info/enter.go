package Goods_info

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ Goods_infoRouter }

var GinfoApi = api.ApiGroupApp.Goods_infoApiGroup.Goods_infoApi
