/**配置
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取所有配置树
export const getAllConfigTree = async () => {
  return axiosReq({
    url: '/config/allTree',
    method: 'get',
    isParams: true,
    data: {},
  });
};

// 获取配置树
export const getConfigTree = async (params: any) => {
  return axiosReq({
    url: '/config/tree',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加配置
export const addConfig = async (params: any) => {
  return axiosReq({
    url: '/config/add',
    method: 'post',
    data: params,
  });
};

// 更新配置
export const updateConfig = async (params: any) => {
  return axiosReq({
    url: '/config/update',
    method: 'put',
    data: params,
  });
};

// 删除配置
export const deleteConfig = async (params: any) => {
  return axiosReq({
    url: '/config/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除配置
export const batchDeleteConfig = async (params: any) => {
  return axiosReq({
    url: '/config/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新配置状态
export const updateConfigStatus = async (params: any) => {
  return axiosReq({
    url: '/config/status',
    method: 'put',
    data: params,
  });
};
