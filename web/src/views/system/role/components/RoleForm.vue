<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加角色' : '编辑角色'"
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
      <el-form-item label="角色名称" prop="name">
        <el-input v-model="props.data.name" placeholder="请输入角色名称" />
      </el-form-item>
      <el-form-item label="角色状态" prop="status">
        <el-switch
          v-model="props.data.status"
          :active-value="1"
          :inactive-value="0"
        />
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input-number v-model="props.data.sort" :min="1" />
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
import { updateRole, addRole } from '@/api/system/role';
import { Role } from '~/api/system/role';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: Role;
    visible: boolean;
    type: string; // add/edit
    width?: string;
  }>(),
  {
    width: '100%',
  },
);

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  name: [
    { required: true, message: '请输入角色名称', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
  sort: [{ required: true, message: '最小为1', trigger: 'blur' }],
});

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
        await addRole(props.data);
      } else {
        await updateRole(props.data);
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
