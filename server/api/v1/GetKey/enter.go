package GetKey

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ GetKeysApi }

var GKeyService = service.ServiceGroupApp.GetKeyServiceGroup.GetKeysService
