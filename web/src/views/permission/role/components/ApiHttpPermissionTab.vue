<template>
  <div class="permission">
    <el-button type="primary" @click="handleSubmit">提交</el-button>
  </div>
  <el-tree
    ref="treeRef"
    :data="treeData"
    show-checkbox
    default-expand-all
    node-key="id"
    highlight-current
    :props="{
      children: 'children',
      label: 'name',
    }"
  >
    <template #default="{ node, data }">
      <span class="custom-tree-node">
        <span>{{ node.label }}</span>
        <el-tag
          v-for="(item, _) in methodOptions.filter(
            (v) => v.value === data.method,
          )"
          class="api-http-method-item"
          size="small"
          :key="item.value"
          :type="item.type"
        >
          {{ data.method }}
        </el-tag>
        <span class="api-http-uri-item">{{ data.uri }}</span>
      </span>
    </template>
  </el-tree>
</template>

<script setup lang="ts">
import { ElMessage, ElTree } from 'element-plus';
import { getAllApiHttpTree } from '@/api/api-auth/api-http';
import {
  getApiRoleHttpRelList,
  updateApiRoleHttpRel,
} from '@/api/api-auth/api-role-http-rel';
import { ApiHttpTreeRsp, ApiHttp } from '~/api/api-auth/api-http';
import { ApiRoleHttpRelRsp } from '~/api/api-auth/api-role-http-rel';

const props = withDefaults(
  defineProps<{
    roleId: number;
  }>(),
  {},
);

// 请求类型列表
const methodOptions = [
  {
    value: 'GET',
    type: 'info',
  },
  {
    value: 'POST',
    type: '',
  },
  {
    value: 'PUT',
    type: 'warning',
  },
  {
    value: 'DELTE',
    type: 'danger',
  },
];

const treeRef = ref<InstanceType<typeof ElTree>>();
const treeData = ref<ApiHttp[]>([]);

onBeforeMount(() => {
  fetchAllApiHttpTree();
});

// 获取所有Http协议接口信息列表
const fetchAllApiHttpTree = async () => {
  try {
    const resp = (await getAllApiHttpTree()).data as ApiHttpTreeRsp;
    treeData.value = resp.data_list;
  } catch (error) {
    console.log(error);
  }
};

// 获取角色与Http协议接口关系列表
const fetchRApiRoleHttpReList = async (roleId: number) => {
  try {
    const resp = (
      await getApiRoleHttpRelList({
        role_id: roleId,
      })
    ).data as ApiRoleHttpRelRsp;
    const ids = resp.data_list.map((v) => v.api_id);
    // 设置已关联的列表
    treeRef.value!.setCheckedKeys(ids, false);
    console.log(ids);
  } catch (error) {
    console.log(error);
  }
};

// 提交
const handleSubmit = async () => {
  try {
    const data = {
      role_id: props.roleId,
      api_ids: treeRef.value!.getCheckedKeys(false),
    };
    await updateApiRoleHttpRel(data);
    fetchRApiRoleHttpReList(props.roleId);
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

watch(
  () => props.roleId,
  () => {
    if (!props.roleId) {
      return;
    }
    fetchRApiRoleHttpReList(props.roleId);
  },
  { immediate: true },
);
</script>

<style scoped lang="scss">
.permission {
  display: flex;
  justify-content: flex-end;
}

.custom-tree-node {
  .api-http-method-item {
    margin-left: 20px;
  }
  .api-http-uri-item {
    margin-left: 20px;
  }
}
</style>
