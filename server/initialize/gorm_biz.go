package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/GetKey"
	"github.com/flipped-aurora/gin-vue-admin/server/model/GetPay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Goods_info"
	"github.com/flipped-aurora/gin-vue-admin/server/model/KeyUse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Key_use"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(GetPay.Order{}, GetPay.Get_pay{}, KeyUse.Key{}, Key_use.Key_uses{}, GetKey.GetKeys{}, Goods_info.Goods_info{})
	if err != nil {
		return err
	}
	return nil
}
