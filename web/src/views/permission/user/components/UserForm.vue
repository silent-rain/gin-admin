<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加用户' : '编辑用户'"
    :width="props.width"
    :before-close="handleClose"
  >
    <el-form
      ref="ruleFormRef"
      :rules="rules"
      :model="props.data"
      label-width="100px"
      style="width: 100%; padding-right: 20px"
    >
      <el-row>
        <el-col :span="12">
          <el-form-item label="昵称" prop="nickname">
            <el-input v-model="props.data.nickname" placeholder="请输入昵称" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="姓名" prop="realname">
            <el-input v-model="props.data.realname" placeholder="请输入姓名" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item class="form-phone" label="手机号码" prop="phone">
            <el-input
              v-if="props.type === 'add'"
              v-model="props.data.phone"
              placeholder="请输入手机号码"
            />
            <div v-else>
              <template v-if="!state.isPhoneEdit">
                <span>{{ props.data.phone }}</span>
                <el-button
                  link
                  type="primary"
                  size="small"
                  :icon="EditPen"
                  @click="state.isPhoneEdit = true"
                >
                  修改
                </el-button>
              </template>
              <template v-else>
                <el-input
                  v-model="state.newPhone"
                  placeholder="请输入新手机号码"
                  style="width: 180px"
                />
                <el-button
                  link
                  type="primary"
                  size="small"
                  style="margin-left: 3px"
                  @click="handleUpdatePhone"
                >
                  保存
                </el-button>
                <el-button link size="small" @click="state.isPhoneEdit = false">
                  取消
                </el-button>
              </template>
            </div>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item class="form-email" label="邮箱" prop="email">
            <el-input
              v-if="props.type === 'add'"
              v-model="props.data.email"
              placeholder="请输入邮箱"
            />
            <div v-else>
              <template v-if="!state.isEmailEdit">
                <span>{{ props.data.email }}</span>
                <el-button
                  link
                  type="primary"
                  size="small"
                  :icon="EditPen"
                  @click="state.isEmailEdit = true"
                >
                  修改
                </el-button>
              </template>
              <template v-else>
                <el-input
                  v-model="state.newEmail"
                  placeholder="请输入新邮箱"
                  style="width: 180px"
                />
                <el-button
                  link
                  type="primary"
                  size="small"
                  style="margin-left: 3px"
                  @click="handleUpdateEmail"
                >
                  保存
                </el-button>
                <el-button link size="small" @click="state.isEmailEdit = false">
                  取消
                </el-button>
              </template>
            </div>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="性别" prop="gender">
            <el-radio-group v-model="props.data.gender">
              <el-radio :label="1">女</el-radio>
              <el-radio :label="2">男性</el-radio>
              <el-radio :label="0">保密</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="年龄" prop="age">
            <el-input-number
              v-model="props.data.age"
              :min="1"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="出生日期" prop="birthday">
            <el-date-picker
              v-model="props.data.birthday"
              type="date"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              placeholder="请选择出生日期"
            />
          </el-form-item>
        </el-col>
        <el-col v-if="props.type === 'add'" :span="24">
          <el-form-item label="密码" prop="password">
            <el-input v-model="props.data.password" placeholder="请输入密码" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="头像" prop="avatar">
            <UploadAvatar
              class="avatar-uploader"
              v-model:url="props.data.avatar"
            ></UploadAvatar>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="分配角色" prop="role_ids">
            <el-select
              v-model="roleIds"
              multiple
              filterable
              clearable
              placeholder="Select"
              style="width: 100%"
            >
              <el-option
                v-for="item in roleList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否启用" prop="status">
            <el-switch
              v-model="props.data.status"
              :active-value="1"
              :inactive-value="0"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="排序" prop="sort">
            <el-input-number
              v-model="props.data.sort"
              :min="1"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="介绍" prop="intro">
            <el-input
              v-model="props.data.intro"
              type="textarea"
              placeholder="介绍内容"
            />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="备注" prop="note">
            <el-input
              v-model="props.data.note"
              type="textarea"
              placeholder="请输入备注"
            />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElMessage, FormInstance, FormRules, UploadProps } from 'element-plus';
import { EditPen, Plus } from '@element-plus/icons-vue';
import {
  updateUser,
  addUser,
  updatePhone,
  updateEmail,
} from '@/api/permission/user';
import { getAllRole } from '@/api/permission/role';
import { useUserStore } from '@/store/user';
import { User } from '@/typings/api/permission/user';
import { RoleListRsp, Role } from '~/api/permission/role';
import UploadAvatar from '@/components/Upload/UploadAvatar.vue';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: User;
    visible: boolean;
    type: string; // add/edit
    width?: string;
  }>(),
  {
    width: '100%',
  },
);

const userStore = useUserStore();

// const { data } = toRefs(props);
// const data = toRef(props, 'data');

const state = reactive({
  isPhoneEdit: false,
  newPhone: '',
  isEmailEdit: false,
  newEmail: '',
  imageUrl: '',
});
const headerObj = {
  authorization: userStore.token,
};

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
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
  sort: [{ required: true, message: '最小为1', trigger: 'blur' }],
});
const roleList = ref<Role[]>([]);
const roleIds = ref<number[]>([]);

onBeforeMount(() => {
  fetchAllRole();
  state.imageUrl = props.data.avatar;
});

// 获取所有角色
const fetchAllRole = async () => {
  try {
    const resp = (await getAllRole()).data as RoleListRsp;
    roleList.value = resp.data_list;
  } catch (error) {
    console.log(error);
  }
};

// 关闭
const handleClose = () => {
  emit('update:visible', false);
  emit('update:data', {});
  state.imageUrl = '';
};

// 取消
const handleCancel = () => {
  emit('update:visible', false);
  emit('update:data', {});
  state.imageUrl = '';
};
// 提交
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }
    const data = { ...props.data };
    data.role_ids = roleIds.value;
    data.avatar = state.imageUrl;
    try {
      if (props.type === 'add') {
        await addUser(data);
      } else {
        await updateUser(data);
      }
      emit('update:visible', false);
      emit('update:data', {});
      emit('refresh');
      ElMessage.success('操作成功');
    } catch (error) {
      console.log(error);
    }
  });
};

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

  const data = {
    id: props.data.id,
    phone: state.newPhone.trim(),
  };
  try {
    await updatePhone(data);
    state.isPhoneEdit = false;
    emit('refresh');
    emit('update:data', { ...props.data, phone: state.newPhone });
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

  const data = {
    id: props.data.id,
    email: state.newEmail.trim(),
  };
  try {
    await updateEmail(data);
    state.isEmailEdit = false;
    emit('refresh');
    emit('update:data', { ...props.data, email: state.newEmail });
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 上传头像成功事件
const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile,
) => {
  // state.imageBlob = URL.createObjectURL(uploadFile.raw!);
  state.imageUrl = `${response.data.url}`;
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
  if (!state.imageUrl) {
    return '';
  }
  return import.meta.env.VITE_APP_IMAGE_URL + state.imageUrl;
});

// 角色赋值
watch(
  () => props.data.roles,
  () => {
    if (!props.data.roles) {
      return;
    }
    roleIds.value = props.data.roles.map((v) => v.id);
  },
  { immediate: true },
);

// 头像赋值
watch(
  () => props.data.avatar,
  () => {
    if (!props.data.avatar) {
      return;
    }
    state.imageUrl = props.data.avatar;
  },
  { immediate: true },
);
</script>

<style scoped lang="scss">
.form-phone,
.form-email {
  .el-button + .el-button {
    margin-left: 0px;
  }
}
</style>
