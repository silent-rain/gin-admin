<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加API令牌' : '编辑API令牌'"
    :width="props.width"
    :before-close="handleClose"
  >
    <el-form
      ref="ruleFormRef"
      :rules="rules"
      :model="props.data"
      label-width="100px"
      style="width: 100%"
    >
      <el-form-item label="用户昵称" prop="user_id">
        <el-select
          v-model="userInfo"
          filterable
          clearable
          remote
          reserve-keyword
          placeholder="请选择用户(支持远程检索)"
          remote-show-suffix
          value-key="id"
          :remote-method="remoteMethod"
          :loading="loading"
          @change="handleChangeUser"
        >
          <el-option
            v-for="item in userOptions"
            :key="item.id"
            :label="item.nickname"
            :value="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="访问权限" prop="permission">
        <el-select
          v-model="permission"
          multiple
          clearable
          placeholder="请选择权限"
          @change="handleChangePermission"
        >
          <el-option
            v-for="item in permissionOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="API口令" prop="passphrase">
        <el-input v-model="props.data.passphrase" placeholder="请输入API口令" />
      </el-form-item>
      <el-form-item label="启用状态" prop="status">
        <el-switch
          v-model="props.data.status"
          :active-value="1"
          :inactive-value="0"
        />
      </el-form-item>
      <el-form-item label="备注" prop="note">
        <el-input v-model="props.data.note" placeholder="请输入备注" />
      </el-form-item>
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
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import {
  updateUserApiToken,
  addUserApiToken,
} from '@/api/permission/user-api-token';
import { UserApiToken } from '~/api/permission/user-api-token';
import { User, UserListRsp } from '~/api/permission/user';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: UserApiToken;
    visible: boolean;
    type: string; // add/edit
    width?: string;
  }>(),
  {
    width: '500px',
  },
);

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  user_id: [{ required: true, message: '请选择用户', trigger: 'change' }],
  passphrase: [{ required: true, message: '请输入口令', trigger: 'blur' }],
  status: [{ required: true, message: '请选择启用状态', trigger: 'change' }],
});

const permissionOptions = ['GET', 'POST', 'PUT', 'DELETE'];
const permission = ref<string[]>([]);
const userOptions = ref<User[]>([]);
const userInfo = ref<User>({} as User);
const loading = ref(false);

onBeforeMount(() => {
  fetchUserList();
});

// 获取用户列表
const fetchUserList = async () => {
  // 默认显示200个用户
  const data = {
    page: 1,
    page_size: 200,
  };
  try {
    const resp = (await getUserList(data)).data as UserListRsp;
    userOptions.value = resp.data_list;
    userInfo.value.id = props.data.user_id;
    userInfo.value.id = 21;
    permission.value = props.data.permission.split(';');
  } catch (error) {
    console.log(error);
  }
};

// 获取用户列表
const remoteMethod = async (query: string) => {
  if (!query) {
    return;
  }

  // 远程搜索
  try {
    loading.value = true;
    const data = {
      nickname: query,
      page: 1,
      page_size: 50,
    };
    const resp = (await getUserList(data)).data as UserListRsp;
    userOptions.value = resp.data_list;
    loading.value = false;
  } catch (error) {
    console.log(error);
    loading.value = false;
  }
};

// 切换用户
const handleChangeUser = (value: User) => {
  emit('update:data', { ...props.data, user_id: value.id });
};
// 切换访问权限
const handleChangePermission = (value: string[]) => {
  emit('update:data', { ...props.data, permission: value.join(';') });
};

// 关闭
const handleClose = () => {
  emit('update:visible', false);
  emit('update:data', {});
};

// 取消
const handleCancel = () => {
  emit('update:visible', false);
  emit('update:data', {});
};
// 提交
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }
    try {
      if (props.type === 'add') {
        await addUserApiToken(props.data);
      } else {
        await updateUserApiToken(props.data);
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
</script>

<style scoped lang="scss"></style>
