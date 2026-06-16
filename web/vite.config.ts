import { resolve } from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons';
import { viteMockServe } from 'vite-plugin-mock';
import Components from 'unplugin-vue-components/vite';
import UnoCSS from 'unocss/vite';
import { presetAttributify, presetIcons, presetUno } from 'unocss';
import AutoImport from 'unplugin-auto-import/vite';
import setting from './src/settings';
import DefineOptions from 'unplugin-vue-define-options/vite';
// import { visualizer } from 'rollup-plugin-visualizer'

const prodMock = setting.openProdMock;
const pathSrc = resolve(__dirname, 'src');

export default defineConfig(({ command, mode }) => {
  const isProduction = mode === 'production';

  return {
    base: setting.viteBasePath,
    define: {
      // define global var
      GLOBAL_STRING: JSON.stringify(
        'i am global var from vite.config.js define',
      ),
      GLOBAL_VAR: { test: 'i am global var from vite.config.js define' },
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: isProduction,
    },
    clearScreen: false, // 设为 false 可以避免 Vite 清屏而错过在终端中打印某些关键信息
    server: {
      hmr: {
        overlay: false, // 设置 server.hmr.overlay 为 false 可以禁用开发服务器错误的屏蔽。方便错误查看
      },
      port: 5005, // 指定服务器端口
      open: false, // 在服务器启动时自动在浏览器中打开应用程序
      host: true,
      // https: { // 使用 Vite 内置 HTTPS 支持
      //   key: './localhost-key.pem',
      //   cert: './localhost.pem',
      // },
      watch: {
        usePolling: false, // 禁用轮询以提高性能
      },
    },
    preview: {
      port: 5006,
      host: true,
      strictPort: true,
    },
    plugins: [
      vue(),
      vueJsx(),
      DefineOptions(), // 使用 DefineOptions 替代 vitePluginSetupExtend
      UnoCSS({
        presets: [presetUno(), presetAttributify(), presetIcons()],
      }),
      createSvgIconsPlugin({
        iconDirs: [
          resolve(process.cwd(), 'src/icons/common'),
          resolve(process.cwd(), 'src/icons/nav-bar'),
        ],
        symbolId: 'icon-[dir]-[name]',
      }),
      // https://github.com/anncwb/vite-plugin-mock/blob/HEAD/README.zh_CN.md
      viteMockServe({
        supportTs: true,
        mockPath: 'mock',
        localEnabled: command === 'serve',
        prodEnabled: prodMock,
        injectCode: `
          import { setupProdMockServer } from '../utils/mock-prod-server';
          setupProdMockServer();
        `,
        logger: true,
      }),
      Components({
        dirs: ['src/components', 'src/icons'],
        extensions: ['vue'],
        deep: true,
        dts: './src/typings/components.d.ts',
      }),
      AutoImport({
        imports: ['vue', 'vue-router', 'pinia'],
        // 配置后会自动扫描目录下的文件
        dirs: [
          'src/hooks/**',
          'src/utils/**',
          'src/store/**',
          'src/api/**',
          'src/directives/**',
        ],
        dts: './src/typings/auto-imports.d.ts',
        vueTemplate: true, // 在 Vue 模板中自动导入
      }),
      // 依赖分析插件
      // visualizer({
      //   open: true,
      //   gzipSize: true,
      //   brotliSize: true
      // })
    ],
    build: {
      target: 'es2015', // 构建目标
      cssTarget: 'chrome61', // CSS 构建目标
      chunkSizeWarningLimit: 10000, // 消除触发警告的 chunk, 默认500k
      outDir: '../server/assets/dist', // 指定输出路径
      emptyOutDir: true, // 指定清空路径
      assetsDir: 'static/assets', // 指定生成静态文件目录
      minify: isProduction ? 'esbuild' : false, // 使用 esbuild 替代 terser 以获得更好的性能
      sourcemap: !isProduction, // 生产环境不生成 sourcemap
      rollupOptions: {
        output: {
          chunkFileNames: 'static/js/[name]-[hash].js',
          entryFileNames: 'static/js/[name]-[hash].js',
          assetFileNames: 'static/[ext]/[name]-[hash].[ext]',
          // Vite 8+ (Rolldown) 需要使用函数形式的 manualChunks
          manualChunks(id) {
            if (id.includes('node_modules')) {
              // Vue 相关库
              if (id.includes('vue') || id.includes('vue-router') || id.includes('pinia')) {
                return 'vue-vendor';
              }
              // UI 组件库
              if (id.includes('element-plus') || id.includes('@element-plus')) {
                return 'ui-vendor';
              }
            }
          },
        },
      },
    },
    resolve: {
      alias: {
        '@/': `${pathSrc}/`,
        // vue-i18n 新版本不需要这个别名
      },
    },
    optimizeDeps: {
      include: ['moment-mini', 'vue', 'vue-router', 'pinia'],
      exclude: ['vue-demi'],
    },
    esbuild: {
      drop: isProduction ? ['console', 'debugger'] : [],
      pure: isProduction ? ['console.log', 'console.warn'] : [],
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/styles/element/index.scss" as *;`,
        },
      },
      postcss: {
        plugins: [
          // 可以在这里添加 postcss 插件
        ],
      },
    },
  };
});
