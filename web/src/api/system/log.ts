/** 日志
 *
 */
import axiosReq from '@/utils/axios-req';

// 获取网络请求日志列表
export const getHttpLogList = async (params: any) => {
  return axiosReq({
    url: '/httpLog/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 获取系统日志列表
export const getSystemLogList = async (params: any) => {
  return axiosReq({
    url: '/systemLog/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};
