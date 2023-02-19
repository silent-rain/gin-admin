<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加配置' : '编辑配置'"
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
      <el-form-item label="上级配置" prop="parent_id">
        <el-tree-select
          v-model="props.data.parent_id"
          :data="configOptions"
          node-key="id"
          :props="{
            children: 'children',
            label: 'name',
          }"
          :render-after-expand="false"
          filterable
          accordion
          :check-strictly="true"
          placeholder="请选择上级配置"
          style="width: 100%"
        >
          <template #default="{ node, _data }">
            <span class="custom-tree-node">
              <span>{{ node.label }}</span>
            </span>
          </template>
        </el-tree-select>
      </el-form-item>
      <el-form-item label="配置名称" prop="name">
        <el-input v-model="props.data.name" placeholder="请输入配置名称" />
      </el-form-item>
      <el-form-item label="配置KEY" prop="key">
        <el-input
          v-model="props.data.key"
          :disabled="props.type !== 'add'"
          placeholder="请输入配置KEY"
        />
      </el-form-item>
      <el-form-item label="配置值" prop="value">
        <el-input
          v-model="props.data.value"
          type="textarea"
          :rows="4"
          placeholder="请输入配置值"
        />
      </el-form-item>
      <el-form-item label="排序号" prop="sort">
        <el-input-number v-model="props.data.sort" :min="1" />
      </el-form-item>
      <el-form-item label="启用状态" prop="status">
        <el-radio-group v-model="props.data.status">
          <el-radio :label="1">启用</el-radio>
          <el-radio :label="0">停用</el-radio>
        </el-radio-group>
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
import { QuestionFilled } from '@element-plus/icons-vue';
import { updateConfig, addConfig, getAllConfigTree } from '@/api/system/config';
import { Config, ConfigListRsp } from '~/api/system/config';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: Config;
    visible: boolean;
    type: string; // add/edit
  }>(),
  {},
);

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  title: [
    { required: true, message: '请输入配置名称', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
});
// 配置列表
const configOptions = ref<Config[]>([]);

onBeforeMount(() => {
  fetchAllConfigTree();
});

// 获取所有配置树
const fetchAllConfigTree = async () => {
  try {
    const resp = (await getAllConfigTree()).data as ConfigListRsp;
    configOptions.value = resp.data_list.filter((v: any) => {
      // 过滤自身选择, 防止自依赖
      if (v.id !== props.data.id) {
        return true;
      }
    });
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

    try {
      if (props.type === 'add') {
        await addConfig(props.data);
      } else {
        await updateConfig(props.data);
      }
      fetchAllConfigTree();
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

<style scoped lang="scss">
.el-divider--horizontal {
  margin: 10px, 0;
}
</style>
