import { WebErrorType } from '@/constant/system/log';
import axiosReq from '@/utils/axios-req';
import { useBasicStore } from '@/store/basic';
export const getHttpLogList = async (params) => {
    return axiosReq({
        url: '/httpLog/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getHttpLogBody = async (params) => {
    return axiosReq({
        url: '/httpLog/body',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getSystemLogList = async (params) => {
    return axiosReq({
        url: '/systemLog/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getWebList = async (params) => {
    return axiosReq({
        url: '/webLog/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addWebCodeLog = async (params) => {
    if (!params.level) {
        params.level = 'ERROR';
    }
    params.error_type = WebErrorType.Code;
    params.os_type = useBasicStore().osType;
    params.url = window.location.href;
    return axiosReq({
        url: '/webLog/add',
        method: 'post',
        data: params,
    });
};
export const addWebApiLog = async (params) => {
    if (!params.level) {
        params.level = 'ERROR';
    }
    params.error_type = WebErrorType.Api;
    params.os_type = useBasicStore().osType;
    params.url = window.location.href;
    return axiosReq({
        url: '/webLog/add',
        method: 'post',
        data: params,
    });
};
