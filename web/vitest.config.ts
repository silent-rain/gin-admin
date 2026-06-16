import { defineConfig } from 'vitest/config';
import Vue from '@vitejs/plugin-vue';
import VueJsx from '@vitejs/plugin-vue-jsx';
import { fileURLToPath, URL } from 'node:url';

export default defineConfig({
  plugins: [Vue(), VueJsx()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  test: {
    // 启用类似 Jest 的全局测试 API
    globals: true,

    // 清除 mock 调用、实例和结果
    clearMocks: true,

    // 测试环境
    environment: 'jsdom',

    // 环境配置
    environmentOptions: {
      jsdom: {
        // 模拟浏览器 API
      },
    },

    // 测试文件匹配模式
    include: ['src/**/*.{test,spec}.{js,ts,jsx,tsx}'],
    exclude: ['node_modules', 'dist', '.idea', '.vscode', '*.config.*'],

    // 设置文件
    setupFiles: ['./vitest.setup.ts'],

    // 覆盖率配置
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html'],
      include: ['src/**/*.{js,ts,jsx,tsx,vue}'],
      exclude: [
        'node_modules',
        'dist',
        '**/*.config.*',
        '**/*.d.ts',
        '**/*.test.*',
        '**/*.spec.*',
        '**/__tests__/**',
        'src/typings/**',
        'src/types/**',
      ],
      thresholds: {
        lines: 80,
        functions: 80,
        branches: 80,
        statements: 80,
      },
    },

    // 依赖优化
    deps: {
      optimizer: {
        web: {
          enabled: true,
        },
      },
    },

    // 转换模式
    transformMode: {
      web: [/\.[jt]sx$/],
    },

    // 服务器配置
    server: {
      deps: {
        inline: ['@vue/test-utils'],
      },
    },

    // 输出配置
    outputFile: {
      json: './test-results.json',
    },

    // 其他配置
    mockReset: true,
    restoreMocks: true,
    unstubEnvs: true,
    unstubGlobals: true,
  },
});
