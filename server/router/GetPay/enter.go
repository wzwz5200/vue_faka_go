package GetPay

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ Get_payRouter }

var PayApi = api.ApiGroupApp.GetPayApiGroup.Get_payApi
