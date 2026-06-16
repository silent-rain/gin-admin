import axios from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';
let reqConfig;
let loadingE;
const service = axios.create();
service.interceptors.request.use((request) => {
    const { axiosPromiseArr } = useBasicStore();
    request.cancelToken = new axios.CancelToken((cancel) => {
        axiosPromiseArr.push({
            url: request.url,
            cancel,
        });
    });
    request.headers.authorization = useUserStore().token;
    if (request.isDownLoadFile) {
        request.responseType = 'blob';
    }
    if (request.isUploadFile) {
        request.headers['Content-Type'] = 'multipart/form-data';
    }
    reqConfig = request;
    if (request.bfLoading) {
        loadingE = ElLoading.service({
            lock: true,
            text: '数据载入中',
            background: 'rgba(0, 0, 0, 0.1)',
        });
    }
    if (request.isParams) {
        request.params = request.data;
        request.data = {};
    }
    return request;
}, (err) => {
    Promise.reject(err);
});
service.interceptors.response.use((res) => {
    if (reqConfig.afHLoading && loadingE) {
        loadingE.close();
    }
    if (reqConfig.isDownLoadFile) {
        return res;
    }
    const { msg, isNeedUpdateToken, data, code } = res.data;
    if (isNeedUpdateToken && data.token) {
    }
    const successCode = '0,200,10000';
    if (successCode.includes(code)) {
        return res.data;
    }
    const noAuthCode = '401,403,10400,10401,10402,10403,10404,10405,10406,10407';
    if (noAuthCode.includes(code) && !location.href.includes('/login')) {
        ElMessageBox.confirm(res.data.msg ? res.data.msg : '请重新登录', {
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
    return Promise.reject(res.data);
}, (err) => {
    if (loadingE) {
        loadingE.close();
    }
    ElMessage.error({
        message: err,
        duration: 2 * 1000,
    });
    const errObj = {
        msg: err.toString(),
        reqUrl: reqConfig.baseURL + reqConfig.url,
        params: reqConfig.isParams ? reqConfig.params : reqConfig.data,
    };
    return Promise.reject(JSON.stringify(errObj));
});
export function axiosReq2(config) {
    return service({
        baseURL: import.meta.env.VITE_APP_BASE_URL,
        timeout: 8000,
        ...config,
    });
}
export default function axiosReq({ url, method, data = {}, isParams = false, bfLoading = false, afHLoading = false, isUploadFile = false, isDownLoadFile = false, baseURL = import.meta.env.VITE_APP_BASE_URL, timeout = 150000, isAlertErrorMsg = true, }) {
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
function AxiosHeaderValue() {
    throw new Error('Function not implemented.');
}
