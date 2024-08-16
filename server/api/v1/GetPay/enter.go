package GetPay

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ Get_payApi }

var PayService = service.ServiceGroupApp.GetPayServiceGroup.Get_payService
