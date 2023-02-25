<template>
  <div class="title">
    <el-button
      v-if="!state.userFormEdit"
      class="button"
      text
      type="primary"
      @click="handleEdit"
    >
      编辑
    </el-button>
    <template v-else>
      <el-button
        class="button"
        text
        type="primary"
        @click="handleUserCancelEdit"
      >
        取消
      </el-button>
      <el-button
        class="button"
        text
        type="primary"
        @click="handleUserSubmit(userRuleFormRef)"
      >
        保存
      </el-button>
    </template>
    <el-divider />
  </div>
  <div class="content">
    <el-form
      ref="userRuleFormRef"
      :rules="userRules"
      :model="props.data"
      label-width="100px"
      style="width: 100%; padding-right: 20px"
    >
      <el-row>
        <el-col :span="12">
          <el-form-item label="昵称" prop="nickname">
            <el-input
              v-if="state.userFormEdit"
              v-model="props.data.nickname"
              placeholder="请输入昵称"
            />
            <span v-else>{{ props.data.nickname }}</span>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="姓名" prop="realname">
            <el-input
              v-if="state.userFormEdit"
              v-model="props.data.realname"
              placeholder="请输入姓名"
            />
            <span v-else>{{ props.data.realname }}</span>
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="年龄" prop="age">
            <el-input-number
              v-if="state.userFormEdit"
              v-model="props.data.age"
              :min="1"
              style="width: 100%"
            />
            <span v-else>{{ props.data.age }}</span>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="出生日期" prop="birthday">
            <el-date-picker
              v-if="state.userFormEdit"
              v-model="props.data.birthday"
              type="date"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              placeholder="请选择出生日期"
            />
            <span v-else>{{ props.data.birthday }}</span>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="性别" prop="gender">
            <el-radio-group
              v-if="state.userFormEdit"
              v-model="props.data.gender"
            >
              <el-radio :label="1">女</el-radio>
              <el-radio :label="2">男性</el-radio>
              <el-radio :label="0">保密</el-radio>
            </el-radio-group>
            <template v-else>
              <el-tag v-if="props.data.gender === 0" type="info">保密</el-tag>
              <el-tag v-else-if="props.data.gender === 1" type="success">
                女
              </el-tag>
              <el-tag v-else>男</el-tag>
            </template>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="头像" prop="avatar">
            <UploadAvatar
              v-if="state.userFormEdit"
              class="avatar-uploader"
              v-model:url="props.data.avatar"
            ></UploadAvatar>
            <el-avatar v-else :size="60" :src="remoteImageUrl" />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="介绍" prop="intro">
            <el-input
              v-if="state.userFormEdit"
              v-model="props.data.intro"
              type="textarea"
              placeholder="介绍内容"
            />
            <p class="user-intro" v-else>{{ props.data.intro }}</p>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { User } from '@/typings/api/permission/user';
import UploadAvatar from '@/components/Upload/UploadAvatar.vue';

const props = withDefaults(
  defineProps<{
    data: User;
  }>(),
  {},
);

const emits = defineEmits(['update:data', 'refresh']);

const state = reactive({
  userFormEdit: false,
  dataRaw: {} as User,
});
const userRuleFormRef = ref<FormInstance>();
const userRules = reactive<FormRules>({
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
  gender: [{ required: true, message: '请选择性别', trigger: 'blur' }],
  birthday: [{ required: true, message: '请选择出生日期', trigger: 'blur' }],
});

// 远程图片地址
const remoteImageUrl = computed(() => {
  if (!props.data.avatar) {
    return '';
  }
  return import.meta.env.VITE_APP_IMAGE_URL + props.data.avatar;
});

// 编辑
const handleEdit = () => {
  state.userFormEdit = true;
  state.dataRaw = { ...props.data };
};

// 取消用户编辑
const handleUserCancelEdit = () => {
  emits('update:data', { ...state.dataRaw });
  state.userFormEdit = false;
};

// 提交用户编辑
const handleUserSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }
    try {
      await updateUser(props.data);
      emits('refresh');
      state.userFormEdit = false;
      ElMessage.success('操作成功');
    } catch (error) {
      console.log(error);
    }
  });
};
</script>

<style scoped lang="scss">
.user-config {
  .user-config-user {
    .title {
      text-align: right;
    }
    .content {
      margin-top: 25px;
    }
  }
}
.user-intro {
  padding: 0;
  margin: 0;
}
</style>
