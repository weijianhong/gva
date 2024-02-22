package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voln"
    volnReq "github.com/flipped-aurora/gin-vue-admin/server/model/voln/request"
)

type WeakpassService struct {
}

// CreateWeakpass 创建weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService) CreateWeakpass(weakpass *voln.Weakpass) (err error) {
	err = global.GVA_DB.Create(weakpass).Error
	return err
}

// DeleteWeakpass 删除weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService)DeleteWeakpass(id string) (err error) {
	err = global.GVA_DB.Delete(&voln.Weakpass{},"id = ?",id).Error
	return err
}

// DeleteWeakpassByIds 批量删除weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService)DeleteWeakpassByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]voln.Weakpass{},"id in ?",ids).Error
	return err
}

// UpdateWeakpass 更新weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService)UpdateWeakpass(weakpass voln.Weakpass) (err error) {
	err = global.GVA_DB.Save(&weakpass).Error
	return err
}

// GetWeakpass 根据id获取weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService)GetWeakpass(id string) (weakpass voln.Weakpass, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&weakpass).Error
	return
}

// GetWeakpassInfoList 分页获取weakpass表记录
// Author [piexlmax](https://github.com/piexlmax)
func (weakpassService *WeakpassService)GetWeakpassInfoList(info volnReq.WeakpassSearch) (list []voln.Weakpass, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&voln.Weakpass{})
    var weakpasss []voln.Weakpass
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
	
	err = db.Find(&weakpasss).Error
	return  weakpasss, total, err
}
