package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voln"
	volnReq "github.com/flipped-aurora/gin-vue-admin/server/model/voln/request"
)

type VulnsService struct {
}

// CreateVulns 创建vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) CreateVulns(vulns *voln.Vulns) (err error) {
	err = global.GVA_DB.Create(vulns).Error
	return err
}

// DeleteVulns 删除vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) DeleteVulns(id string) (err error) {
	err = global.GVA_DB.Delete(&voln.Vulns{}, "id = ?", id).Error
	return err
}

// DeleteVulnsByIds 批量删除vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) DeleteVulnsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]voln.Vulns{}, "id in ?", ids).Error
	return err
}

// UpdateVulns 更新vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) UpdateVulns(vulns voln.Vulns) (err error) {
	err = global.GVA_DB.Save(&vulns).Error
	return err
}

// GetVulns 根据id获取vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) GetVulns(id string) (vulns voln.Vulns, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&vulns).Error
	return
}

// GetVulnsInfoList 分页获取vulns表记录
// Author [piexlmax](https://github.com/piexlmax)
func (vulnsService *VulnsService) GetVulnsInfoList(info volnReq.VulnsSearch) (list []voln.Vulns, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&voln.Vulns{})
	var vulnss []voln.Vulns
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&vulnss).Error
	return vulnss, total, err
}
