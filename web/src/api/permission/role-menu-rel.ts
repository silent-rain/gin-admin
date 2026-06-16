/** 角色菜单关系 */
import axiosReq from '@/utils/axios-req'

// 获取角色菜单列表
export async function getRoleMenuRelList(params: any) {
  return axiosReq({
    url: '/roleMenuRel/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 更新角色菜单关联关系
export async function updateRoleMenuRel(params: any) {
  return axiosReq({
    url: '/roleMenuRel/update',
    method: 'put',
    data: params,
  })
}
