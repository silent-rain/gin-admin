import { addWebCodeLog } from '@/api/system/log';
export const warnHandler = async (msg, vm, trace) => {
    console.warn(msg);
    console.warn(trace);
    addWebCodeLog({
        level: 'WARN',
        caller_line: vm.$.vnode.type.__file,
        msg: JSON.stringify(msg),
        stack: trace,
    });
};
export const errorHandler = async (msg, vm, trace) => {
    console.error(msg);
    console.error(trace);
    addWebCodeLog({
        level: 'ERROR',
        caller_line: vm.$.vnode.type.__file,
        msg: JSON.stringify(msg),
        stack: trace,
    });
};
export const onerrorHandler = async (message, source, lineno, colno, error) => {
    console.log(message, source, lineno, colno, error);
    addWebCodeLog({
        level: 'ERROR',
        caller_line: source + ':' + lineno + ':' + colno,
        msg: JSON.stringify(message),
        stack: JSON.stringify(error),
    });
};
