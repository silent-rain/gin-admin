/*
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 02:06:57
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/hooks/use-common.ts
 * @Descripttion:
 */
// 复制文本
import useClipboard from 'vue-clipboard3';
import { ElMessage } from 'element-plus';

// i18n language  match title
import { i18n } from '@/lang';
// the keys using  zh file
import langEn from '@/lang/zh';
import settings from '@/settings';

export const sleepTimeout = (time: number) => {
  return new Promise((resolve) => {
    const timer = setTimeout(() => {
      clearTimeout(timer);
      resolve(null);
    }, time);
  });
};

// 深拷贝
export function cloneDeep(value) {
  return JSON.parse(JSON.stringify(value));
}

// copyValueToClipboard
const { toClipboard } = useClipboard();
export const copyValueToClipboard = (value: any) => {
  toClipboard(JSON.stringify(value));
  ElMessage.success('复制成功');
};
const { t, te } = i18n.global;
export const langTitle = (title) => {
  if (!title) {
    return settings.title;
  }
  for (const key of Object.keys(langEn)) {
    if (te(`${key}.${title}`) && t(`${key}.${title}`)) {
      return t(`${key}.${title}`);
    }
  }
  return title;
};

// get i18n instance
export const getLangInstance = () => {
  return i18n.global as ObjKeys;
};
