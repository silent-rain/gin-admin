<template>
  <el-card id="role-manage">
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
        <el-button type="primary" :icon="Plus" @click="handleAdd"
          >添加
        </el-button>
        <el-button type="danger" :icon="Delete" @click="handleBatchDelete"
          >批量删除
        </el-button>
      </div>
      <div class="right-button">
        <el-tooltip content="刷新" placement="top">
          <el-button :icon="RefreshRight" @click="fetchRoleList" />
        </el-tooltip>
        <el-tooltip content="密度" placement="top">
          <el-dropdown @command="handleTableSizeCommand">
            <el-button
              class="el-dropdown-link"
              :icon="Expand"
              trigger="click"
            />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="(item, _) in tableSizeOptions"
                  :command="item.value"
                  >{{ item.label }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>

        <el-popover placement="bottom" :width="180" trigger="hover">
          <template #reference>
            <span>
              <el-tooltip content="设置" placement="top">
                <el-button :icon="Setting" @click="handleFilter" />
              </el-tooltip>
            </span>
          </template>
          <div
            class="operation-settings-show"
            style="
              display: flex;
              justify-content: space-between;
              align-items: center;
            "
          >
            <el-checkbox
              v-model="checkAll"
              :indeterminate="isIndeterminate"
              @change="handleCheckAllChange"
              >列展示
            </el-checkbox>
            <el-button type="primary" text>重置</el-button>
          </div>

          <el-divider style="margin: 4px 0" />

          <el-checkbox-group v-model="tableColsCheckList">
            <el-checkbox label="Option A" />
            <el-checkbox label="Option B" />
            <el-checkbox label="Option C" />
            <el-checkbox label="disabled" disabled />
            <el-checkbox label="selected and disabled" disabled />
          </el-checkbox-group>
        </el-popover>
        <el-tooltip content="全屏" placement="top">
          <el-button :icon="FullScreen" @click="handleScreenFull" />
        </el-tooltip>
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

    <el-table
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column property="id" label="角色ID" width="80" />
      <el-table-column property="name" label="角色名称" show-overflow-tooltip />
      <el-table-column
        property="sort"
        label="排序"
        show-overflow-tooltip
        width="80"
      />
      <el-table-column
        property="status"
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
      <el-table-column property="note" label="备注" show-overflow-tooltip />
      <el-table-column property="created_at" label="创建时间" width="165" />
      <el-table-column property="updated_at" label="更新时间" width="165" />
      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            @click="handleEdit(scope.row)"
            >编辑
          </el-button>
          <el-button
            link
            type="primary"
            size="small"
            @click="handleDelete(scope.row)"
            >删除
          </el-button>
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
import {
  Search,
  Delete,
  Plus,
  Setting,
  RefreshRight,
  Expand,
  FullScreen,
} from '@element-plus/icons-vue';
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { ElMessage } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getRoleList,
  updateRoleStatus,
  deleteRole,
  batchDelete,
} from '@/api/system/role';
import { RoleListRsp, Role } from '~/api/system/role';
import Pagination from '@/components/Pagination.vue';
import RoleForm from './components/RoleForm.vue';

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

// 表格尺寸列表
const tableSizeOptions = [
  {
    label: '宽松',
    value: 'large',
  },
  {
    label: '默认',
    value: 'default',
  },
  {
    label: '紧凑',
    value: 'small',
  },
];
// 表格尺寸
const tableSize = ref(settings.value.defaultSize);
// 表格尺寸选择事件
const handleTableSizeCommand = (data: string) => {
  tableSize.value = data;
};

const checkAll = ref(false);
const isIndeterminate = ref(true);
const tableColsCheckList = ref(['selected and disabled', 'Option A']);
const handleCheckAllChange = (val: boolean) => {
  // checkedCities.value = val ? cities : [];
  isIndeterminate.value = false;
};

// 全屏
const screenFullFlag = ref(false);
const handleScreenFull = () => {
  let element = document.getElementById('role-manage');
  if (!element) {
    return;
  }
  // 不全屏是null,返回false,
  screenFullFlag.value = document.fullscreenElement === null ? false : true;
  // false是进入全屏状态
  if (screenFullFlag.value) {
    // 退出全屏
    if (document.exitFullscreen) {
      document.exitFullscreen();
    }
  } else {
    // 全屏
    element.requestFullscreen();
  }
  // 切换文本状态（只是用在文本上，文本不是动态可以忽略）
  screenFullFlag.value = !screenFullFlag.value;
};

const state = reactive({
  roleForm: {
    data: {
      sort: 1,
    } as Role,
    visible: false,
    type: 'add',
    width: '500px',
  },
});
const tableData = ref<Role[]>();
const tableDataTotal = ref<number>(0);

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
  state.roleForm.visible = true;
  state.roleForm.data = row;
  state.roleForm.type = 'edit';
};
// 添加
const handleAdd = async () => {
  state.roleForm.visible = true;
};
// 批量删除
const handleBatchDelete = async () => {
  const data = {
    ids: [],
  };
  try {
    await batchDelete(data);
    fetchRoleList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
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

:deep(.operation-settings-show) {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
