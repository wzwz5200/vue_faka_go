package Goods_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Goods_info"
	Goods_infoReq "github.com/flipped-aurora/gin-vue-admin/server/model/Goods_info/request"
)

type Goods_infoService struct{}

// CreateGoods_info 创建获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) CreateGoods_info(Ginfo *Goods_info.Goods_info) (err error) {
	err = global.GVA_DB.Create(Ginfo).Error
	return err
}

// DeleteGoods_info 删除获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) DeleteGoods_info(ID string) (err error) {
	err = global.GVA_DB.Delete(&Goods_info.Goods_info{}, "id = ?", ID).Error
	return err
}

// DeleteGoods_infoByIds 批量删除获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) DeleteGoods_infoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Goods_info.Goods_info{}, "id in ?", IDs).Error
	return err
}

// UpdateGoods_info 更新获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) UpdateGoods_info(Ginfo Goods_info.Goods_info) (err error) {
	err = global.GVA_DB.Model(&Goods_info.Goods_info{}).Where("id = ?", Ginfo.ID).Updates(&Ginfo).Error
	return err
}

// GetGoods_info 根据ID获取获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) GetGoods_info(ID string) (Ginfo Goods_info.Goods_info, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Ginfo).Error
	return
}

// GetGoods_infoInfoList 分页获取获取商品信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (GinfoService *Goods_infoService) GetGoods_infoInfoList(info Goods_infoReq.Goods_infoSearch) (list []Goods_info.Goods_info, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Goods_info.Goods_info{})
	var Ginfos []Goods_info.Goods_info
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Price != "" {
		db = db.Where("price = ?", info.Price)
	}
	if info.Goods_id != "" {
		db = db.Where("goods_id = ?", info.Goods_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Ginfos).Error
	return Ginfos, total, err
}

// GetGoods_info 请实现方法
// Author [yourname](https://github.com/yourname)
func (GinfoService *Goods_infoService) GetGoods() (list []Goods_info.Goods_info, err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&Goods_info.Goods_info{})

	db.Find(&list)

	return list, db.Error
}

// Goods_Int 请实现方法
// Author [yourname](https://github.com/yourname)
func (GinfoService *Goods_infoService) Goods_Int(Goods_id string) (list Goods_info.Goods_info, err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&Goods_info.Goods_info{})

	db.Where("goods_id = ?", Goods_id).First(&list)

	if list.Goods_id == "" {
		return

	}

	return list, db.Error
}
