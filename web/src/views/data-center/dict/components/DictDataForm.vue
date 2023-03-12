<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加字典项' : '编辑字典项'"
    width="500px"
    :before-close="handleClose"
  >
    <el-form
      ref="ruleFormRef"
      :rules="rules"
      :model="props.data"
      label-width="100px"
      style="width: 100%"
    >
      <el-form-item label="字典项名称" prop="name">
        <el-input v-model="props.data.name" placeholder="请输入字典名称" />
      </el-form-item>
      <el-form-item label="字典项值" prop="value">
        <el-input v-model="props.data.value" placeholder="请输入字典项值" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-switch
          v-model="props.data.status"
          :active-value="1"
          :inactive-value="0"
        />
      </el-form-item>
      <el-form-item label="备注" prop="note">
        <el-input
          v-model="props.data.note"
          type="textarea"
          placeholder="请输入备注"
        />
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
import { updateDictData, addDictData } from '@/api/data-center/dict-data';
import { DictData } from '~/api/data-center/dict-data';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: DictData;
    visible: boolean;
    type: string; // add/edit
  }>(),
  {},
);

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  name: [
    { required: true, message: '请输入字典项名称', trigger: 'blur' },
    { min: 2, message: '至少输入2个字符', trigger: 'blur' },
  ],
  value: [
    { required: true, message: '请输入字典项值', trigger: 'blur' },
    { min: 6, message: '至少输入6个字符', trigger: 'blur' },
  ],
  status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
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
        await addDictData(props.data);
      } else {
        await updateDictData(props.data);
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
