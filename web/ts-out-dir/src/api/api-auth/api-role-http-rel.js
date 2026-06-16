import axiosReq from '@/utils/axios-req';
export const getApiRoleHttpRelList = async (params) => {
    return axiosReq({
        url: '/apiRoleHttpRel/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const updateApiRoleHttpRel = async (params) => {
    return axiosReq({
        url: '/apiRoleHttpRel/update',
        method: 'put',
        data: params,
    });
};
