import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import ElementPlus from 'element-plus';

import VXETable from 'vxe-table';
import App from './App.vue';
import router from './router';

// import theme
import './theme/index.scss';

// import unocss
import 'uno.css';

// i18n
import { setupI18n } from '@/lang';

import '@/styles/index.scss'; // global css

// svg-icon
import 'virtual:svg-icons-register';
import svgIcon from '@/icons/SvgIcon.vue';
import directive from '@/directives';

// import router intercept
import './permission';

// import element-plus
import 'element-plus/dist/index.css';

// import vxe-table
import 'vxe-table/lib/style.css';

// highlight 的样式，依赖包，组件
import 'highlight.js/styles/atom-one-dark.css';
import 'highlight.js/lib/common';
import hljsVuePlugin from '@highlightjs/vue-plugin';

const app = createApp(App);

// router
app.use(router);

// pinia
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
app.use(pinia);

// i18n
app.use(setupI18n);
app.component('SvgIcon', svgIcon);
directive(app);

// element-plus
app.use(ElementPlus, { size: 'small', zIndex: 3000 });
app.use(VXETable);

// highlight;
app.use(hljsVuePlugin);

app.mount('#app');
