/**
 * 字典数据管理
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取字典数据信息列表
export async function getDictDataList(params: any) {
  return axiosReq({
    url: '/dictData/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加字典数据信息
export async function addDictData(params: any) {
  return axiosReq({
    url: '/dictData/add',
    method: 'post',
    data: params,
  })
}

// 更新字典数据信息
export async function updateDictData(params: any) {
  return axiosReq({
    url: '/dictData/update',
    method: 'put',
    data: params,
  })
}

// 删除字典数据信息
export async function deleteDictData(params: any) {
  return axiosReq({
    url: '/dictData/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除字典数据信息
export async function batchDeleteDictData(params: any) {
  return axiosReq({
    url: '/dictData/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新字典数据信息状态
export async function updateDictDataStatus(params: any) {
  return axiosReq({
    url: '/dictData/updateStatus',
    method: 'put',
    data: params,
  })
}
