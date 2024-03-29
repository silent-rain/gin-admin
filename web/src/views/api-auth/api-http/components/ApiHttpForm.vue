<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加HTTP接口' : '编辑HTTP接口'"
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
      <el-form-item label="上级接口" prop="parent_id">
        <el-tree-select
          v-model="props.data.parent_id"
          :data="apiOptions"
          node-key="id"
          :props="{
            children: 'children',
            label: 'name',
          }"
          :render-after-expand="false"
          filterable
          accordion
          :check-strictly="true"
          placeholder="请选择上级接口"
        >
          <template #default="{ node, _data }">
            <span class="custom-tree-node">
              <span>{{ node.label }}</span>
            </span>
          </template>
        </el-tree-select>
      </el-form-item>
      <el-form-item label="接口名称" prop="name">
        <el-input v-model="props.data.name" placeholder="请输入接口名称" />
      </el-form-item>
      <el-form-item label="URI资源" prop="uri">
        <el-input v-model="props.data.uri" placeholder="请输入URI资源地址" />
      </el-form-item>
      <el-form-item label="请求类型" prop="method">
        <el-select
          v-model="props.data.method"
          clearable
          placeholder="请选择请求类型"
        >
          <el-option
            v-for="item in methodOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="启用状态" prop="status">
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
import {
  updateApiHttp,
  addApiHttp,
  getAllApiHttpTree,
} from '@/api/api-auth/api-http';
import { ApiHttp, ApiHttpTreeRsp } from '~/api/api-auth/api-http';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: ApiHttp;
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
  name: [{ required: true, message: '请输入接口名称', trigger: 'change' }],
  uri: [{ required: true, message: '请输入URI资源地址', trigger: 'blur' }],
  method: [{ required: true, message: '请选择请求类型', trigger: 'blur' }],
  status: [{ required: true, message: '请选择启用状态', trigger: 'change' }],
});

const methodOptions = ['GET', 'POST', 'PUT', 'DELETE'];
// Http协议接口列表
const apiOptions = ref<ApiHttp[]>([]);

onBeforeMount(() => {
  fetchAllApiHttpTree();
});

// 获取所有Http协议接口信息树
const fetchAllApiHttpTree = async () => {
  try {
    const resp = (await getAllApiHttpTree()).data as ApiHttpTreeRsp;
    apiOptions.value = resp.data_list.filter((v: ApiHttp) => {
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
        await addApiHttp(props.data);
      } else {
        await updateApiHttp(props.data);
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
