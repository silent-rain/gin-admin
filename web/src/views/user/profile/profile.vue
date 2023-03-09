<template>
  <el-row :gutter="10">
    <el-col :span="8">
      <el-card class="user-preview">
        <div class="user-avatar">
          <el-avatar :size="80" :src="remoteImageUrl" />
        </div>
        <div class="user-nickname">
          <h1>{{ state.user.nickname }}</h1>
        </div>
        <div class="user-intro">
          <p>{{ state.user.intro }}</p>
        </div>

        <div class="user-info">
          <p>
            <label class="icon">
              <el-icon><Iphone /></el-icon>
            </label>
            <span class="text">{{ state.user.phone }}</span>
          </p>
          <p>
            <label class="icon">
              <el-icon><Connection /></el-icon>
            </label>
            <el-link
              class="text"
              type="primary"
              :href="'mailto:' + state.user.email"
              target="_blank"
            >
              {{ state.user.email }}
            </el-link>
          </p>
          <p v-if="state.roles.length > 0" class="user-role">
            <label class="icon">
              <el-icon><UserFilled /></el-icon>
            </label>
            <el-tag class="text" v-for="(item, i) in state.roles">
              {{ item.name }}
            </el-tag>
          </p>
        </div>
      </el-card>
    </el-col>
    <el-col :span="16">
      <el-card class="user-config">
        <el-tabs v-model="activeName" tab-position="top">
          <el-tab-pane class="user-config-user" label="基本信息" name="user">
            <UserInfoTab
              v-model:data="state.user"
              @refresh="fetchUserInfo"
            ></UserInfoTab>
          </el-tab-pane>
          <el-tab-pane label="修改密码" name="password">
            <PasswordTab></PasswordTab>
          </el-tab-pane>
          <el-tab-pane label="账号绑定" name="account">
            <AccountTab
              :data="state.user"
              @refresh="fetchUserInfo"
            ></AccountTab>
          </el-tab-pane>
          <el-tab-pane label="Token令牌" name="token">
            <ApiTokenTab></ApiTokenTab>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import { UserFilled, Iphone, Connection } from '@element-plus/icons-vue';
import { getUserInfo } from '@/api/permission/user';
import { User } from '~/api/permission/user';
import { Role } from '~/api/permission/role';
import UserInfoTab from './components/UserInfoTab.vue';
import PasswordTab from './components/PasswordTab.vue';
import AccountTab from './components/AccountTab.vue';
import ApiTokenTab from './components/ApiTokenTab/index.vue';

const activeName = ref('user');
const state = reactive({
  user: {} as User,
  roles: [] as Role[],
});

onBeforeMount(() => {
  fetchUserInfo();
});

// 获取用户信息
const fetchUserInfo = async () => {
  try {
    const resp = await getUserInfo();
    state.user = resp.data.user as User;
    state.roles = resp.data.roles as Role[];
  } catch (error) {
    console.log(error);
  }
};

// 远程图片地址
const remoteImageUrl = computed(() => {
  if (!state.user.avatar) {
    return '';
  }
  return import.meta.env.VITE_APP_IMAGE_URL + state.user.avatar;
});
</script>

<style scoped lang="scss">
.user-preview {
  min-height: 500px;
}
.user-avatar {
  text-align: center;
}
.user-nickname {
  text-align: center;
  margin-top: 10px;
}
.user-intro {
  margin-top: 25px;
}
.user-role {
  .el-tag {
    margin-left: 5px;
  }
}
.user-info {
  margin-top: 25px;
}
.user-info p {
  display: flex;
  align-items: center;

  .text {
    margin-left: 5px;
  }
}
</style>
