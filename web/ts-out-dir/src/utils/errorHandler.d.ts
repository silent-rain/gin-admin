export declare const warnHandler: (msg: any, vm: any, trace: string) => Promise<void>;
export declare const errorHandler: (msg: any, vm: any, trace: string) => Promise<void>;
export declare const onerrorHandler: (message: string | Event, source?: string, lineno?: number, colno?: number, error?: any) => Promise<void>;
