import axios, { AxiosRequestConfig } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';

let reqConfig: any;
let loadingE: any;

// 使用axios.create()创建一个axios请求实例
const service = axios.create();

// 请求拦截
// @ts-ignore
service.interceptors.request.use(
  (request) => {
    const { axiosPromiseArr } = useBasicStore();
    // axiosPromiseArr收集请求地址,用于取消请求
    request.cancelToken = new axios.CancelToken((cancel) => {
      axiosPromiseArr.push({
        url: request.url,
        cancel,
      });
    });

    // token setting
    // @ts-ignore
    request.headers.authorization = useUserStore().token;
    /* download file */
    // @ts-ignore
    if (request.isDownLoadFile) {
      request.responseType = 'blob';
    }
    /* upload file */
    // @ts-ignore
    if (request.isUploadFile) {
      // @ts-ignore
      request.headers['Content-Type'] = 'multipart/form-data';
    }

    reqConfig = request;
    // @ts-ignore
    if (request.bfLoading) {
      // @ts-ignore
      loadingE = ElLoading.service({
        lock: true,
        text: '数据载入中',
        // spinner: 'el-icon-ElLoading',
        background: 'rgba(0, 0, 0, 0.1)',
      });
    }

    // params会拼接到url上
    // @ts-ignore
    if (request.isParams) {
      request.params = request.data;
      request.data = {};
    }
    return request;
  },
  (err) => {
    // 发送请求失败
    Promise.reject(err);
  },
);

// 请求后拦截
service.interceptors.response.use(
  (res) => {
    if (reqConfig.afHLoading && loadingE) {
      loadingE.close();
    }

    // 如果是下载文件直接返回
    if (reqConfig.isDownLoadFile) {
      return res;
    }

    const { msg, isNeedUpdateToken, data, code } = res.data;
    // 更新token保持登录状态
    if (isNeedUpdateToken && data.token) {
      // setToken(data.token);
    }

    const successCode = '0,200,10000';
    if (successCode.includes(code)) {
      return res.data;
    }

    const noAuthCode = '401,403,10401,10402,10403,10404,10405,10406';
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

    if (reqConfig.isAlertErrorMsg) {
      ElMessage({
        message: msg,
        type: 'error',
        duration: 2 * 1000,
      });
    }

    // 返回错误信息
    // 如果未catch 走unhandledrejection进行收集
    // 注：如果没有return 则，会放回到请求方法中.then ,返回的res为 undefined
    return Promise.reject(res.data);
  },
  // 响应报错
  (err) => {
    if (loadingE) {
      loadingE.close();
    }
    ElMessage.error({
      message: err,
      duration: 2 * 1000,
    });

    // 如果是跨域
    // Network Error,cross origin
    const errObj = {
      msg: err.toString(),
      reqUrl: reqConfig.baseURL + reqConfig.url,
      params: reqConfig.isParams ? reqConfig.params : reqConfig.data,
    };
    return Promise.reject(JSON.stringify(errObj));
  },
);

// 导出service实例给页面调用 , config->页面的配置
export function axiosReq2(config: AxiosRequestConfig<any>) {
  return service({
    baseURL: import.meta.env.VITE_APP_BASE_URL,
    timeout: 8000,
    ...config,
  });
}

// 导出service实例给页面调用, 自定义配置
export default function axiosReq({
  url,
  method,
  data = {},
  isParams = false,
  bfLoading = false,
  afHLoading = false,
  isUploadFile = false,
  isDownLoadFile = false,
  baseURL = import.meta.env.VITE_APP_BASE_URL,
  timeout = 150000,
  isAlertErrorMsg = true,
}) {
  // @ts-ignore
  return service({
    url,
    method,
    data,
    isParams,
    bfLoading,
    afHLoading,
    isUploadFile,
    isDownLoadFile,
    isAlertErrorMsg,
    baseURL,
    timeout,
  });
}

function AxiosHeaderValue(): any {
  throw new Error('Function not implemented.');
}
