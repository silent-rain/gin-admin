import { ElMessage } from 'element-plus'
// 复制文本
import useClipboard from 'vue-clipboard3'

// i18n language  match title
import { i18n } from '@/lang'
// the keys using  zh file
import langEn from '@/lang/zh'
import settings from '@/settings'

export function sleepTimeout(time: number) {
  return new Promise((resolve) => {
    const timer = setTimeout(() => {
      clearTimeout(timer)
      resolve(null)
    }, time)
  })
}

// 深拷贝
export function cloneDeep(value) {
  return JSON.parse(JSON.stringify(value))
}

// copyValueToClipboard
const { toClipboard } = useClipboard()
export function copyValueToClipboard(value: any) {
  toClipboard(JSON.stringify(value))
  ElMessage.success('复制成功')
}
const { t, te } = i18n.global
export function langTitle(title) {
  if (!title) {
    return settings.title
  }
  for (const key of Object.keys(langEn)) {
    if (te(`${key}.${title}`) && t(`${key}.${title}`)) {
      return t(`${key}.${title}`)
    }
  }
  return title
}

// get i18n instance
export function getLangInstance() {
  return i18n.global as ObjKeys
}
