import axiosReq from '@/utils/axios-req';
export const getAllMenuTree = async () => {
    return axiosReq({
        url: '/menu/allTree',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getMenuTree = async (params) => {
    return axiosReq({
        url: '/menu/tree',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addMenu = async (params) => {
    return axiosReq({
        url: '/menu/add',
        method: 'post',
        data: params,
    });
};
export const updateMenu = async (params) => {
    return axiosReq({
        url: '/menu/update',
        method: 'put',
        data: params,
    });
};
export const deleteMenu = async (params) => {
    return axiosReq({
        url: '/menu/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteMenu = async (params) => {
    return axiosReq({
        url: '/menu/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateMenuStatus = async (params) => {
    return axiosReq({
        url: '/menu/updateStatus',
        method: 'put',
        data: params,
    });
};
