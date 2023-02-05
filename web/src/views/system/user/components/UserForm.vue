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
          <el-form-item label="手机号码" prop="phone">
            <el-input v-model="props.data.phone" placeholder="请输入手机号码" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="props.data.email" placeholder="请输入邮箱" />
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
        <el-col :span="24">
          <el-form-item label="分配角色" prop="role_ids">
            <el-transfer
              v-model="roleIds"
              filterable
              filter-placeholder="筛选角色"
              :data="roleList"
              :props="{ key: 'id', label: 'name' }"
              :titles="['可选角色', '已选角色']"
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
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { updateUser, addUser } from '@/api/system/user';
import { getAllRole } from '@/api/system/role';
import { User } from '~/api/system/user';
import { RoleListRsp, Role } from '~/api/system/role';

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

// const { data } = toRefs(props);
// const data = toRef(props, 'data');

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
    const data = { ...props.data };
    data.role_ids = roleIds.value;
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
</script>

<style scoped lang="scss"></style>
