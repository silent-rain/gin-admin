<template>
  <div class="register-container columnCE">
    <div class="register-hero">
      <img src="@/assets/layout/login.svg" :alt="settings.title" />
    </div>

    <el-form
      ref="ruleFormRef"
      class="register-form"
      :model="userForm"
      label-position="right"
      label-width="80px"
      :rules="rules"
    >
      <div class="title-container">
        <h3 class="title text-center">用户注册</h3>
      </div>

      <el-form-item label="姓名" prop="realname">
        <el-input v-model="userForm.realname" placeholder="请输入姓名" />
      </el-form-item>
      <el-form-item label="昵称" prop="nickname">
        <el-input v-model="userForm.nickname" placeholder="请输入昵称" />
      </el-form-item>
      <el-form-item class="form-phone" label="手机号码" prop="phone">
        <el-input v-model="userForm.phone" placeholder="请输入手机号码" />
      </el-form-item>
      <el-form-item class="form-email" label="邮箱" prop="email">
        <el-input v-model="userForm.email" placeholder="请输入邮箱" />
      </el-form-item>
      <el-form-item label="性别" prop="gender">
        <el-radio-group v-model="userForm.gender">
          <el-radio :label="1">女</el-radio>
          <el-radio :label="2">男性</el-radio>
          <el-radio :label="0">保密</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="年龄" prop="age">
        <el-input-number v-model="userForm.age" :min="1" />
      </el-form-item>
      <el-form-item label="出生日期" prop="birthday">
        <el-date-picker
          v-model="userForm.birthday"
          type="date"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          placeholder="请选择出生日期"
        />
      </el-form-item>

      <el-form-item label="密码" prop="password">
        <el-input v-model="userForm.password" placeholder="请输入密码" />
      </el-form-item>
      <el-form-item label="确认密码" prop="password2">
        <el-input v-model="userForm.password2" placeholder="请再次输入密码" />
      </el-form-item>

      <!-- 验证码 -->
      <el-form-item label="">
        <img class="captcha" :src="state.captcha" @click="fetchCaptcha" />
      </el-form-item>
      <el-form-item label="验证码" prop="captcha">
        <div class="form-captcha">
          <el-input v-model="userForm.captcha" placeholder="请输入验证码" />
          <el-tooltip content="刷新验证码" placement="top">
            <el-button :icon="RefreshRight" @click="fetchCaptcha" />
          </el-tooltip>
        </div>
      </el-form-item>

      <el-button
        class="form-submit"
        type="primary"
        @click="submitForm(ruleFormRef)"
      >
        提交
      </el-button>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { RefreshRight } from '@element-plus/icons-vue';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';
import { register, getCaptcha } from '@/api/system/login';
import { User } from '~/api/system/user';
import { GetCaptchaRsp } from '~/api/system/login';

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();
const { settings } = useBasicStore();

const state: any = reactive({
  otherQuery: {},
  redirect: undefined,
  captcha: '',
});
const ruleFormRef = ref<FormInstance>();
const userForm = reactive({
  age: 0,
} as User);

// 检查密码
const validatePass2 = (_rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入密码'));
  } else if (value !== userForm.password) {
    callback(new Error('两次密码不一致!'));
  } else {
    callback();
  }
};
const rules = reactive<FormRules>({
  realname: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
  gender: [{ required: true, message: '请选择性别', trigger: 'blur' }],
  birthday: [{ required: true, message: '请选择出生日期', trigger: 'blur' }],
  phone: [
    { required: true, message: '请输入手机号码', trigger: 'blur' },
    { min: 11, max: 11, message: '手机号码不合法', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '至少输入6个字符', trigger: 'blur' },
  ],
  password2: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      validator: validatePass2,
      trigger: 'blur',
    },
  ],
});

onBeforeMount(() => {
  fetchCaptcha();
});

// 获取验证码
const fetchCaptcha = async () => {
  try {
    const resp = (await getCaptcha()).data as GetCaptchaRsp;
    state.captcha = resp.b64s;
    userForm.captcha_id = resp.captcha_id;
  } catch (error) {
    console.log(error);
  }
};

// 提交
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }
    const data = { ...userForm };
    try {
      await register(data);
      ElMessage.success('注册成功');
      router.push('/login');
    } catch (error) {
      console.log(error);
      fetchCaptcha();
    }
  });
};
</script>

<style lang="scss" scoped>
$bg: #fbfcff;
$dark_gray: #333;
$gray: #999;
$light_gray: #eee;
.register-container {
  height: 100vh;
  position: relative;
  overflow-y: hidden;
  width: 100%;
  background-color: $bg;
  .register-form {
    width: 400px;
    padding: 40px 30px;
    background: #fff;
    box-shadow: 0px 4px 16px rgba(4, 61, 175, 0.15);
    border-radius: 8px;
    margin-right: 14vw;
    z-index: 10;
    @media screen and (min-width: 769px) and (max-width: 992px) {
      margin-right: 10vw;
    }
    @media only screen and (max-width: 768px) {
      margin-right: auto;
      margin-left: auto;
    }
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

.captcha {
  cursor: pointer;
}
.form-captcha {
  display: flex;
  align-items: center;

  img {
    width: 115px;
    height: auto;
  }
}
.form-submit {
  float: right;
}
.register-hero {
  width: 40vw;
  position: absolute;
  top: 50%;
  left: 15vw;
  z-index: 0;
  transform: translateY(-50%);
  @media screen and (min-width: 769px) and (max-width: 992px) {
    width: 60vw;
    left: 5vw;
  }
  @media screen and (max-width: 768px) {
    width: 100vw;
    left: 0;
  }
  img {
    width: 100%;
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
</style>
