import type { defineOptions as _defineOptions } from 'unplugin-vue-define-options/macros.d.ts';

declare global {
  interface ObjKeys {
    [propName: string]: any;
  }
  const GLOBAL_VAR: string;
  const defineOptions: typeof _defineOptions;
  const $ref: any;
}
export {};
