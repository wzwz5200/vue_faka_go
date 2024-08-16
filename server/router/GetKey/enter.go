package GetKey

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ GetKeysRouter }

var GKeyApi = api.ApiGroupApp.GetKeyApiGroup.GetKeysApi
