<template>
  <el-upload
    v-model:file-list="fileList"
    :action="uploadImage"
    :headers="headerObj"
    list-type="picture-card"
    multiple
    :before-upload="beforeImageUpload"
    :on-preview="handlePictureCardPreview"
    :on-remove="handleRemove"
    :on-success="handleImageSuccess"
  >
    <el-icon><Plus /></el-icon>

    <template #tip>
      <div class="el-upload__tip text-red">
        {{ props.tip }}
      </div>
    </template>
  </el-upload>

  <el-dialog v-model="dialogVisible">
    <img w-full :src="dialogImageUrl" alt="Preview Image" />
  </el-dialog>
</template>

<script setup lang="ts">
import { Plus } from '@element-plus/icons-vue';
import { uploadImage } from '@/api/system/upload';
import { useUserStore } from '@/store/user';
import { ElMessage, UploadProps, UploadUserFile } from 'element-plus';

const props = withDefaults(
  defineProps<{
    fileList: UploadUserFile[];
    imgSize?: number; // mb
    fileType?: string[];
    tip?: string;
  }>(),
  {
    imgSize: 2,
    fileType: () => {
      return [
        'image/gif',
        'image/jpg',
        'image/jpeg',
        'image/x-png',
        'image/png',
      ];
    },
    tip: '',
  },
);

const emits = defineEmits(['update:fileList']);

const userStore = useUserStore();

const fileList = ref<UploadUserFile[]>();
const headerObj = {
  authorization: userStore.token,
};

const dialogImageUrl = ref('');
const dialogVisible = ref(false);

onBeforeMount(() => {});
onMounted(() => {
  fileList.value = [
    ...props.fileList.map((v) => {
      v.url = import.meta.env.VITE_APP_IMAGE_URL + v.url;
      return v;
    }),
  ];
});

// 上传之前文件校验事件
const beforeImageUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (props.fileType.indexOf(rawFile.type) === -1) {
    ElMessage.error(
      `file must be ${props.fileType.join(', ')} format! , err: ${
        rawFile.type
      }`,
    );
    return false;
  }
  if (rawFile.size / 1024 / 1024 > props.imgSize) {
    ElMessage.error(`filee size can not exceed ${props.imgSize}MB!`);
    return false;
  }
  return true;
};

// 移除照片
const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
  console.log(uploadFile, uploadFiles);
};

// 预览照片
const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url!;
  dialogVisible.value = true;
};

// 上传成功事件
const handleImageSuccess: UploadProps['onSuccess'] = (response, uploadFile) => {
  // state.imageBlob = URL.createObjectURL(uploadFile.raw!);
  emits(
    'update:fileList',
    fileList.value?.map((v) => {
      const resp = v.response as any;
      const url = resp?.data.url.replace(import.meta.env.VITE_APP_IMAGE_URL);
      return {
        name: v.name,
        url: url,
      };
    }),
  );
};
</script>

<style scoped lang="scss"></style>
