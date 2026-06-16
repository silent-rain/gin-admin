import { AxiosRequestConfig } from 'axios';
export declare function axiosReq2(config: AxiosRequestConfig<any>): Promise<import("axios").AxiosResponse<any, any, {}>>;
export default function axiosReq({ url, method, data, isParams, bfLoading, afHLoading, isUploadFile, isDownLoadFile, baseURL, timeout, isAlertErrorMsg, }: {
    url: any;
    method: any;
    data?: {} | undefined;
    isParams?: boolean | undefined;
    bfLoading?: boolean | undefined;
    afHLoading?: boolean | undefined;
    isUploadFile?: boolean | undefined;
    isDownLoadFile?: boolean | undefined;
    baseURL?: string | undefined;
    timeout?: number | undefined;
    isAlertErrorMsg?: boolean | undefined;
}): Promise<import("axios").AxiosResponse<any, any, {}>>;
