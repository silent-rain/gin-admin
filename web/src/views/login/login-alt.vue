<template>
  <div class="login-container columnCC">
    <div class="login-bg">
      <img src="@/assets/layout/login-bg.svg" :alt="settings.title" />
    </div>

    <div class="login-pane">
      <img
        src="@/assets/layout/login-top.svg"
        class="login-top"
        :alt="settings.title"
      />
      <img
        src="@/assets/layout/login-front.svg"
        class="login-front"
        :alt="settings.title"
      />

      <el-form
        ref="refLoginForm"
        class="login-form"
        :model="subForm"
        :rules="formRules"
      >
        <div class="title-container">
          <h3 class="title text-center">{{ settings.title }}</h3>
        </div>

        <el-form-item prop="username">
          <el-input
            v-model="subForm.username"
            placeholder="请输入手机号码/邮箱"
            :prefix-icon="User"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            ref="refPassword"
            v-model="subForm.password"
            :prefix-icon="Lock"
            show-password
            placeholder="请输入密码"
          />
        </el-form-item>

        <!-- 验证码 -->
        <el-form-item prop="captcha">
          <div class="form-captcha">
            <el-input
              v-model="subForm.captcha"
              placeholder="请输入验证码"
              @keyup.enter="handleLogin(refLoginForm)"
            />
            <img
              class="captcha"
              :src="state.captchaSrc"
              @click="fetchCaptcha"
            />
          </div>
        </el-form-item>

        <el-button
          :loading="subLoading"
          type="warning"
          class="login-btn"
          size="default"
          round
          @click.prevent="handleLogin(refLoginForm)"
        >
          登录
        </el-button>
        <div class="register-btn">
          <el-button type="warning" link @click="handleRegister">
            没有用户?点击注册
          </el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount, reactive, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { User, Lock } from '@element-plus/icons-vue';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';
import { login } from '@/api/system/login';
import { getCaptcha } from '@/api/system/login';
import { GetCaptchaRsp } from '~/api/system/login';
import { md5Encode } from '@/utils/md5';

/* listen router change and set the query  */
const { settings } = useBasicStore();
const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const subForm = reactive({
  username: '18312465088',
  password: '888888',
  captcha_id: '',
  captcha: '',
});
const state: any = reactive({
  otherQuery: {},
  redirect: undefined,
  captchaSrc: '',
});
const formRules = reactive<FormRules>({
  username: [
    { required: true, message: '请输入手机号码/邮箱', trigger: 'blur' },
  ],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
});

onBeforeMount(() => {
  fetchCaptcha();
});

// 获取验证码
const fetchCaptcha = async () => {
  try {
    const resp = (await getCaptcha()).data as GetCaptchaRsp;
    state.captchaSrc = resp.b64s;
    subForm.captcha_id = resp.captcha_id;
  } catch (error) {
    console.log(error);
  }
};

const getOtherQuery = (query: any) => {
  return Object.keys(query).reduce((acc, cur) => {
    if (cur !== 'redirect') {
      acc[cur] = query[cur];
    }
    return acc;
  }, {});
};

watch(
  () => route.query,
  (query) => {
    if (query) {
      state.redirect = query.redirect;
      state.otherQuery = getOtherQuery(query);
    }
  },
  { immediate: true },
);

// 登录
let subLoading = ref(false);
const refLoginForm = ref<FormInstance>();
const handleLogin = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      subLoading.value = false;
      return;
    }
    subLoading.value = true;
    const data = {
      username: subForm.username,
      password: md5Encode(subForm.password),
      captcha_id: subForm.captcha_id,
      captcha: subForm.captcha,
    };
    try {
      const resp = (await login(data)).data;
      ElMessage.success('登录成功');
      userStore.setToken(resp?.token);
      router.push('/');
      subLoading.value = false;
    } catch (error) {
      console.log(error);
      subLoading.value = false;
      fetchCaptcha();
    }
  });
};

// 注册用户
const handleRegister = () => {
  router.push('/register');
};
</script>

<style lang="scss" scoped>
$bg: #ffe4b5;
$dark_gray: #333;
$gray: #999;
$light_gray: #eee;
.login-container {
  height: 100vh;
  position: relative;
  overflow-y: hidden;
  width: 100%;
  background-color: $bg;

  :deep(.el-form-item) {
    // border: 1px solid #e0e0e0;
    background: #fff;
    border-radius: 50px;
    color: #999;
    .el-form-item__content {
      position: relative;
    }
    .el-form-item__error {
      padding-left: 40px;
    }
  }
  .form-captcha {
    display: flex;
    align-items: center;

    img {
      width: 115px;
      height: auto;
      cursor: pointer;
    }
  }
  .login-pane {
    position: relative;
    .login-top,
    .login-front {
      position: absolute;
      top: 0;
      left: 50%;
    }
    .login-top {
      z-index: 0;
      transform: translateY(-85%) translateX(-50%);
    }
    .login-front {
      z-index: 11;
      transform: translateY(-35%) translateX(-50%);
    }
  }
  .login-form {
    width: 340px;
    padding: 40px 30px;
    background: #fff;
    box-shadow: 0px 4px 16px rgba(158, 105, 25, 0.15);
    border-radius: 8px;
    position: relative;
    z-index: 10;
  }
  .title-container {
    .title {
      font-size: 18px;
      color: $dark_gray;
      margin: 0px auto 25px auto;
      text-align: center;
      font-weight: bold;
    }
  }
}
.login-bg {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 0;
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.svg-container {
  padding-left: 16px;
  color: $gray;
  text-align: center;
  width: 30px;
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
}

//登录按钮
.login-btn {
  width: 100%;
  margin-bottom: 10px;
  --el-button-bg-color: #fbcf47;
  --el-button-border-color: #fbcf47;
  --el-button-text-color: #8f5c0e;
  --el-button-hover-text-color: #8f5c0e;
}

// 输入框
:deep(.el-input__wrapper) {
  height: 40px;
  border-radius: 20px;
}
:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px $bg inset;
}
</style>
