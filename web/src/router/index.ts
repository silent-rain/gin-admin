import { createRouter, createWebHashHistory } from 'vue-router';
import Layout from '@/layout/index.vue';
import type { RouteRawConfig, RouterTypes } from '~/basic';
import basicDemo from './modules/basic-demo';
import charts from './modules/charts';
import richText from './modules/rich-text';
import table from './modules/table';
import excel from './modules/excel';
import directive from './modules/directive';
import other from './modules/other';
import guid from './modules/guid';

// 系统固定公开路由
export const constantRoutes: RouterTypes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect'),
      },
    ],
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404.vue'),
    hidden: true,
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401.vue'),
    hidden: true,
  },
  {
    path: '/login',
    component: () => import('@/views/login/index.vue'),
    hidden: true,
    meta: { title: '登录' },
  },
  {
    path: '/register',
    component: () => import('@/views/login/register.vue'),
    hidden: true,
    meta: { title: '用户注册' },
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        // using el svg icon, the elSvgIcon first when at the same time using elSvgIcon and icon
        meta: { title: 'Dashboard', elSvgIcon: 'Fold', affix: true },
      },
    ],
  },
  {
    path: '/system',
    component: Layout,
    meta: { title: '系统管理', elSvgIcon: 'Setting' },
    alwaysShow: true,
    children: [
      {
        path: '/system/user',
        component: () => import('@/views/system/user/index.vue'),
        // name: '用户管理',
        meta: { title: '用户管理', elSvgIcon: 'User' },
      },

      {
        path: '/system/menu',
        component: () => import('@/views/system/menu/index.vue'),
        name: '菜单管理',
        meta: { title: '菜单管理', elSvgIcon: 'Operation' },
      },
    ],
  },
];

export const catchRoutes: RouterTypes = [
  {
    hide: true,
    path: '/:pathMatch(.*)*',
    redirect: '/404',
  },
  {
    path: '/:catchAll(.*)',
    name: 'CatchAll',
    redirect: '/404',
    hidden: true,
  },
];

