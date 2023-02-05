<template>
  <el-card>
    <!-- 过滤条件 -->
    <div class="filter">
      <label>角色名称: </label>
      <el-input
        v-model="listQuery.name"
        class="filter-name"
        placeholder="请输入角色名称"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button type="primary" :icon="Search" @click="handleFilter"
          >查询
        </el-button>
        <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
      </el-button-group>
    </div>

    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button">
        <ConvenienButtons
          :buttonList="['add', 'batchDelete']"
          @add-event="handleAdd"
          @batch-delete-event="handleBatchDelete"
        />
      </div>
      <div class="right-button">
        <ConvenienTools
          @refreshEvent="fetchRoleList"
          v-model:size="tableSize"
          :screenFullElement="'el-table-role'"
          :checkAllList="checkAllList"
          v-model:checkedDict="checkedDict"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <RoleForm
      v-model:data="state.roleForm.data"
      v-model:visible="state.roleForm.visible"
      :type="state.roleForm.type"
      :width="state.roleForm.width"
      @refresh="fetchRoleList"
    />

    <!-- 分配权限 -->
    <MenuPermission
      v-model:data="state.menuPermission.data"
      v-model:visible="state.menuPermission.visible"
    />

    <el-table
      class="el-table-role"
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column
        v-if="checkedDict.id"
        prop="id"
        label="角色ID"
        width="80"
      />
      <el-table-column
        v-if="checkedDict.name"
        prop="name"
        label="角色名称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.sort"
        prop="sort"
        label="排序"
        show-overflow-tooltip
        width="80"
      />
      <el-table-column
        v-if="checkedDict.status"
        prop="status"
        label="角色状态"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.note"
        prop="note"
        label="备注"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.created_at"
        prop="created_at"
        label="创建时间"
        width="165"
      />
      <el-table-column
        v-if="checkedDict.updated_at"
        prop="updated_at"
        label="更新时间"
        width="165"
      />
      <el-table-column
        v-if="checkedDict.operation"
        fixed="right"
        label="操作"
        align="center"
        width="186"
      >
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            :icon="EditPen"
            @click="handleEdit(scope.row)"
            >修改
          </el-button>
          <el-button
            link
            type="primary"
            size="small"
            :icon="Finished"
            @click="handleMenuPermission(scope.row)"
            >分配权限
          </el-button>
          <el-popconfirm
            confirm-button-text="确认"
            cancel-button-text="取消"
            :icon="InfoFilled"
            icon-color="#E6A23C"
            title="确定删除吗?"
            @confirm="handleDelete(scope.row)"
            @cancel="handleCancelEvent"
          >
            <template #reference>
              <el-button link type="danger" size="small" :icon="Delete"
                >删除
              </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <Pagination
      v-model:currentPage="listQuery.page"
      v-model:pageSize="listQuery.page_size"
      :total="tableDataTotal"
      @pagination="fetchRoleList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { useBasicStore } from '@/store/basic';
import {
  EditPen,
  Search,
  Delete,
  Finished,
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import {
  getRoleList,
  updateRoleStatus,
  deleteRole,
  batchDeleteRole,
} from '@/api/system/role';
import { RoleListRsp, Role } from '~/api/system/role';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ConvenienButtons from '@/components/ConvenienButtons/index.vue';
import RoleForm from './components/RoleForm.vue';
import MenuPermission from './components/MenuPermission.vue';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = reactive({
  name: '',
  page: 1,
  page_size: 10,
});
// 过滤事件
const handleFilter = () => {
  fetchRoleList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.name = '';
};

const state = reactive({
  roleForm: {
    data: {} as Role,
    visible: false,
    type: '',
    width: '500px',
  },
  menuPermission: {
    data: {} as Role,
    visible: false,
  },
});

const checkAllList = [
  { label: '角色ID', value: 'id', disabled: false, enabled: true },
  { label: '角色名称', value: 'name', disabled: true, enabled: true },
  { label: '排序', value: 'sort', disabled: false, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: true },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<Role[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<Role[]>([]);

onBeforeMount(() => {
  fetchRoleList();
});

// 获取角色列表
const fetchRoleList = async () => {
  try {
    const resp = (await getRoleList(listQuery)).data as RoleListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: Role) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteRole(data);
    fetchRoleList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: Role) => {
  state.roleForm.data = { ...row };
  state.roleForm.type = 'edit';
  state.roleForm.visible = true;
};
// 添加
const handleAdd = async () => {
  state.roleForm.data.sort = 1;
  state.roleForm.type = 'add';
  state.roleForm.visible = true;
};
// 多选事件
const handleSelectionChange = (val: Role[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: Role) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteRole(data);
    fetchRoleList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 删除取消事件
const handleCancelEvent = () => {
  ElMessage.warning('取消操作');
};

// 状态变更
const handleStatusChange = async (row: Role) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    const resp = (await updateRoleStatus(data)).data as RoleListRsp;
    tableData.value = resp.data_list;
    fetchRoleList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 分配权限
const handleMenuPermission = async (row: Role) => {
  state.menuPermission.data = { ...row };
  state.menuPermission.visible = true;
};
</script>

<style scoped lang="scss">
.filter {
  .filter-name {
    width: 200px;
    margin-left: 8px;
  }
}

.operation-button {
  margin-top: 20px;

  display: flex;
  justify-content: space-between;

  .left-button {
    .el-button + .el-button {
      margin-left: 8px;
    }
  }

  .el-button + .el-button {
    margin-left: 0px;
  }
}

.el-button + .el-button {
  margin-left: 0px;
}
</style>
