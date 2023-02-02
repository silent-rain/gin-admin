/**角色
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取所有角色列表
export const getAllRole = async () => {
  return axiosReq({
    url: '/role/all',
    method: 'get',
    isParams: true,
    data: {},
  });
};

// 获取角色列表
export const getRoleList = async (params: any) => {
  return axiosReq({
    url: '/role/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加角色
export const addRole = async (params: any) => {
  return axiosReq({
    url: '/role/add',
    method: 'post',
    data: params,
  });
};

// 更新角色
export const updateRole = async (params: any) => {
  return axiosReq({
    url: '/role/update',
    method: 'put',
    data: params,
  });
};

// 删除角色
export const deleteRole = async (params: any) => {
  return axiosReq({
    url: '/role/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除角色
export const batchDeleteRole = async (params: any) => {
  return axiosReq({
    url: '/role/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新角色状态
export const updateRoleStatus = async (params: any) => {
  return axiosReq({
    url: '/role/status',
    method: 'put',
    data: params,
  });
};
