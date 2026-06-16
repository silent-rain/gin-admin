/**
 * 角色
 *
 */
import axiosReq from '@/utils/axios-req'

// 获取所有角色列表
export async function getAllRole() {
  return axiosReq({
    url: '/role/all',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取角色列表
export async function getRoleList(params: any) {
  return axiosReq({
    url: '/role/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加角色
export async function addRole(params: any) {
  return axiosReq({
    url: '/role/add',
    method: 'post',
    data: params,
  })
}

// 更新角色
export async function updateRole(params: any) {
  return axiosReq({
    url: '/role/update',
    method: 'put',
    data: params,
  })
}

// 删除角色
export async function deleteRole(params: any) {
  return axiosReq({
    url: '/role/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除角色
export async function batchDeleteRole(params: any) {
  return axiosReq({
    url: '/role/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新角色状态
export async function updateRoleStatus(params: any) {
  return axiosReq({
    url: '/role/updateStatus',
    method: 'put',
    data: params,
  })
}
