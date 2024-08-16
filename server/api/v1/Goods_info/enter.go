package Goods_info

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ Goods_infoApi }

var GinfoService = service.ServiceGroupApp.Goods_infoServiceGroup.Goods_infoService
