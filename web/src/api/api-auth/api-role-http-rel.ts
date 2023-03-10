/** 角色与Http协议接口关联表 */
import axiosReq from '@/utils/axios-req';

// 获取角色与Http协议接口关系列表
export const getApiRoleHttpRelList = async (params: any) => {
  return axiosReq({
    url: '/apiRoleHttpRel/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 更新角色与Http协议接口关系
export const updateApiRoleHttpRel = async (params: any) => {
  return axiosReq({
    url: '/apiRoleHttpRel/update',
    method: 'put',
    data: params,
  });
};
