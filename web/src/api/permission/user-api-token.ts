/** 用户API接口Token令牌表
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取 Token 令牌列表
export const getUserApiTokenList = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加 Token 令牌
export const addUserApiToken = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/add',
    method: 'post',
    data: params,
  });
};

// 更新 Token 令牌
export const updateUserApiToken = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/update',
    method: 'put',
    data: params,
  });
};

// 删除 Token 令牌
export const deleteUserApiToken = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除 Token 令牌
export const batchDeleteUserApiToken = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新 Token 令牌状态
export const updateUserApiTokenStatus = async (params: any) => {
  return axiosReq({
    url: '/userApiToken/updateStatus',
    method: 'put',
    data: params,
  });
};
