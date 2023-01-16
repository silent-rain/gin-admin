import { nextTick } from 'vue';
import { defineStore } from 'pinia';
import router from '@/router';
import { getUserInfo } from '@/api/system/user';

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      // 令牌
      token: '',
      // 登录标识
      getUserInfo: false,
      // user info
      userInfo: { username: '', avatar: '' },
      roles: [] as Array<number>,
      codes: [] as Array<number>,
    };
  },
  persist: {
    storage: localStorage,
    paths: ['token'],
  },
  actions: {
    // 获取用户信息
    async userInfo() {
      const userData = (await getUserInfo({})).data;
      const roles = userData.roles.map((v: any) => {
        return v.id;
      });
      return {
        userInfo: userData,
        roles: roles,
        codes: [],
        menuList: [],
      };
    },

    // 设置 token
    setToken(data: string) {
      this.token = data;
    },
    // 设置用户状态
    setUserInfo({ userInfo, roles, codes }) {
      const { nickname, avatar } = userInfo;
      this.$patch((state) => {
        state.roles = roles;
        state.codes = codes;
        state.getUserInfo = true;
        state.userInfo.username = nickname;
        state.userInfo.avatar = avatar;
      });
    },
    // 重置登录状态
    resetState() {
      this.$patch((state) => {
        // reset token
        state.token = '';

        state.roles = [];
        state.codes = [];

        // reset router

        // reset userInfo
        state.userInfo.username = '';
        state.userInfo.avatar = '';
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
