import service from '@/utils/request'

// @Tags Weakpass
// @Summary 创建weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Weakpass true "创建weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weakpass/createWeakpass [post]
export const createWeakpass = (data) => {
  return service({
    url: '/weakpass/createWeakpass',
    method: 'post',
    data
  })
}

// @Tags Weakpass
// @Summary 删除weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Weakpass true "删除weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weakpass/deleteWeakpass [delete]
export const deleteWeakpass = (params) => {
  return service({
    url: '/weakpass/deleteWeakpass',
    method: 'delete',
    params
  })
}

// @Tags Weakpass
// @Summary 批量删除weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weakpass/deleteWeakpass [delete]
export const deleteWeakpassByIds = (params) => {
  return service({
    url: '/weakpass/deleteWeakpassByIds',
    method: 'delete',
    params
  })
}

// @Tags Weakpass
// @Summary 更新weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Weakpass true "更新weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weakpass/updateWeakpass [put]
export const updateWeakpass = (data) => {
  return service({
    url: '/weakpass/updateWeakpass',
    method: 'put',
    data
  })
}

// @Tags Weakpass
// @Summary 用id查询weakpass表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Weakpass true "用id查询weakpass表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weakpass/findWeakpass [get]
export const findWeakpass = (params) => {
  return service({
    url: '/weakpass/findWeakpass',
    method: 'get',
    params
  })
}

// @Tags Weakpass
// @Summary 分页获取weakpass表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取weakpass表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weakpass/getWeakpassList [get]
export const getWeakpassList = (params) => {
  return service({
    url: '/weakpass/getWeakpassList',
    method: 'get',
    params
  })
}
