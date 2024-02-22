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

type WeakpassApi struct {
}

var weakpassService = service.ServiceGroupApp.VolnServiceGroup.WeakpassService


// CreateWeakpass 创建weakpass表
// @Tags Weakpass
// @Summary 创建weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Weakpass true "创建weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weakpass/createWeakpass [post]
func (weakpassApi *WeakpassApi) CreateWeakpass(c *gin.Context) {
	var weakpass voln.Weakpass
	err := c.ShouldBindJSON(&weakpass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := weakpassService.CreateWeakpass(&weakpass); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWeakpass 删除weakpass表
// @Tags Weakpass
// @Summary 删除weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Weakpass true "删除weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weakpass/deleteWeakpass [delete]
func (weakpassApi *WeakpassApi) DeleteWeakpass(c *gin.Context) {
	id := c.Query("ID")
	if err := weakpassService.DeleteWeakpass(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWeakpassByIds 批量删除weakpass表
// @Tags Weakpass
// @Summary 批量删除weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /weakpass/deleteWeakpassByIds [delete]
func (weakpassApi *WeakpassApi) DeleteWeakpassByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := weakpassService.DeleteWeakpassByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWeakpass 更新weakpass表
// @Tags Weakpass
// @Summary 更新weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body voln.Weakpass true "更新weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weakpass/updateWeakpass [put]
func (weakpassApi *WeakpassApi) UpdateWeakpass(c *gin.Context) {
	var weakpass voln.Weakpass
	err := c.ShouldBindJSON(&weakpass)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := weakpassService.UpdateWeakpass(weakpass); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWeakpass 用id查询weakpass表
// @Tags Weakpass
// @Summary 用id查询weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query voln.Weakpass true "用id查询weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weakpass/findWeakpass [get]
func (weakpassApi *WeakpassApi) FindWeakpass(c *gin.Context) {
	id := c.Query("ID")
	if reweakpass, err := weakpassService.GetWeakpass(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reweakpass": reweakpass}, c)
	}
}

// GetWeakpassList 分页获取weakpass表列表
// @Tags Weakpass
// @Summary 分页获取weakpass表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query volnReq.WeakpassSearch true "分页获取weakpass表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weakpass/getWeakpassList [get]
func (weakpassApi *WeakpassApi) GetWeakpassList(c *gin.Context) {
	var pageInfo volnReq.WeakpassSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := weakpassService.GetWeakpassInfoList(pageInfo); err != nil {
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
