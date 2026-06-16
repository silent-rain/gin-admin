import type { RouteRawConfig, RouterTypes } from '~/store/router';
import 'nprogress/nprogress.css';
interface menuRow {
    category: number;
    code: number;
    children: RouterTypes;
}
export declare const filterAsyncRoutesByMenuList: (menuList: menuRow[]) => RouterTypes;
export declare function filterAsyncRoutesByRoles(routes: RouteRawConfig[], roles: number[]): RouterTypes;
export declare function filterAsyncRouterByCodes(codesRoutes: RouteRawConfig[], codes: number[]): RouterTypes;
export {};
