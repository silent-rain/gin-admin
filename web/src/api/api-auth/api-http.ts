/**
 * Http协议接口管理
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取所有Http协议接口信息列表
export async function getAllApiHttpTree() {
  return axiosReq({
    url: '/apiHttp/allTree',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取Http协议接口信息列表
export async function getApiHttpTree(params: any) {
  return axiosReq({
    url: '/apiHttp/tree',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加Http协议接口信息
export async function addApiHttp(params: any) {
  return axiosReq({
    url: '/apiHttp/add',
    method: 'post',
    data: params,
  })
}

// 更新Http协议接口信息
export async function updateApiHttp(params: any) {
  return axiosReq({
    url: '/apiHttp/update',
    method: 'put',
    data: params,
  })
}

// 删除Http协议接口信息
export async function deleteApiHttp(params: any) {
  return axiosReq({
    url: '/apiHttp/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除Http协议接口信息
export async function batchDeleteApiHttp(params: any) {
  return axiosReq({
    url: '/apiHttp/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新Http协议接口信息状态
export async function updateApiHttpStatus(params: any) {
  return axiosReq({
    url: '/apiHttp/updateStatus',
    method: 'put',
    data: params,
  })
}
