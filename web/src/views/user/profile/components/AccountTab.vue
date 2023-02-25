<template>
  <el-card>
    <div class="account">
      <div class="header">
        <h3>密保手机</h3>
        <el-button
          text
          type="primary"
          @click="state.isPhoneEdit = !state.isPhoneEdit"
        >
          编辑
        </el-button>
      </div>
      <p class="info">
        <label class="info-title-item">已绑定手机:</label>
        <span class="info-text-item">
          {{ props.data.phone !== '' ? props.data.phone : '暂无' }}
        </span>
      </p>
    </div>
    <el-divider />
    <div class="account">
      <div class="header">
        <h3>密保邮箱</h3>
        <el-button
          text
          type="primary"
          @click="state.isEmailEdit = !state.isEmailEdit"
        >
          编辑
        </el-button>
      </div>
      <p class="info">
        <label class="info-title-item">已绑定邮箱:</label>
        <span class="info-text-item">
          {{ props.data.email !== '' ? props.data.email : '暂无' }}
        </span>
      </p>
    </div>
  </el-card>

  <!-- 修改密保手机 -->
  <el-dialog v-model="state.isPhoneEdit" title="修改密保手机" width="30%">
    <label>手机号码:</label>
    <el-input v-model="state.newPhone" placeholder="请输入手机号码" />
    <label>密码验证:</label>
    <el-input
      v-model="state.password"
      type="password"
      show-password
      placeholder="请输入密码"
    />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="state.isPhoneEdit = false">取消</el-button>
        <el-button class="button" type="primary" @click="handleUpdatePhone">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>

  <!-- 修改密保邮箱 -->
  <el-dialog v-model="state.isEmailEdit" title="修改密保邮箱" width="30%">
    <label>邮箱:</label>
    <el-input v-model="state.newEmail" placeholder="请输入邮箱" />
    <label>密码验证:</label>
    <el-input
      v-model="state.password"
      type="password"
      show-password
      placeholder="请输入密码"
    />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="state.isEmailEdit = false">取消</el-button>
        <el-button class="button" type="primary" @click="handleUpdateEmail">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus';
import { updatePhone, updateEmail } from '@/api/permission/user';
import { User } from '@/typings/api/permission/user';
import { md5Encode } from '@/utils/md5';

const props = withDefaults(
  defineProps<{
    data: User;
  }>(),
  {},
);

const emits = defineEmits(['refresh']);

const state = reactive({
  isPhoneEdit: false,
  newPhone: '',
  isEmailEdit: false,
  newEmail: '',
  password: '',
});

// 更新手机号码
const handleUpdatePhone = async () => {
  if (state.newPhone.trim() === '') {
    ElMessage.warning('新手机号码不能为空');
    return;
  }
  if (state.newPhone.trim().length !== 11) {
    ElMessage.warning('非法手机号码');
    return;
  }
  if (state.password.trim() === '') {
    ElMessage.warning('密码不能为空');
    return;
  }

  const data = {
    id: props.data.id,
    phone: state.newPhone.trim(),
    password: md5Encode(state.password),
  };
  try {
    await updatePhone(data);
    state.isPhoneEdit = false;
    emits('refresh');
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 更新邮箱
const handleUpdateEmail = async () => {
  if (state.newEmail.trim() === '') {
    ElMessage.warning('新邮箱不能为空');
    return;
  }
  if (state.password.trim() === '') {
    ElMessage.warning('密码不能为空');
    return;
  }

  const data = {
    id: props.data.id,
    email: state.newEmail.trim(),
    password: md5Encode(state.password),
  };
  try {
    await updateEmail(data);
    state.isEmailEdit = false;
    emits('refresh');
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
</script>

<style scoped lang="scss">
.account {
  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .info {
    color: #97a8be;

    .info-text-item {
      margin-left: 5px;
    }
  }
}
</style>
