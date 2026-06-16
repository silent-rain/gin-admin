import axiosReq from '@/utils/axios-req';
export const getAllRole = async () => {
    return axiosReq({
        url: '/role/all',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getRoleList = async (params) => {
    return axiosReq({
        url: '/role/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addRole = async (params) => {
    return axiosReq({
        url: '/role/add',
        method: 'post',
        data: params,
    });
};
export const updateRole = async (params) => {
    return axiosReq({
        url: '/role/update',
        method: 'put',
        data: params,
    });
};
export const deleteRole = async (params) => {
    return axiosReq({
        url: '/role/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteRole = async (params) => {
    return axiosReq({
        url: '/role/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateRoleStatus = async (params) => {
    return axiosReq({
        url: '/role/updateStatus',
        method: 'put',
        data: params,
    });
};
