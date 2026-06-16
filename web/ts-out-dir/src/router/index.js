import { createRouter, createWebHashHistory } from 'vue-router';
import Layout from '@/layout/index.vue';
export const constantRoutes = [
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
                meta: { title: 'Dashboard', elSvgIcon: 'Fold', affix: true },
            },
        ],
    },
];
export const catchRoutes = [
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
const router = createRouter({
    history: createWebHashHistory(),
    scrollBehavior: () => ({ top: 0 }),
    routes: constantRoutes,
});
export default router;
