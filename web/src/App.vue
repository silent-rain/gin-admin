<script setup lang="ts">
import en from 'element-plus/dist/locale/en.mjs'
// element-plus lang
import zh from 'element-plus/dist/locale/zh-cn.mjs'
import { storeToRefs } from 'pinia'
import { onBeforeMount, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useErrorLog } from '@/hooks/use-error-log'
import { useBasicStore } from '@/store/basic'
import { useConfigStore } from '@/store/config'
import { useUserStore } from '@/store/user'
// reshow default setting
import { toggleHtmlClass } from '@/theme/utils'

const { settings } = storeToRefs(useBasicStore())
const { size, language } = storeToRefs(useConfigStore())
const route = useRoute()

const lang = { zh, en }

onBeforeMount(() => {
  // set tmp token when setting isNeedLogin false
  if (!settings.value.isNeedLogin) {
    useUserStore().setToken(settings.value.tmpToken)
  }
})
onMounted(() => {
  // lanch the errorLog collection
  useErrorLog()
})
onMounted(() => {
  const { setTheme, theme, setSize, size, setLanguage, language }
    = useConfigStore()
  setTheme(theme)
  setLanguage(language, route.meta?.title)
  setSize(size)
  toggleHtmlClass(theme)
})
</script>

<template>
  <el-config-provider :locale="lang[language]" namespace="el" :size="size">
    <router-view />
  </el-config-provider>
</template>

<style lang="scss">
//修改进度条样式
#nprogress .bar {
  background: var(--pregress-bar-color) !important;
}
</style>
