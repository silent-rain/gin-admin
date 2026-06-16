/**
 * 字典维度管理
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取字典维度信息列表
export async function getDictList(params: any) {
  return axiosReq({
    url: '/dict/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加字典维度信息
export async function addDict(params: any) {
  return axiosReq({
    url: '/dict/add',
    method: 'post',
    data: params,
  })
}

// 更新字典维度信息
export async function updateDict(params: any) {
  return axiosReq({
    url: '/dict/update',
    method: 'put',
    data: params,
  })
}

// 删除字典维度信息
export async function deleteDict(params: any) {
  return axiosReq({
    url: '/dict/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除字典维度信息
export async function batchDeleteDict(params: any) {
  return axiosReq({
    url: '/dict/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新字典维度信息状态
export async function updateDictStatus(params: any) {
  return axiosReq({
    url: '/dict/updateStatus',
    method: 'put',
    data: params,
  })
}
