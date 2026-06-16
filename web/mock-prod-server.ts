/// <reference types="vite/client" />

import type { MockMethod } from 'vite-plugin-mock';
import Mock from 'mockjs';

// 自动导入所有 mock 模块
const mockModules: MockMethod[] = [];

// 使用 Vite 的 glob 导入
const modules = import.meta.glob('./mock/*.ts', { eager: true });

for (const path in modules) {
  const module = modules[path] as { default: MockMethod[] };
  if (module.default) {
    mockModules.push(...module.default);
  }
}

export function setupProdMockServer() {
  if (mockModules.length > 0) {
    // 使用 mockjs 注册 mock 路由
    mockModules.forEach((mock) => {
      Mock.mock(mock.url, mock.method, mock.response);
    });
    console.log(`[Mock] ${mockModules.length} mock interfaces loaded`);
  }
}
