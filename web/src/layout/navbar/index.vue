<script setup lang="ts">
import { CaretBottom } from '@element-plus/icons-vue'
import { nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { logout } from '@/api/system/login'
import { langTitle } from '@/hooks/use-common'
import { elMessage } from '@/hooks/use-element'
import { resetState } from '@/hooks/use-permission'
import { useBasicStore } from '@/store/basic'
import { useUserStore } from '@/store/user'
import Breadcrumb from './component/Breadcrumb.vue'
import Hamburger from './component/Hamburger.vue'
import LangSelect from './component/LangSelect.vue'
import ScreenFull from './component/ScreenFull.vue'
import ScreenLock from './component/ScreenLock.vue'
import SizeSelect from './component/SizeSelect.vue'
import ThemeSelect from './component/ThemeSelect.vue'

const basicStore = useBasicStore()
const { settings, sidebar, setToggleSideBar } = basicStore
const router = useRouter()
const userStore = useUserStore()

// 切换sidebar按钮
function toggleSideBar() {
  setToggleSideBar()
}

// 退出登录
async function loginOut() {
  try {
    await logout()
    nextTick(() => {
      resetState()
    })
    elMessage('退出登录成功')
    router.push('/login?redirect=/')
  }
  catch (error) {
    console.log(error)
  }
}
</script>

<template>
  <div class="navbar">
    <div class="heard-left">
      <!-- 切换sidebar按钮  -->
      <Hamburger
        v-if="settings.showHamburger"
        :is-active="sidebar.opened"
        class="hamburger-container"
        @toggle-click="toggleSideBar"
      />
      <!-- 面包屑导航  -->
      <Breadcrumb
        v-if="basicStore.device === 'desktop'"
        class="breadcrumb-container"
      />
    </div>

    <!-- 导航标题 -->
    <div v-if="settings.showNavbarTitle" class="heard-center-title">
      {{ settings.title }}
    </div>

    <!-- 下拉操作菜单 -->
    <div v-if="settings.ShowDropDown" class="heard-right">
      <div v-if="basicStore.device === 'desktop'" class="heard-righ-btn">
        <ScreenFull />
        <ScreenLock />
        <ThemeSelect />
        <SizeSelect />
        <LangSelect />
      </div>

      <el-dropdown trigger="click" size="medium">
        <div class="avatar-wrapper">
          <el-avatar shape="square" :size="40" :src="userStore.userAvatar" />
          <CaretBottom style="width: 1em; height: 1em; margin-left: 4px" />
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <router-link to="/">
              <el-dropdown-item>首页</el-dropdown-item>
            </router-link>

            <router-link to="/user/profile">
              <el-dropdown-item>个人中心</el-dropdown-item>
            </router-link>

            <a target="_blank" href="https://github.com/silent-rain/gin-admin">
              <el-dropdown-item>{{ langTitle('Github') }}</el-dropdown-item>
            </a>

            <el-dropdown-item divided @click="loginOut">
              登出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<style lang="scss" scoped>
// navbar
.navbar {
  height: var(--nav-bar-height);

  display: flex;
  align-items: center;
  justify-content: space-between;

  .heard-left {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .heard-right {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .heard-right .heard-right-btn {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }
}

// logo
.avatar-wrapper {
  margin: 5px 0;
  cursor: pointer;
}

// 导航标题
.heard-center-title {
  text-align: center;
  font-weight: 600;
  font-size: 20px;
}
</style>
