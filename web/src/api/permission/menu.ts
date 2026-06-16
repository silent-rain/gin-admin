/**
 * 菜单
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取所有菜单树
export async function getAllMenuTree() {
  return axiosReq({
    url: '/menu/allTree',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取菜单树
export async function getMenuTree(params: any) {
  return axiosReq({
    url: '/menu/tree',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加菜单
export async function addMenu(params: any) {
  return axiosReq({
    url: '/menu/add',
    method: 'post',
    data: params,
  })
}

// 更新菜单
export async function updateMenu(params: any) {
  return axiosReq({
    url: '/menu/update',
    method: 'put',
    data: params,
  })
}

// 删除菜单
export async function deleteMenu(params: any) {
  return axiosReq({
    url: '/menu/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除菜单
export async function batchDeleteMenu(params: any) {
  return axiosReq({
    url: '/menu/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新菜单状态
export async function updateMenuStatus(params: any) {
  return axiosReq({
    url: '/menu/updateStatus',
    method: 'put',
    data: params,
  })
}
