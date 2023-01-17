/**角色
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取所有角色列表
export const getAllRole = async (params: any) => {
  return axiosReq({
    url: '/role/all',
    method: 'get',
    isParams: true,
    data: params,
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
export const addRole = async () => {
  return axiosReq({
    url: '/role/add',
    method: 'post',
  });
};

// 更新角色
export const updateRole = async () => {
  return axiosReq({
    url: '/role/update',
    method: 'put',
  });
};

// 删除角色
export const deleteRole = async () => {
  return axiosReq({
    url: '/role/delete',
    method: 'delete',
  });
};

// 更新角色状态
export const updateRoleStatus = async () => {
  return axiosReq({
    url: '/role/status',
    method: 'put',
  });
};
