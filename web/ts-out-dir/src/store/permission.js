import { defineStore } from 'pinia';
export const usePermissionStore = defineStore('permission', {
    state: () => {
        return {
            menus: [],
            permissions: [],
            permissionHash: {},
            asyncRoutes: [],
            allRoutes: [],
        };
    },
    actions: {
        setFilterAsyncRoutes(menus, asyncRoutes, allRoutes) {
            this.$patch((state) => {
                state.menus = menus;
                state.asyncRoutes = asyncRoutes;
                state.allRoutes = allRoutes;
            });
        },
        setButtonPermission(permissions, permissionHash) {
            this.$patch((state) => {
                state.permissions = permissions;
                state.permissionHash = permissionHash;
            });
        },
        resetState() {
            this.$patch((state) => {
                state.menus = [];
                state.allRoutes = [];
                state.permissions = [];
                state.asyncRoutes = [];
            });
        },
    },
});
