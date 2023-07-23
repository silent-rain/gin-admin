/** 字典维度管理
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取字典维度信息列表
export const getDictList = async (params: any) => {
  return axiosReq({
    url: '/dict/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加字典维度信息
export const addDict = async (params: any) => {
  return axiosReq({
    url: '/dict/add',
    method: 'post',
    data: params,
  });
};

// 更新字典维度信息
export const updateDict = async (params: any) => {
  return axiosReq({
    url: '/dict/update',
    method: 'put',
    data: params,
  });
};

// 删除字典维度信息
export const deleteDict = async (params: any) => {
  return axiosReq({
    url: '/dict/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除字典维度信息
export const batchDeleteDict = async (params: any) => {
  return axiosReq({
    url: '/dict/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新字典维度信息状态
export const updateDictStatus = async (params: any) => {
  return axiosReq({
    url: '/dict/updateStatus',
    method: 'put',
    data: params,
  });
};
