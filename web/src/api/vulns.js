import service from '@/utils/request'

// @Tags Vulns
// @Summary 创建vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulns true "创建vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /vulns/createVulns [post]
export const createVulns = (data) => {
  return service({
    url: '/vulns/createVulns',
    method: 'post',
    data
  })
}

// @Tags Vulns
// @Summary 删除vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulns true "删除vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vulns/deleteVulns [delete]
export const deleteVulns = (params) => {
  return service({
    url: '/vulns/deleteVulns',
    method: 'delete',
    params
  })
}

// @Tags Vulns
// @Summary 批量删除vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vulns/deleteVulns [delete]
export const deleteVulnsByIds = (params) => {
  return service({
    url: '/vulns/deleteVulnsByIds',
    method: 'delete',
    params
  })
}

// @Tags Vulns
// @Summary 更新vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulns true "更新vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vulns/updateVulns [put]
export const updateVulns = (data) => {
  return service({
    url: '/vulns/updateVulns',
    method: 'put',
    data
  })
}

// @Tags Vulns
// @Summary 用id查询vulns表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Vulns true "用id查询vulns表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vulns/findVulns [get]
export const findVulns = (params) => {
  return service({
    url: '/vulns/findVulns',
    method: 'get',
    params
  })
}

// @Tags Vulns
// @Summary 分页获取vulns表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取vulns表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vulns/getVulnsList [get]
export const getVulnsList = (params) => {
  return service({
    url: '/vulns/getVulnsList',
    method: 'get',
    params
  })
}
