package voln

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WeakpassRouter struct {
}

// InitWeakpassRouter 初始化 weakpass表 路由信息
func (s *WeakpassRouter) InitWeakpassRouter(Router *gin.RouterGroup) {
	weakpassRouter := Router.Group("weakpass").Use(middleware.OperationRecord())
	weakpassRouterWithoutRecord := Router.Group("weakpass")
	var weakpassApi = v1.ApiGroupApp.VolnApiGroup.WeakpassApi
	{
		weakpassRouter.POST("createWeakpass", weakpassApi.CreateWeakpass)   // 新建weakpass表
		weakpassRouter.DELETE("deleteWeakpass", weakpassApi.DeleteWeakpass) // 删除weakpass表
		weakpassRouter.DELETE("deleteWeakpassByIds", weakpassApi.DeleteWeakpassByIds) // 批量删除weakpass表
		weakpassRouter.PUT("updateWeakpass", weakpassApi.UpdateWeakpass)    // 更新weakpass表
	}
	{
		weakpassRouterWithoutRecord.GET("findWeakpass", weakpassApi.FindWeakpass)        // 根据ID获取weakpass表
		weakpassRouterWithoutRecord.GET("getWeakpassList", weakpassApi.GetWeakpassList)  // 获取weakpass表列表
	}
}
