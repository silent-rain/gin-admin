<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:config:list')" class="filter">
      <label>一级配置筛选:</label>
      <el-input
        v-model="listQuery.name"
        class="filter-name"
        :disabled="isDisabledButton('sys:config:list')"
        placeholder="请输入一级配置名称"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:config:list')"
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
          permission="sys:config:add"
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
              permission="sys:config:delall"
              type="danger"
              :icon="Delete"
            >
              批量删除
            </ButtonPermission>
          </template>
        </el-popconfirm>
        <ButtonPermission
          permission="sys:config:expand"
          type=""
          @click="handleExpandAllEvent(true)"
        >
          全部展开
        </ButtonPermission>
        <ButtonPermission
          permission="sys:config:collapse"
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
          :screen-full-element="'el-table-config'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchConfigTree"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <ConfigForm
      v-model:data="state.formData.data"
      v-model:visible="state.formData.visible"
      :type="state.formData.type"
      @refresh="fetchConfigTree"
    />

    <el-table
      ref="tableRef"
      class="el-table-config"
      :data="tableData"
      :size="tableSize"
      row-key="id"
      style="width: 100%; margin-top: 10px"
      :default-expand-all="tableExpandAll"
      @selection-change="handleSelectionChange"
    >
      <el-table-column v-if="checkedDict.id" prop="id" label="ID" />
      <el-table-column
        v-if="checkedDict.name"
        prop="name"
        label="配置名称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.key"
        prop="key"
        label="配置KEY"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.value"
        prop="value"
        label="配置值"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.status"
        prop="status"
        label="配置状态"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            :disabled="isDisabledButton('sys:config:status')"
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
        width="186"
      >
        <template #default="scope">
          <ButtonPermission
            permission="sys:config:addchild"
            link
            type="primary"
            size="small"
            :icon="Plus"
            @click="handleAddById(scope.row)"
          >
            添加
          </ButtonPermission>
          <ButtonPermission
            permission="sys:config:update"
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
                permission="sys:config:delete"
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
      @pagination="fetchConfigTree"
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
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage, TableInstance } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getConfigTree,
  updateConfigStatus,
  deleteConfig,
  batchDeleteConfig,
} from '@/api/data-center/config';
import { ConfigListRsp, Config } from '~/api/data-center/config';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ButtonPermission from '@/components/ButtonPermission.vue';
import ConfigForm from './components/ConfigForm.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const basicStore = useBasicStore();

// 筛选过滤条件
const listQuery = reactive({
  page: 1,
  page_size: 10,
  name: '',
});
// 过滤事件
const handleFilter = () => {
  fetchConfigTree();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.name = '';
};

const state = reactive({
  formData: {
    data: {} as Config,
    visible: false,
    type: '',
  },
});

const checkAllList = [
  { label: '配置ID', value: 'id', disabled: true, enabled: true },
  { label: '配置名称', value: 'name', disabled: true, enabled: true },
  { label: '配置KEY', value: 'key', disabled: true, enabled: true },
  { label: '配置值', value: 'value', disabled: true, enabled: true },
  { label: '排序', value: 'sort', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '备注', value: 'note', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(basicStore.settings.defaultSize);
const tableExpandAll = ref<boolean>(false);
const tableRef = ref<TableInstance>();
const tableData = ref<Config[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<Config[]>([]);

onBeforeMount(() => {
  fetchConfigTree();
});

// 获取配置树
const fetchConfigTree = async () => {
  try {
    const resp = (await getConfigTree(listQuery)).data as ConfigListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: Config) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteConfig(data);
    fetchConfigTree();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: Config) => {
  state.formData.data = { ...row };
  state.formData.type = 'edit';
  state.formData.visible = true;
};
// 添加
const handleAdd = async () => {
  state.formData.data.status = 1;
  state.formData.data.sort = 1;
  state.formData.type = 'add';
  state.formData.visible = true;
};
// 指定上级配置添加
const handleAddById = async (row: Config) => {
  state.formData.data.parent_id = row.id;
  state.formData.data.status = 1;
  state.formData.data.sort = 1;
  state.formData.type = 'add';
  state.formData.visible = true;
};
// 多选事件
const handleSelectionChange = (val: Config[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: Config) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteConfig(data);
    fetchConfigTree();
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
const handleStatusChange = async (row: Config) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    const resp = (await updateConfigStatus(data)).data as ConfigListRsp;
    tableData.value = resp.data_list;
    fetchConfigTree();
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
  dataList: Config[] | undefined,
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
