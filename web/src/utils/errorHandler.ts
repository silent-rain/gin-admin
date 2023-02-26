/**web 端错误日志 */
import { useBasicStore } from '@/store/basic';
import { addWebCodeLog } from '@/api/system/log';

// vue 警告监听
export const warnHandler = async (msg: any, vm: any, trace: string) => {
  console.warn(msg);
  console.warn(trace);
  addWebCodeLog({
    level: 'WARN',
    caller_line: vm.$.vnode.type.__file,
    msg: JSON.stringify(msg),
    stack: trace,
  });
};
/**vue 错误监听
 * Vue全局错误监听，所有组件错误都会汇总到这里
 * errorCaptured 返回 false ，错误会被提前拦截阻止，这里无法捕获
 * err：具体错误信息
 * vm：当前错误所在的Vue实例
 * info：Vue特定的错误信息，错误所在的生命周期钩子
 */
export const errorHandler = async (msg: any, vm: any, trace: string) => {
  console.error(msg);
  console.error(trace);
  addWebCodeLog({
    level: 'ERROR',
    caller_line: vm.$.vnode.type.__file,
    msg: JSON.stringify(msg),
    stack: trace,
  });
};

/**JS全局onerror
 * 全局监听所有JS错误
 * 无法识别 Vue 组件信息
 * 可以捕获一些 Vue 监听不到的错误，如：异步错误
 *
 * event – 人类可读的错误信息，以字符串的形式描述问题。
 * source – 产生错误的脚本文件的URL。它也是一个字符串的形式。
 * lineno – 产生错误的脚本文件的行号。它是整数格式。
 * colon – 产生错误的脚本文件的列号。它是整数格式的。
 * error – 这是被抛出的错误对象。
 */
export const onerrorHandler = async (
  message: string | Event,
  source?: string,
  lineno?: number,
  colno?: number,
  error?: any,
) => {
  console.log(message, source, lineno, colno, error);
  addWebCodeLog({
    level: 'ERROR',
    caller_line: source + ':' + lineno + ':' + colno,
    msg: JSON.stringify(message),
    stack: JSON.stringify(error),
  });
};
