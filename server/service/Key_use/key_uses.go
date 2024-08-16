package Key_use

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Key_use"
	Key_useReq "github.com/flipped-aurora/gin-vue-admin/server/model/Key_use/request"
)

type Key_usesService struct{}

// CreateKey_uses 创建key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) CreateKey_uses(Key_usess *Key_use.Key_uses) (err error) {
	err = global.GVA_DB.Create(Key_usess).Error
	return err
}

// DeleteKey_uses 删除key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) DeleteKey_uses(id string) (err error) {
	err = global.GVA_DB.Delete(&Key_use.Key_uses{}, "id = ?", id).Error
	return err
}

// DeleteKey_usesByIds 批量删除key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) DeleteKey_usesByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]Key_use.Key_uses{}, "id in ?", ids).Error
	return err
}

// UpdateKey_uses 更新key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) UpdateKey_uses(Key_usess Key_use.Key_uses) (err error) {
	err = global.GVA_DB.Model(&Key_use.Key_uses{}).Where("id = ?", Key_usess.Id).Updates(&Key_usess).Error
	return err
}

// GetKey_uses 根据id获取key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) GetKey_uses(id string) (Key_usess Key_use.Key_uses, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Key_usess).Error
	return
}

// GetKey_usesInfoList 分页获取key表记录
// Author [piexlmax](https://github.com/piexlmax)
func (Key_usessService *Key_usesService) GetKey_usesInfoList(info Key_useReq.Key_usesSearch) (list []Key_use.Key_uses, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Key_use.Key_uses{})
	var Key_usesss []Key_use.Key_uses
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Key_usesss).Error
	return Key_usesss, total, err
}

// Getkeys 请实现方法
// Author [yourname](https://github.com/yourname)
func (Key_usessService *Key_usesService) Getkeys(Keysz []string) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&Key_use.Key_uses{})

	var newz []Key_use.Key_z
	//	nk := newkeys

	for i := 0; i < 10; i++ {
		newkey := Key_use.Key_z{
			Keys: Keysz[i],
		}

		newz = append(newz, newkey)
	}
	nk := newz

	db.Create(&nk)
	return db.Error
}
