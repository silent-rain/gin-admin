/**菜单
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取所有菜单列表
export const getAllMenu = async () => {
  return axiosReq({
    url: '/menu/all',
    method: 'get',
    isParams: true,
    data: {},
  });
};

// 获取菜单列表
export const getMenuList = async (params: any) => {
  return axiosReq({
    url: '/menu/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加菜单
export const addMenu = async (params: any) => {
  return axiosReq({
    url: '/menu/add',
    method: 'post',
    data: params,
  });
};

// 更新菜单
export const updateMenu = async (params: any) => {
  return axiosReq({
    url: '/menu/update',
    method: 'put',
    data: params,
  });
};

// 删除菜单
export const deleteMenu = async (params: any) => {
  return axiosReq({
    url: '/menu/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除菜单
export const batchDeleteMenu = async (params: any) => {
  return axiosReq({
    url: '/menu/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新菜单状态
export const updateMenuStatus = async (params: any) => {
  return axiosReq({
    url: '/menu/status',
    method: 'put',
    data: params,
  });
};
