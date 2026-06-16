export declare const useTable: (searchForm: any, selectPageReq: any) => {
    pageNum: globalThis.Ref<number, number>;
    pageSize: globalThis.Ref<number, number>;
    totalPage: globalThis.Ref<number, number>;
    tableListData: globalThis.Ref<never[], never[]>;
    tableListReq: (config: any) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    dateRangePacking: (timeArr: any) => void;
    multipleSelection: globalThis.Ref<ObjKeys[], ObjKeys[]>;
    handleSelectionChange: (val: any) => void;
    handleCurrentChange: (val: any) => void;
    handleSizeChange: (val: any) => void;
    resetPageReq: () => void;
    multiDelBtnDill: (reqConfig: any) => void;
    tableDelDill: (row: any, reqConfig: any) => void;
};
