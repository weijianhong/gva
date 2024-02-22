package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/voln"
    volnReq "github.com/flipped-aurora/gin-vue-admin/server/model/voln/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type VulnsApi struct {
}

var vulnsService = service.ServiceGroupApp.VolnServiceGroup.VulnsService


// CreateVulns 创建vulns表
// @Tags Vulns
// @Summary 创建vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Vulns true "创建vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /vulns/createVulns [post]
func (vulnsApi *VulnsApi) CreateVulns(c *gin.Context) {
	var vulns voln.Vulns
	err := c.ShouldBindJSON(&vulns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := vulnsService.CreateVulns(&vulns); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVulns 删除vulns表
// @Tags Vulns
// @Summary 删除vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Vulns true "删除vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vulns/deleteVulns [delete]
func (vulnsApi *VulnsApi) DeleteVulns(c *gin.Context) {
	id := c.Query("ID")
	if err := vulnsService.DeleteVulns(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVulnsByIds 批量删除vulns表
// @Tags Vulns
// @Summary 批量删除vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /vulns/deleteVulnsByIds [delete]
func (vulnsApi *VulnsApi) DeleteVulnsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := vulnsService.DeleteVulnsByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVulns 更新vulns表
// @Tags Vulns
// @Summary 更新vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Vulns true "更新vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vulns/updateVulns [put]
func (vulnsApi *VulnsApi) UpdateVulns(c *gin.Context) {
	var vulns voln.Vulns
	err := c.ShouldBindJSON(&vulns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := vulnsService.UpdateVulns(vulns); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVulns 用id查询vulns表
// @Tags Vulns
// @Summary 用id查询vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voln.Vulns true "用id查询vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vulns/findVulns [get]
func (vulnsApi *VulnsApi) FindVulns(c *gin.Context) {
	id := c.Query("ID")
	if revulns, err := vulnsService.GetVulns(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revulns": revulns}, c)
	}
}

// GetVulnsList 分页获取vulns表列表
// @Tags Vulns
// @Summary 分页获取vulns表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query volnReq.VulnsSearch true "分页获取vulns表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vulns/getVulnsList [get]
func (vulnsApi *VulnsApi) GetVulnsList(c *gin.Context) {
	var pageInfo volnReq.VulnsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := vulnsService.GetVulnsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
