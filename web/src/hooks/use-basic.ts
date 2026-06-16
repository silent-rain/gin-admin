import type { ConfigListRsp } from '@/typings/api/data-center/config'
import NProgress from 'nprogress' // 进度条
import { getWebSiteConfigList } from '@/api/data-center/config'
import 'nprogress/nprogress.css' // 进度条样式

NProgress.configure({ showSpinner: false })
// 开始进度条
export function progressStart() {
  NProgress.start()
}
// 关闭进度条
export function progressClose() {
  NProgress.done()
}

// 查询网站配置列表
export async function fecthWebSiteConfigList() {
  const basicStore = useBasicStore()
  const configHash = {}
  try {
    const resp = (await getWebSiteConfigList()).data as ConfigListRsp
    for (const item of resp.data_list) {
      configHash[item.key] = item
    }
    basicStore.setWebSiteConfig(configHash)
  }
  catch (error) {
    console.log(error)
  }
}
