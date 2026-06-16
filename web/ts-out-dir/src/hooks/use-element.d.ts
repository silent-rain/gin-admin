import type { EpPropMergeType } from 'element-plus/es/utils';
export declare const useElement: () => {
    tableData: globalThis.Ref<never[], never[]>;
    rowDeleteIdArr: globalThis.Ref<never[], never[]>;
    loadingId: globalThis.Ref<null, null>;
    formModel: globalThis.Ref<{}, {}>;
    subForm: globalThis.Ref<{}, {}>;
    searchForm: globalThis.Ref<{}, {}>;
    formRules: globalThis.Ref<{
        isNull: (msg: string) => {
            required: boolean;
            message: string;
            trigger: string;
        }[];
        isNotNull: (msg: string) => {
            required: boolean;
            message: string;
            trigger: string;
        }[];
        upZeroInt: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        zeroInt: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        money: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        phone: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        email: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
    }, {
        isNull: (msg: string) => {
            required: boolean;
            message: string;
            trigger: string;
        }[];
        isNotNull: (msg: string) => {
            required: boolean;
            message: string;
            trigger: string;
        }[];
        upZeroInt: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        zeroInt: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        money: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        phone: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
        email: (msg: string) => {
            required: boolean;
            validator: (rule: any, value: any, callback: any) => void;
            trigger: string;
        }[];
    }>;
    datePickerOptions: globalThis.Ref<{
        disabledDate: (time: any) => boolean;
    }, {
        disabledDate: (time: any) => boolean;
    }>;
    startEndArr: globalThis.Ref<never[], never[]>;
    dialogTitle: globalThis.Ref<string, string>;
    detailDialog: globalThis.Ref<boolean, boolean>;
    isDialogEdit: globalThis.Ref<boolean, boolean>;
    dialogVisible: globalThis.Ref<boolean, boolean>;
    tableLoading: globalThis.Ref<boolean, boolean>;
    treeData: globalThis.Ref<never[], never[]>;
    defaultProps: globalThis.Ref<{
        children: string;
        label: string;
    }, {
        children: string;
        label: string;
    }>;
};
export declare const elMessage: (message: string, type?: any) => void;
export declare const elLoading: (msg?: string) => void;
export declare const closeElLoading: () => void;
export declare const elNotify: (message: string, type: EpPropMergeType<any, any, any> | undefined, title: string, duration: number) => void;
export declare const elConfirmNoCancelBtn: (title: string, message: string) => Promise<import("element-plus").MessageBoxData>;
export declare const elConfirm: (title: string, message: string) => Promise<import("element-plus").MessageBoxData>;
export declare const casHandleChange: () => void;
