<template>
  <el-dialog
    :model-value="props.visible"
    title="分配权限"
    width="400px"
    :before-close="handleClose"
  >
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
      style="height: 400px"
    />

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button type="primary" @click="handleSubmit"> 提交 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElMessage, ElTree } from 'element-plus';
import { getAllMenuTree } from '@/api/system/menu';
import {
  getRoleMenuRelList,
  updateRoleMenuRel,
} from '@/api/system/role-menu-rel';
import { Role } from '~/api/system/role';
import { MenuListRsp, Menu } from '~/api/system/menu';
import { RoleMenuRelListRsp } from '~/api/system/role-menu-rel';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: Role;
    visible: boolean;
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
const handleSubmit = async () => {
  try {
    const data = {
      role_id: props.data.id,
      menu_ids: treeRef.value!.getCheckedKeys(false),
    };
    console.log(data);
    await updateRoleMenuRel(data);
    emit('update:visible', false);
    emit('update:data', {});
    emit('refresh');
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

watch(
  () => props.data.id,
  () => {
    if (!props.data.id) {
      return;
    }
    fetchRoleMenuList(props.data.id);
  },
  { immediate: true },
);
</script>

<style scoped lang="scss"></style>
