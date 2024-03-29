<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:menu:list')" class="filter">
      <label>一级菜单筛选:</label>
      <el-input
        v-model="listQuery.title"
        class="filter-name"
        :disabled="isDisabledButton('sys:menu:list')"
        placeholder="请输入一级菜单名称"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:menu:list')"
          @click="handleFilter"
        >
          查询
        </el-button>
        <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
      </el-button-group>
    </div>

    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button">
        <ButtonPermission
          permission="sys:menu:add"
          type="primary"
          :icon="Plus"
          @click="handleAdd"
        >
          添加
        </ButtonPermission>
        <el-popconfirm
          confirm-button-text="确认"
          cancel-button-text="取消"
          :icon="InfoFilled"
          icon-color="#E6A23C"
          title="确定删除吗?"
          @confirm="handleBatchDelete"
          @cancel="handleBatchDeleteCancel"
        >
          <template #reference>
            <ButtonPermission
              permission="sys:menu:delall"
              type="danger"
              :icon="Delete"
            >
              批量删除
            </ButtonPermission>
          </template>
        </el-popconfirm>
        <ButtonPermission
          permission="sys:menu:expand"
          type=""
          @click="handleExpandAllEvent(true)"
        >
          全部展开
        </ButtonPermission>
        <ButtonPermission
          permission="sys:menu:collapse"
          type=""
          @click="handleExpandAllEvent(false)"
        >
          全部折叠
        </ButtonPermission>
      </div>
      <div class="right-button">
        <ConvenienTools
          v-model:size="tableSize"
          v-model:checkedDict="checkedDict"
          :screen-full-element="'el-table-menu'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchMenuTree"
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
      ref="tableRef"
      class="el-table-menu"
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
        width="120"
      />
      <el-table-column
        v-if="checkedDict.menu_type"
        prop="menu_type"
        label="菜单类型"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag v-if="scope.row.menu_type === MenuType.Button" type="success">
            按钮
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
          <el-tag v-if="scope.row.open_type === OpenType.Component">
            菜单
          </el-tag>
          <el-tag
            v-else-if="scope.row.open_type === OpenType.Link"
            type="success"
          >
            按钮
          </el-tag>
          <el-tag
            v-else-if="scope.row.open_type === OpenType.OuterLink"
            type="info"
          >
            按钮
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.name"
        prop="name"
        label="路由别名"
        show-overflow-tooltip
      />
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
        v-if="checkedDict.redirect"
        prop="redirect"
        label="路由重定向"
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
        v-if="checkedDict.hidden"
        prop="hidden"
        label="是否隐藏"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag
            v-if="scope.row.menu_type === MenuType.Button && scope.row.hidden"
            type="info"
          >
            禁用
          </el-tag>
          <el-tag
            v-else-if="
              scope.row.menu_type === MenuType.Button && !scope.row.hidden
            "
            type="success"
          >
            可用
          </el-tag>
          <el-tag v-else-if="scope.row.hidden">隐藏</el-tag>
          <el-tag v-else type="success">显示</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.always_show"
        prop="always_show"
        label="显示根菜单"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-tag v-if="scope.row.always_show" type="success">显示</el-tag>
          <el-tag v-else>隐藏</el-tag>
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
            :disabled="isDisabledButton('sys:menu:status')"
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
        :fixed="basicStore.isMobile() ? false : 'right'"
        label="操作"
        align="center"
        width="210"
      >
        <template #default="scope">
          <ButtonPermission
            permission="sys:menu:addchild"
            link
            type="primary"
            size="small"
            :icon="Plus"
            @click="handleAddById(scope.row)"
          >
            添加
          </ButtonPermission>
          <ButtonPermission
            permission="sys:menu:update"
            link
            type="primary"
            size="small"
            :icon="DocumentCopy"
            @click="handleCopy(scope.row)"
          >
            拷贝
          </ButtonPermission>
          <ButtonPermission
            permission="sys:menu:update"
            link
            type="primary"
            size="small"
            :icon="EditPen"
            @click="handleEdit(scope.row)"
          >
            修改
          </ButtonPermission>
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
              <ButtonPermission
                permission="sys:menu:delete"
                link
                type="danger"
                size="small"
                :icon="Delete"
              >
                删除
              </ButtonPermission>
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
import {
  EditPen,
  Search,
  Delete,
  Plus,
  DocumentCopy,
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage, TableInstance } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getMenuTree,
  updateMenuStatus,
  deleteMenu,
  batchDeleteMenu,
} from '@/api/permission/menu';
import { MenuListRsp, Menu } from '~/api/permission/menu';
import { MenuType, OpenType } from '@/constant/permission/menu';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ButtonPermission from '@/components/ButtonPermission.vue';
import MenuForm from './components/MenuForm.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const basicStore = useBasicStore();

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
  { label: '路由别名', value: 'name', disabled: true, enabled: false },
  { label: '路由地址', value: 'path', disabled: true, enabled: true },
  { label: '组件路径', value: 'component', disabled: false, enabled: true },
  { label: '路由重定向', value: 'redirect', disabled: false, enabled: false },
  { label: '链接地址', value: 'link', disabled: false, enabled: true },
  { label: '权限标识', value: 'permission', disabled: true, enabled: true },
  { label: '是否隐藏', value: 'hidden', disabled: false, enabled: true },
  {
    label: '显示根菜单',
    value: 'always_show',
    disabled: false,
    enabled: false,
  },
  { label: '排序', value: 'sort', disabled: false, enabled: false },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: false },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(basicStore.settings.defaultSize);
const tableExpandAll = ref<boolean>(false);
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
  state.menuForm.data.hidden = 0;
  state.menuForm.data.always_show = 1;
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
  state.menuForm.data.hidden = 0;
  state.menuForm.data.always_show = 1;
  state.menuForm.data.status = 1;
  state.menuForm.data.sort = 1;
  state.menuForm.type = 'add';
  state.menuForm.visible = true;
};
// 拷贝当前菜单
const handleCopy = async (row: Menu) => {
  state.menuForm.data = { ...row };

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
// 取消批量删除事件
const handleBatchDeleteCancel = () => {
  ElMessage.warning('取消操作');
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
  flex-wrap: wrap;

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
