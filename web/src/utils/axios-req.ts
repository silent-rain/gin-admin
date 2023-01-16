/*
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 01:39:59
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/utils/axios-req.ts
 * @Descripttion:
 */
import axios, { AxiosRequestConfig } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';

// 使用axios.create()创建一个axios请求实例
const service = axios.create();

// 请求前拦截
service.interceptors.request.use(
  (req) => {
    const { axiosPromiseArr } = useBasicStore();
    const { token } = useUserStore();
    // axiosPromiseArr收集请求地址,用于取消请求
    req.cancelToken = new axios.CancelToken((cancel) => {
      axiosPromiseArr.push({
        url: req.url,
        cancel,
      });
    });
    // 设置token到header
    // @ts-ignore
    req.headers['Authorization'] = token;
    // @ts-ignore
    req.headers['Content-type'] = 'application/json;charset=UTF-8';
    // 如果req.method给get 请求参数设置为 ?name=xxx
    if ('get'.includes(req.method?.toLowerCase() as string))
      req.params = req.data;
    return req;
  },
  (err) => {
    // 发送请求失败
    Promise.reject(err);
  },
);

// 请求后拦截
service.interceptors.response.use(
  (res) => {
    const { code } = res.data;
    const successCode = '0,200,10000,20000';
    const noAuthCode = '401,403,10401,10402,10403,10404,10405,10406';
    if (successCode.includes(code)) {
      return res.data;
    }
    if (noAuthCode.includes(code) && !location.href.includes('/login')) {
      ElMessageBox.confirm('请重新登录', {
        confirmButtonText: '重新登录',
        closeOnClickModal: false,
        showCancelButton: false,
        showClose: false,
        type: 'warning',
      }).then(() => {
        useUserStore().resetStateAndToLogin();
      });
    }
    return Promise.reject(res.data);
  },
  // 响应报错
  (err) => {
    ElMessage.error({
      message: err,
      duration: 2 * 1000,
    });
    return Promise.reject(err);
  },
);

// 导出service实例给页面调用 , config->页面的配置
export default function axiosReq(config: AxiosRequestConfig<any>) {
  return service({
    baseURL: import.meta.env.VITE_APP_BASE_URL,
    timeout: 8000,
    ...config,
  });
}
function AxiosHeaderValue(): any {
  throw new Error('Function not implemented.');
}
