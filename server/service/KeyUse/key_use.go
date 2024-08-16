package KeyUse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/KeyUse"
    KeyUseReq "github.com/flipped-aurora/gin-vue-admin/server/model/KeyUse/request"
)

type KeyService struct {}

// CreateKey 创建卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService) CreateKey(KEYS *KeyUse.Key) (err error) {
	err = global.GVA_DB.Create(KEYS).Error
	return err
}

// DeleteKey 删除卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService)DeleteKey(ID string) (err error) {
	err = global.GVA_DB.Delete(&KeyUse.Key{},"id = ?",ID).Error
	return err
}

// DeleteKeyByIds 批量删除卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService)DeleteKeyByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]KeyUse.Key{},"id in ?",IDs).Error
	return err
}

// UpdateKey 更新卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService)UpdateKey(KEYS KeyUse.Key) (err error) {
	err = global.GVA_DB.Model(&KeyUse.Key{}).Where("id = ?",KEYS.ID).Updates(&KEYS).Error
	return err
}

// GetKey 根据ID获取卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService)GetKey(ID string) (KEYS KeyUse.Key, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&KEYS).Error
	return
}

// GetKeyInfoList 分页获取卡密记录
// Author [piexlmax](https://github.com/piexlmax)
func (KEYSService *KeyService)GetKeyInfoList(info KeyUseReq.KeySearch) (list []KeyUse.Key, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&KeyUse.Key{})
    var KEYSs []KeyUse.Key
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&KEYSs).Error
	return  KEYSs, total, err
}