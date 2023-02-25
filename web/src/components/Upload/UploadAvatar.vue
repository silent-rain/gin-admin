<template>
  <el-upload
    class="avatar-uploader"
    :action="uploadAvatar"
    :headers="headerObj"
    :show-file-list="false"
    :on-success="handleAvatarSuccess"
    :before-upload="beforeAvatarUpload"
  >
    <img v-if="remoteImageUrl" :src="remoteImageUrl" class="upload-avatar" />
    <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
  </el-upload>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue';
import { uploadAvatar } from '@/api/system/upload';
import { useUserStore } from '@/store/user';
import { ElMessage, UploadProps } from 'element-plus';

const props = withDefaults(
  defineProps<{
    url: string;
    remoteUrl?: string;
  }>(),
  {},
);

const emits = defineEmits(['update:url', 'update:remoteUrl']);

const userStore = useUserStore();

const headerObj = {
  authorization: userStore.token,
};

// 上传头像成功事件
const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile,
) => {
  // state.imageBlob = URL.createObjectURL(uploadFile.raw!);
  emits('update:url', response.data.url);
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
  if (!props.url) {
    return '';
  }
  const url = import.meta.env.VITE_APP_IMAGE_URL + props.url;

  emits('update:remoteUrl', url);
  return url;
});
</script>

<style scoped lang="scss">
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
