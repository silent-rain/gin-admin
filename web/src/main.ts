import hljsVuePlugin from '@highlightjs/vue-plugin'
import ElementPlus from 'element-plus'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import { createApp } from 'vue'
import VXETable from 'vxe-table'
import directive from '@/directives'

import svgIcon from '@/icons/SvgIcon.vue'

// i18n
import { setupI18n } from '@/lang'

import App from './App.vue'

import router from './router'

// error 日志上报
import * as errorLog from './utils/errorHandler'
// import theme
import './theme/index.scss'
// import unocss
import 'uno.css'

import '@/styles/index.scss' // global css

// svg-icon
import 'virtual:svg-icons-register'

// import router intercept
import './permission'

// import element-plus
import 'element-plus/dist/index.css'
// import vxe-table
import 'vxe-table/lib/style.css'
// highlight 的样式，依赖包，组件
import 'highlight.js/styles/atom-one-dark.css'

import 'highlight.js/lib/common'

const app = createApp(App)

// router
app.use(router)

// pinia
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

// i18n
app.use(setupI18n)
app.component('SvgIcon', svgIcon)
directive(app)

// element-plus
app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.use(VXETable)

// highlight;
app.use(hljsVuePlugin)

// vue error 错误日志监听上报
app.config.warnHandler = errorLog.warnHandler
app.config.errorHandler = errorLog.errorHandler
window.onerror = errorLog.onerrorHandler

app.mount('#app')
