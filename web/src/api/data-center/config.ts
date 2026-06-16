/**
 * 配置
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取所有配置树
export async function getAllConfigTree() {
  return axiosReq({
    url: '/config/allTree',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取所有配置
export async function getConfigList(params: any) {
  return axiosReq({
    url: '/config/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 通过上级 key 获取子配置列表
export async function getConfigChildrensByKey(params: any) {
  return axiosReq({
    url: '/config/childrensByKey',
    method: 'get',
    isParams: true,
    data: params,
  })
}
// 获取配置树
export async function getConfigTree(params: any) {
  return axiosReq({
    url: '/config/tree',
    method: 'get',
    isParams: true,
    data: params,
  })
}
// 获取配置信息
export async function getConfigInfo(params: any) {
  return axiosReq({
    url: '/config/info',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加配置
export async function addConfig(params: any) {
  return axiosReq({
    url: '/config/add',
    method: 'post',
    data: params,
  })
}

// 更新配置
export async function updateConfig(params: any) {
  return axiosReq({
    url: '/config/update',
    method: 'put',
    data: params,
  })
}
// 批量更新配置
export async function batchUpdateConfig(params: any) {
  return axiosReq({
    url: '/config/batchUpdate',
    method: 'put',
    data: params,
  })
}

// 删除配置
export async function deleteConfig(params: any) {
  return axiosReq({
    url: '/config/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除配置
export async function batchDeleteConfig(params: any) {
  return axiosReq({
    url: '/config/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新配置状态
export async function updateConfigStatus(params: any) {
  return axiosReq({
    url: '/config/updateStatus',
    method: 'put',
    data: params,
  })
}

// 查询网站配置列表
export async function getWebSiteConfigList() {
  return axiosReq({
    url: '/config/webSiteConfigList',
    method: 'get',
    isParams: true,
    data: {},
  })
}
