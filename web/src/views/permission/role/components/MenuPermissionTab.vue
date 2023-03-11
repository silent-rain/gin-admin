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
      label: 'title',
    }"
  />
</template>

<script setup lang="ts">
import { ElMessage, ElTree } from 'element-plus';
import { getAllMenuTree } from '@/api/permission/menu';
import {
  getRoleMenuRelList,
  updateRoleMenuRel,
} from '@/api/permission/role-menu-rel';
import { MenuListRsp, Menu } from '~/api/permission/menu';
import { RoleMenuRelListRsp } from '~/api/permission/role-menu-rel';

const props = withDefaults(
  defineProps<{
    roleId: number;
  }>(),
  {},
);

const treeRef = ref<InstanceType<typeof ElTree>>();
const treeData = ref<Menu[]>([]);

onBeforeMount(() => {
  fetchAllMenuTree();
});

// 获取菜单树
const fetchAllMenuTree = async () => {
  try {
    const resp = (await getAllMenuTree()).data as MenuListRsp;
    treeData.value = resp.data_list;
  } catch (error) {
    console.log(error);
  }
};

// 获取角色关联的菜单列表
const fetchRoleMenuList = async (roleId: number) => {
  try {
    const resp = (
      await getRoleMenuRelList({
        role_id: roleId,
      })
    ).data as RoleMenuRelListRsp;
    const menuIds = resp.data_list.map((v) => v.menu_id);
    // 设置已关联的列表
    treeRef.value!.setCheckedKeys(menuIds, false);
  } catch (error) {
    console.log(error);
  }
};

// 提交
const handleSubmit = async () => {
  try {
    const data = {
      role_id: props.roleId,
      menu_ids: treeRef.value!.getCheckedKeys(false),
    };
    await updateRoleMenuRel(data);
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
    fetchRoleMenuList(props.roleId);
  },
  { immediate: true },
);
</script>

<style scoped lang="scss">
.permission {
  display: flex;
  justify-content: flex-end;
}
</style>
