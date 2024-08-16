package GetKey

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/GetKey"
    GetKeyReq "github.com/flipped-aurora/gin-vue-admin/server/model/GetKey/request"
)

type GetKeysService struct {}

// CreateGetKeys 创建获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService) CreateGetKeys(GKey *GetKey.GetKeys) (err error) {
	err = global.GVA_DB.Create(GKey).Error
	return err
}

// DeleteGetKeys 删除获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService)DeleteGetKeys(ID string) (err error) {
	err = global.GVA_DB.Delete(&GetKey.GetKeys{},"id = ?",ID).Error
	return err
}

// DeleteGetKeysByIds 批量删除获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService)DeleteGetKeysByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]GetKey.GetKeys{},"id in ?",IDs).Error
	return err
}

// UpdateGetKeys 更新获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService)UpdateGetKeys(GKey GetKey.GetKeys) (err error) {
	err = global.GVA_DB.Model(&GetKey.GetKeys{}).Where("id = ?",GKey.ID).Updates(&GKey).Error
	return err
}

// GetGetKeys 根据ID获取获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService)GetGetKeys(ID string) (GKey GetKey.GetKeys, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&GKey).Error
	return
}

// GetGetKeysInfoList 分页获取获取key记录
// Author [piexlmax](https://github.com/piexlmax)
func (GKeyService *GetKeysService)GetGetKeysInfoList(info GetKeyReq.GetKeysSearch) (list []GetKey.GetKeys, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&GetKey.GetKeys{})
    var GKeys []GetKey.GetKeys
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
	
	err = db.Find(&GKeys).Error
	return  GKeys, total, err
}
// GetKey_psot 请实现方法
// Author [yourname](https://github.com/yourname)
func (GKeyService *GetKeysService)GetKey_psot() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&GetKey.GetKeys{})

	


    return db.Error
}