// export const constantRoutes2: RouterTypes = [
//   {
//     path: '/redirect',
//     component: Layout,
//     hidden: true,
//     children: [
//       {
//         path: '/redirect/:path(.*)',
//         component: () => import('@/views/redirect'),
//       },
//     ],
//   },
//   {
//     path: '/login',
//     component: () => import('@/views/login/index.vue'),
//     hidden: true,
//     meta: { title: '登录' },
//   },
//   {
//     path: '/register',
//     component: () => import('@/views/login/register.vue'),
//     hidden: true,
//     meta: { title: '用户注册' },
//   },
//   {
//     path: '/404',
//     component: () => import('@/views/error-page/404.vue'),
//     hidden: true,
//   },
//   {
//     path: '/401',
//     component: () => import('@/views/error-page/401.vue'),
//     hidden: true,
//   },
//   {
//     path: '/',
//     component: Layout,
//     redirect: '/dashboard',
//     children: [
//       {
//         path: 'dashboard',
//         name: 'Dashboard',
//         component: () => import('@/views/dashboard/index.vue'),
//         // using el svg icon, the elSvgIcon first when at the same time using elSvgIcon and icon
//         meta: { title: 'Dashboard', elSvgIcon: 'Fold', affix: true },
//       },
//     ],
//   },
//   {
//     path: '/system',
//     component: Layout,
//     meta: { title: '系统管理', elSvgIcon: 'Setting' },
//     alwaysShow: true,
//     children: [
//       {
//         path: '/system/user',
//         component: () => import('@/views/system/user/index.vue'),
//         name: '用户管理',
//         meta: { title: '用户管理', elSvgIcon: 'User' },
//       },
//       {
//         path: '/system/role',
//         component: () => import('@/views/system/role/index.vue'),
//         name: '角色管理',
//         meta: { title: '角色管理', elSvgIcon: 'Postcard' },
//       },
//       {
//         path: '/system/menu',
//         component: () => import('@/views/system/menu/index.vue'),
//         name: '菜单管理',
//         meta: { title: '菜单管理', elSvgIcon: 'Operation' },
//       },
//       {
//         path: '/system/level',
//         component: () => import('@/views/other/signboard/index.vue'),
//         name: '职级管理',
//         hidden: true,
//         meta: { title: '职级管理', elSvgIcon: 'Histogram' },
//       },
//       {
//         path: '/system/position',
//         component: () => import('@/views/other/signboard/index.vue'),
//         name: '岗位管理',
//         hidden: true,
//         meta: { title: '岗位管理', elSvgIcon: 'MessageBox' },
//       },
//       {
//         path: '/system/dept',
//         component: () => import('@/views/other/signboard/index.vue'),
//         name: '部门管理',
//         hidden: true,
//         meta: { title: '部门管理', elSvgIcon: 'OfficeBuilding' },
//       },
//     ],
//   },
//   guid,
//   {
//     path: '/RBAC',
//     component: Layout,
//     children: [
//       {
//         path: 'https://github.jzfai.top/low-code-plateform/#/permission-center/user-table-query',
//         meta: { title: 'RBAC', icon: 'skill' },
//       },
//     ],
//   },
//   basicDemo,
//   richText,
//   charts,
//   table,
//   directive,
//   excel,
//   other,
//   {
//     path: '/setting-switch',
//     component: Layout,
//     children: [
//       {
//         path: 'index',
//         component: () => import('@/views/setting-switch/index.vue'),
//         name: 'SettingSwitch',
//         meta: { title: 'Setting Switch', icon: 'example' },
//       },
//     ],
//   },
//   {
//     path: '/error-log',
//     component: Layout,
//     meta: { title: 'Error Log', icon: 'eye' },
//     alwaysShow: true,
//     children: [
//       {
//         path: 'error-log',
//         component: () => import('@/views/error-log/index.vue'),
//         name: 'ErrorLog',
//         meta: { title: 'Error Index' },
//       },
//       {
//         path: 'error-generator',
//         component: () => import('@/views/error-log/error-generator.vue'),
//         name: 'ErrorGenerator',
//         meta: { title: 'Error Generator' },
//       },
//     ],
//   },
//   {
//     path: '/nested',
//     component: Layout,
//     redirect: '/nested/menu1',
//     name: 'Nested',
//     meta: {
//       title: 'Nested',
//       icon: 'nested',
//     },
//     children: [
//       {
//         path: 'menu1',
//         component: () => import('@/views/nested/menu1/index.vue'), // Parent router-view
//         name: 'Menu1',
//         meta: { title: 'Menu1' },
//         children: [
//           {
//             path: 'menu1-1',
//             component: () => import('@/views/nested/menu1/menu1-1/index.vue'),
//             name: 'Menu1-1',
//             meta: { title: 'Menu1-1' },
//           },
//           {
//             path: 'menu1-2',
//             component: () => import('@/views/nested/menu1/menu1-2/index.vue'),
//             name: 'Menu1-2',
//             meta: { title: 'Menu1-2' },
//             children: [
//               {
//                 path: 'menu1-2-1',
//                 component: () =>
//                   import('@/views/nested/menu1/menu1-2/menu1-2-1/index.vue'),
//                 name: 'Menu1-2-1',
//                 meta: { title: 'Menu1-2-1' },
//               },
//               {
//                 path: 'menu1-2-2',
//                 component: () =>
//                   import('@/views/nested/menu1/menu1-2/menu1-2-2/index.vue'),
//                 name: 'Menu1-2-2',
//                 meta: { title: 'Menu1-2-2' },
//               },
//             ],
//           },
//           {
//             path: 'menu1-3',
//             component: () => import('@/views/nested/menu1/menu1-3/index.vue'),
//             name: 'Menu1-3',
//             meta: { title: 'Menu1-3' },
//           },
//         ],
//       },
//       {
//         path: 'menu2',
//         component: () => import('@/views/nested/menu2/index.vue'),
//         name: 'Menu2',
//         meta: { title: 'menu2' },
//       },
//     ],
//   },
// ];

// 角色和code数组动态路由
// export const roleCodeRoutes: RouterTypes = [
//   {
//     path: '/roles-codes',
//     component: Layout,
//     redirect: '/roles-codes/page',
//     alwaysShow: true, // will always show the root menu
//     name: 'Permission',
//     meta: {
//       title: 'Permission',
//       icon: 'lock',
//       roles: ['admin', 'editor'], // you can set roles in root nav
//     },
//     children: [
//       {
//         path: 'index',
//         component: () => import('@/views/roles-codes/index.vue'),
//         name: 'RolesCodes',
//         meta: { title: 'Permission Switch' },
//       },
//       {
//         path: 'roleIndex',
//         component: () => import('@/views/roles-codes/role-index.vue'),
//         name: 'RoleIndex',
//         meta: { title: 'Role Index', roles: ['admin'] },
//       },
//       {
//         path: 'code-index',
//         component: () => import('@/views/roles-codes/code-index.vue'),
//         name: 'CodeIndex',
//         meta: { title: 'Code Index', code: 16 },
//       },
//       {
//         path: 'button-permission',
//         component: () => import('@/views/roles-codes/button-permission.vue'),
//         name: 'ButtonPermission',
//         meta: { title: 'Button Permission' },
//       },
//     ],
//   },
// ];

// 404 page must be placed at the end !!!

const router = createRouter({
  history: createWebHashHistory(),
  scrollBehavior: () => ({ top: 0 }),
  routes: constantRoutes,
});

export default router;
