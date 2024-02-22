package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VulnsRouter struct {
}

// InitVulnsRouter 初始化 vulns表 路由信息
func (s *VulnsRouter) InitVulnsRouter(Router *gin.RouterGroup) {
	vulnsRouter := Router.Group("vulns").Use(middleware.OperationRecord())
	vulnsRouterWithoutRecord := Router.Group("vulns")
	var vulnsApi = v1.ApiGroupApp.VolnApiGroup.VulnsApi
	{
		vulnsRouter.POST("createVulns", vulnsApi.CreateVulns)   // 新建vulns表
		vulnsRouter.DELETE("deleteVulns", vulnsApi.DeleteVulns) // 删除vulns表
		vulnsRouter.DELETE("deleteVulnsByIds", vulnsApi.DeleteVulnsByIds) // 批量删除vulns表
		vulnsRouter.PUT("updateVulns", vulnsApi.UpdateVulns)    // 更新vulns表
	}
	{
		vulnsRouterWithoutRecord.GET("findVulns", vulnsApi.FindVulns)        // 根据ID获取vulns表
		vulnsRouterWithoutRecord.GET("getVulnsList", vulnsApi.GetVulnsList)  // 获取vulns表列表
	}
}
