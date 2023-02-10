import { nextTick } from 'vue';
import { defineStore } from 'pinia';
import router from '@/router';
import { getUserInfo } from '@/api/system/user';
import { User } from '~/api/system/user';
import { Role } from '~/api/system/role';
import { UserTry } from '~/store/user';

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      // 令牌
      token: '',
      // 登录标识
      getUserInfo: false,
      // 用户信息
      userInfo: {} as User,
      // 角色列表
      roles: [] as Role[],
      codes: [] as number[],
    } as UserTry;
  },
  persist: {
    storage: localStorage,
    paths: ['token'],
  },
  actions: {
    // 获取用户信息
    async userInfo() {
      const userData = (await getUserInfo({})).data;
      return {
        userInfo: userData.user,
        roles: userData.roles,
        menus: userData.menus,
        permissions: userData.permissions,
        codes: [],
      };
    },

    // 设置 token
    setToken(data: string) {
      this.$patch((state) => {
        state.token = data;
      });
    },
    // 设置用户状态
    setUserInfo(data: any) {
      this.$patch((state) => {
        state.roles = data.roles;
        state.codes = data.codes;
        state.getUserInfo = true;
        state.userInfo.nickname = data.userInfo.nickname;
        state.userInfo.avatar = data.userInfo.avatar;
      });
    },
    // 重置登录状态
    resetState() {
      this.$patch((state) => {
        // reset token
        state.token = '';

        // reset userInfo
        state.userInfo = {} as User;

        state.roles = [];
        state.permissions = [];
        state.codes = [];

        // reset router
        // ...
      });
      this.getUserInfo = false;
    },
    // 重置登录状态, 并跳转到登录页面
    resetStateAndToLogin() {
      this.resetState();
      nextTick(() => {
        router.push({ path: '/login' });
      });
    },
  },
});
