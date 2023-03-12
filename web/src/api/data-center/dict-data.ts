/** 字典数据管理
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取字典数据信息列表
export const getDictDataList = async (params: any) => {
  return axiosReq({
    url: '/dictData/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加字典数据信息
export const addDictData = async (params: any) => {
  return axiosReq({
    url: '/dictData/add',
    method: 'post',
    data: params,
  });
};

// 更新字典数据信息
export const updateDictData = async (params: any) => {
  return axiosReq({
    url: '/dictData/update',
    method: 'put',
    data: params,
  });
};

// 删除字典数据信息
export const deleteDictData = async (params: any) => {
  return axiosReq({
    url: '/dictData/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除字典数据信息
export const batchDeleteDictData = async (params: any) => {
  return axiosReq({
    url: '/dictData/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新字典数据信息状态
export const updateDictDataStatus = async (params: any) => {
  return axiosReq({
    url: '/dictData/status',
    method: 'put',
    data: params,
  });
};
