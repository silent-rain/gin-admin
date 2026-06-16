import { RouteRawConfig, RouterTypes } from '~/store/router';
import { Menu, ButtonPermission } from '~/api/permission/menu';
export declare const usePermissionStore: import("pinia").StoreDefinition<"permission", {
    menus: Menu[];
    permissions: ButtonPermission[];
    permissionHash: {};
    asyncRoutes: RouteRawConfig[];
    allRoutes: RouterTypes;
}, {}, {
    setFilterAsyncRoutes(menus: Menu[], asyncRoutes: RouteRawConfig[], allRoutes: RouteRawConfig[]): void;
    setButtonPermission(permissions: ButtonPermission[], permissionHash: any): void;
    resetState(): void;
}>;
