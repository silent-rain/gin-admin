<template>
  <el-form
    ref="ruleFormRef"
    :rules="userRules"
    :model="state.formData"
    label-width="100px"
    style="width: 100%; padding-right: 20px"
  >
    <el-row>
      <el-col :span="12">
        <el-form-item label="旧密码" prop="old_password">
          <el-input
            v-model="state.formData.old_password"
            placeholder="请输旧密码"
          />
        </el-form-item>
        <el-form-item label="新密码" prop="new_password">
          <el-input
            v-model="state.formData.new_password"
            type="password"
            show-password
            placeholder="请输入新密码"
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="new_password2">
          <el-input
            v-model="state.formData.new_password2"
            type="password"
            show-password
            placeholder="请再次输入新密码"
          />
        </el-form-item>

        <div class="submit">
          <el-button type="primary" @click="handleSubmit(ruleFormRef)">
            提交
          </el-button>
        </div>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup lang="ts">
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { updateUserPwd } from '@/api/permission/user';
import { useUserStore } from '@/store/user';

const userStore = useUserStore();

const state = reactive({
  formData: {
    old_password: '',
    new_password: '',
    new_password2: '',
  },
});
const ruleFormRef = ref<FormInstance>();
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入新密码'));
  } else if (value !== state.formData.new_password) {
    callback(new Error('两次输入的密码不一致'));
  } else {
    callback();
  }
};
const userRules = reactive<FormRules>({
  old_password: [{ required: true, message: '请输旧密码', trigger: 'blur' }],
  new_password: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
  new_password2: [
    { required: true, validator: validatePass2, trigger: 'blur' },
  ],
});

// 提交
const handleSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }
    const data = {
      id: userStore.userId,
      ...state.formData,
    };
    try {
      await updateUserPwd(data);
      state.formData = {} as any;
      ElMessage.success('操作成功');
    } catch (error) {
      console.log(error);
    }
  });
};
</script>

<style scoped lang="scss">
.submit {
  text-align: right;
}
</style>
