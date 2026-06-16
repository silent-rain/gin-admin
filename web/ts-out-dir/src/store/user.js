import { nextTick } from 'vue';
import { defineStore } from 'pinia';
import router from '@/router';
import { getUserInfo } from '@/api/permission/user';
import { TOKEN_PREFIX } from '@/constant/permission/auth';
export const useUserStore = defineStore('user', {
    state: () => {
        return {
            token: '',
            getUserInfo: false,
            userInfo: {},
            userId: undefined,
            userAvatar: undefined,
            roles: [],
            codes: [],
        };
    },
    persist: {
        storage: localStorage,
        pick: ['token'],
    },
    actions: {
        async userInfo() {
            const userData = (await getUserInfo()).data;
            return {
                userInfo: userData.user,
                roles: userData.roles,
                menus: userData.menus,
                permissions: userData.permissions,
                codes: [],
            };
        },
        setToken(data) {
            this.$patch((state) => {
                state.token = TOKEN_PREFIX + data;
            });
        },
        setUserInfo(data) {
            this.$patch((state) => {
                state.roles = data.roles;
                state.codes = data.codes;
                state.getUserInfo = true;
                state.userId = data.userInfo.id;
                if (data.userInfo.avatar) {
                    state.userAvatar =
                        import.meta.env.VITE_APP_IMAGE_URL + data.userInfo.avatar;
                }
                state.userInfo.nickname = data.userInfo.nickname;
                state.userInfo.avatar = data.userInfo.avatar;
            });
        },
        resetState() {
            this.$patch((state) => {
                state.token = '';
                state.userInfo = {};
                state.roles = [];
                state.codes = [];
            });
            this.getUserInfo = false;
        },
        resetStateAndToLogin() {
            this.resetState();
            nextTick(() => {
                router.push({ path: '/login' });
            });
        },
    },
});
