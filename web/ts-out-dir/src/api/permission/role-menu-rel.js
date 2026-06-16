import axiosReq from '@/utils/axios-req';
export const getRoleMenuRelList = async (params) => {
    return axiosReq({
        url: '/roleMenuRel/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const updateRoleMenuRel = async (params) => {
    return axiosReq({
        url: '/roleMenuRel/update',
        method: 'put',
        data: params,
    });
};
