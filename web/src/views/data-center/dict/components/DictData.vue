<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('dataCenter:dictData:list')" class="filter">
      <label>字典筛选:</label>
      <el-input
        v-model="listQuery.name"
        class="filter-name"
        :disabled="isDisabledButton('dataCenter:dictData:list')"
        placeholder="筛选字典项名称"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.value"
        class="filter-name"
        :disabled="isDisabledButton('dataCenter:dictData:list')"
        placeholder="筛选字典项值"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('dataCenter:dictData:list')"
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
          permission="dataCenter:dictData:add"
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
              permission="dataCenter:dictData:delall"
              type="danger"
              :icon="Delete"
            >
              批量删除
            </ButtonPermission>
          </template>
        </el-popconfirm>
      </div>
      <div class="right-button">
        <ConvenienTools
          v-model:size="tableSize"
          v-model:checkedDict="checkedDict"
          :screen-full-element="'el-table-full'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchDictDataList"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <DictDataForm
      v-if="state.form.visible"
      v-model:data="state.form.data"
      v-model:visible="state.form.visible"
      :type="state.form.type"
      @refresh="fetchDictDataList"
    />

    <el-table
      class="el-table-full"
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column
        v-if="checkedDict.id"
        prop="id"
        label="字典项ID"
        width="90"
      />
      <el-table-column
        v-if="checkedDict.dict_id"
        prop="dict_id"
        label="字典维度ID"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.name"
        prop="name"
        label="字典项名称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.value"
        prop="value"
        label="字典项值"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.status"
        prop="status"
        label="状态"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            :disabled="isDisabledButton('dataCenter:dictData:status')"
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
        :fixed="basicStore.isMobile() ? false : 'right'"
        label="操作"
        align="center"
        width="120"
      >
        <template #default="scope">
          <ButtonPermission
            permission="dataCenter:dictData:update"
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
                permission="dataCenter:dictData:delete"
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
      @pagination="fetchDictDataList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import {
  EditPen,
  Search,
  Delete,
  InfoFilled,
  Plus,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ButtonPermission from '@/components/ButtonPermission.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';
import { useBasicStore } from '@/store/basic';
import {
  getDictDataList,
  updateDictDataStatus,
  deleteDictData,
  batchDeleteDictData,
} from '@/api/data-center/dict-data';
import { DictDataListRsp, DictData } from '~/api/data-center/dict-data';
import DictDataForm from './DictDataForm.vue';

const props = withDefaults(
  defineProps<{
    dictId: number;
  }>(),
  {},
);

const basicStore = useBasicStore();

// 筛选过滤条件
const listQuery = ref({
  page: 1,
  page_size: 10,
  name: null,
  dict_id: 0,
  value: null,
});
// 过滤事件
const handleFilter = () => {
  fetchDictDataList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {} as any;
  fetchDictDataList();
};

const state = reactive({
  form: {
    data: {} as DictData,
    visible: false,
    type: '',
  },
});

const checkAllList = [
  { label: '字典项ID', value: 'id', disabled: false, enabled: true },
  { label: '字典维度ID', value: 'dict_id', disabled: false, enabled: false },
  { label: '字典项名称', value: 'name', disabled: false, enabled: true },
  { label: '字典项值', value: 'value', disabled: false, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(basicStore.settings.defaultSize);
const tableData = ref<DictData[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<DictData[]>([]);

onBeforeMount(() => {
  listQuery.value.dict_id = props.dictId;
  fetchDictDataList();
});

// 获取字典数据信息列表
const fetchDictDataList = async () => {
  try {
    const resp = (await getDictDataList(listQuery.value))
      .data as DictDataListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: DictData) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteDictData(data);
    fetchDictDataList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: DictData) => {
  state.form.data = { ...row };
  state.form.type = 'edit';
  state.form.visible = true;
};
// 添加
const handleAdd = async () => {
  state.form.type = 'add';
  state.form.visible = true;
  state.form.data.dict_id = props.dictId;
};
// 多选事件
const handleSelectionChange = (val: DictData[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: DictData) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteDictData(data);
    fetchDictDataList();
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
const handleStatusChange = async (row: DictData) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    await updateDictDataStatus(data);
    fetchDictDataList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

watch(
  () => props.dictId,
  () => {
    if (!props.dictId) {
      return;
    }
    listQuery.value.dict_id = props.dictId;
    fetchDictDataList();
  },
);
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
