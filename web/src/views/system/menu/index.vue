<template>
  <el-card>
    <!-- 过滤条件 -->
    <div class="filter">
      <label>一级菜单筛选: </label>
      <el-input
        v-model="listQuery.title"
        class="filter-name"
        placeholder="请输入一级菜单名称"
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
          :buttonList="['add', 'expand', 'collapse']"
          @add-event="handleAdd"
          @batch-delete-event="handleBatchDelete"
          @expandEvent="handleExpandAllEvent"
        >
          <template v-slot:import> 导入菜单 </template>
          <template v-slot:export> 导出菜单 </template>
        </ConvenienButtons>
      </div>
      <div class="right-button">
        <ConvenienTools
          @refreshEvent="fetchMenuTree"
          v-model:size="tableSize"
          :screenFullElement="'el-table-menu'"
          :checkAllList="checkAllList"
          v-model:checkedDict="checkedDict"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <MenuForm
      v-model:data="state.menuForm.data"
      v-model:visible="state.menuForm.visible"
      :type="state.menuForm.type"
      :width="state.menuForm.width"
      @refresh="fetchMenuTree"
    />

    <el-table
      class="el-table-menu"
      ref="tableRef"
      :data="tableData"
      :size="tableSize"
      row-key="id"
      style="width: 100%; margin-top: 10px"
      :default-expand-all="tableExpandAll"
      @selection-change="handleSelectionChange"
    >
      <el-table-column v-if="checkedDict.id" prop="id" label="ID" />
      <el-table-column
        v-if="checkedDict.title"
        prop="title"
        label="菜单名称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.menu_type"
        prop="menu_type"
        label="菜单类型"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag v-if="scope.row.menu_type === MenuType.Button" type="success"
            >按钮
          </el-tag>
          <el-tag v-else>菜单</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.open_type"
        prop="open_type"
        label="打开方式"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag v-if="scope.row.open_type === OpenType.Component"
            >菜单</el-tag
          >
          <el-tag
            v-else-if="scope.row.open_type === OpenType.Link"
            type="success"
            >按钮
          </el-tag>
          <el-tag
            v-else-if="scope.row.open_type === OpenType.OuterLink"
            type="info"
            >按钮
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.path"
        prop="path"
        label="路由地址"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.component"
        prop="component"
        label="组件路径"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.link"
        prop="link"
        label="链接地址"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.permission"
        prop="permission"
        label="权限标识"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.hide"
        prop="hide"
        label="是否可见"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag v-if="scope.row.hide">显示</el-tag>
          <el-tag v-else type="success">隐藏 </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.status"
        prop="status"
        label="菜单状态"
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
        v-if="checkedDict.sort"
        prop="sort"
        label="排序"
        show-overflow-tooltip
        width="80"
      />
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
            :icon="Plus"
            @click="handleAddById(scope.row)"
            >添加
          </el-button>
          <el-button
            link
            type="primary"
            size="small"
            :icon="EditPen"
            @click="handleEdit(scope.row)"
            >修改
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
      @pagination="fetchMenuTree"
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
  Plus,
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage, TableInstance } from 'element-plus';
import {
  getMenuTree,
  updateMenuStatus,
  deleteMenu,
  batchDeleteMenu,
} from '@/api/system/menu';
import { MenuListRsp, Menu } from '~/api/system/menu';
import { MenuType, OpenType } from '@/constant/system/menu';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ConvenienButtons from '@/components/ConvenienButtons/index.vue';
import MenuForm from './components/MenuForm.vue';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = reactive({
  title: '',
  page: 1,
  page_size: 10,
});
// 过滤事件
const handleFilter = () => {
  fetchMenuTree();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.title = '';
};

const state = reactive({
  menuForm: {
    data: {} as Menu,
    visible: false,
    type: '',
    width: '700px',
  },
});

const checkAllList = [
  { label: '菜单ID', value: 'id', disabled: true, enabled: true },
  { label: '菜单名称', value: 'title', disabled: true, enabled: true },
  { label: '菜单类型', value: 'menu_type', disabled: true, enabled: true },
  { label: '打开方式', value: 'open_type', disabled: false, enabled: false },
  { label: '路由地址', value: 'path', disabled: true, enabled: true },
  { label: '组件路径', value: 'component', disabled: false, enabled: true },
  { label: '链接地址', value: 'link', disabled: false, enabled: true },
  { label: '权限标识', value: 'permission', disabled: true, enabled: true },
  { label: '是否可见', value: 'hide', disabled: false, enabled: true },
  { label: '排序', value: 'sort', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: false },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(settings.value.defaultSize);
const tableExpandAll = ref<boolean>(true);
const tableRef = ref<TableInstance>();
const tableData = ref<Menu[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<Menu[]>([]);

onBeforeMount(() => {
  fetchMenuTree();
});

// 获取菜单树
const fetchMenuTree = async () => {
  try {
    const resp = (await getMenuTree(listQuery)).data as MenuListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: Menu) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteMenu(data);
    fetchMenuTree();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: Menu) => {
  state.menuForm.data = { ...row };
  state.menuForm.type = 'edit';
  state.menuForm.visible = true;
};
// 添加
const handleAdd = async () => {
  state.menuForm.data.menu_type = 0;
  state.menuForm.data.open_type = 0;
  state.menuForm.data.hide = 1;
  state.menuForm.data.status = 1;
  state.menuForm.data.sort = 1;
  state.menuForm.type = 'add';
  state.menuForm.visible = true;
};
// 指定上级菜单添加
const handleAddById = async (row: Menu) => {
  state.menuForm.data.parent_id = row.id;
  state.menuForm.data.menu_type = 0;
  state.menuForm.data.open_type = 0;
  state.menuForm.data.hide = 1;
  state.menuForm.data.status = 1;
  state.menuForm.data.sort = 1;
  state.menuForm.type = 'add';
  state.menuForm.visible = true;
};
// 多选事件
const handleSelectionChange = (val: Menu[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: Menu) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteMenu(data);
    fetchMenuTree();
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
const handleStatusChange = async (row: Menu) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    const resp = (await updateMenuStatus(data)).data as MenuListRsp;
    tableData.value = resp.data_list;
    fetchMenuTree();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 全部展开/全部折叠 事件
const handleExpandAllEvent = (value: boolean) => {
  toggleRowExpansionAll(tableData.value, value);
};

// 全部展开/全部折叠
const toggleRowExpansionAll = (
  dataList: Menu[] | undefined,
  value: boolean,
) => {
  if (!dataList) {
    return;
  }
  dataList.forEach((v) => {
    tableRef.value?.toggleRowExpansion(v, value);
    if (v.children !== undefined && v.children !== null) {
      toggleRowExpansionAll(v.children, value);
    }
  });
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
