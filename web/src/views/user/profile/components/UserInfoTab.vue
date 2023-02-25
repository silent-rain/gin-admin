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
            <el-upload
              v-if="state.userFormEdit"
              class="avatar-uploader"
              :action="uploadAvatar"
              :headers="headerObj"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
            >
              <img
                v-if="remoteImageUrl"
                :src="remoteImageUrl"
                class="upload-avatar"
              />
              <el-icon v-else class="avatar-uploader-icon">
                <Plus />
              </el-icon>
            </el-upload>
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
            <span v-else>{{ props.data.intro }}</span>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue';
import { ElMessage, FormInstance, FormRules, UploadProps } from 'element-plus';
import { uploadAvatar } from '@/api/system/upload';
import { useUserStore } from '@/store/user';
import { User } from '@/typings/api/permission/user';

const props = withDefaults(
  defineProps<{
    data: User;
  }>(),
  {},
);

const userStore = useUserStore();

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

const headerObj = {
  authorization: userStore.token,
};

// 上传头像成功事件
const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile,
) => {
  // state.imageBlob = URL.createObjectURL(uploadFile.raw!);
  props.data.avatar = `${response.data.url}`;
};
// 上传头像事件
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  const imgfileType = [
    'image/gif',
    'image/jpg',
    'image/jpeg',
    'image/x-png',
    'image/png',
  ];
  if (imgfileType.indexOf(rawFile.type) === -1) {
    ElMessage.error('Avatar picture must be JPG/JPEG/PNG/GIF format!');
    return false;
  }
  if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!');
    return false;
  }
  return true;
};

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

// 头像
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color-darker);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}
.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}
.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 100px;
  text-align: center;
  border: 1px dashed var(--el-border-color-darker);
}
.upload-avatar {
  width: 100px;
  height: 100px;
}
</style>
