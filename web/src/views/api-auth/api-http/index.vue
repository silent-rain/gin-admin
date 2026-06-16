<script setup lang="ts">
import type { TableInstance } from 'element-plus'
import type { ApiHttp, ApiHttpTreeRsp } from '~/api/api-auth/api-http'
import {
  Delete,
  DocumentCopy,
  EditPen,
  InfoFilled,
  Plus,
  Search,
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { onBeforeMount, reactive, ref } from 'vue'
import {
  batchDeleteApiHttp,
  deleteApiHttp,
  getApiHttpTree,
  updateApiHttpStatus,
} from '@/api/api-auth/api-http'
import ButtonPermission from '@/components/ButtonPermission.vue'
import ConvenienTools from '@/components/ConvenienTools/index.vue'
import Pagination from '@/components/Pagination.vue'
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission'
import { useBasicStore } from '@/store/basic'
import ApiHttpForm from './components/ApiHttpForm.vue'

const basicStore = useBasicStore()

// 筛选过滤条件
const listQuery = ref<any>({
  page: 1,
  page_size: 10,
  name: null,
  method: null,
  uri: null,
  status: null,
})
const statusOptions = [
  {
    label: '启用',
    value: 1,
  },
  {
    label: '禁用',
    value: 0,
  },
]
// 过滤事件
function handleFilter() {
  fetchApiHttpTree()
}
// 清空过滤条件
function handleCleanFilter() {
  listQuery.value = {}
}
// 状态变更事件-清空处理
function handleChangeStatus(value: any) {
  if (!value) {
    listQuery.value.status = null
  }
  handleFilter()
}

const state = reactive({
  form: {
    data: {} as ApiHttp,
    visible: false,
    type: '',
  },
})

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
    value: 'DELETE',
    type: 'danger',
  },
]

const checkAllList = [
  { label: '自增ID', value: 'id', disabled: false, enabled: false },
  { label: '接口名称', value: 'name', disabled: true, enabled: true },
  { label: '请求类型', value: 'method', disabled: true, enabled: true },
  { label: 'URI资源', value: 'uri', disabled: true, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
]
const checkedDict = ref<any>({})

const tableSize = ref<string>(basicStore.settings.defaultSize)
const tableData = ref<ApiHttp[]>()
const tableDataTotal = ref<number>(0)
const multipleSelection = ref<ApiHttp[]>([])
const tableRef = ref<TableInstance>()
const tableExpandAll = ref<boolean>(false)

onBeforeMount(() => {
  fetchApiHttpTree()
})

// 获取Http协议接口信息树
async function fetchApiHttpTree() {
  try {
    const resp = (await getApiHttpTree(listQuery.value)).data as ApiHttpTreeRsp
    tableData.value = resp.data_list
    tableDataTotal.value = resp.tatol
  }
  catch (error) {
    console.log(error)
  }
}

// 删除
async function handleDelete(row: ApiHttp) {
  const data = {
    id: row.id,
  }
  try {
    await deleteApiHttp(data)
    fetchApiHttpTree()
    ElMessage.success('操作成功')
  }
  catch (error) {
    console.log(error)
  }
}
// 编辑
async function handleEdit(row: ApiHttp) {
  state.form.data = { ...row }
  state.form.type = 'edit'
  state.form.visible = true
}
// 添加
async function handleAdd() {
  state.form.type = 'add'
  state.form.visible = true
  state.form.data.status = 1
}

// 指定上级菜单添加
async function handleAddById(row: ApiHttp) {
  state.form.type = 'add'
  state.form.visible = true

  state.form.data.parent_id = row.id
  state.form.data.status = 1
  state.form.type = 'add'
}
// 拷贝当前菜单
async function handleCopy(row: ApiHttp) {
  state.form.type = 'add'
  state.form.visible = true

  state.form.data = { ...row }
}

// 多选事件
function handleSelectionChange(val: ApiHttp[]) {
  multipleSelection.value = val
}
// 批量删除
async function handleBatchDelete() {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据')
    return
  }
  const data = {
    ids: multipleSelection.value.map((v: ApiHttp) => {
      return v.id
    }),
  }
  try {
    await batchDeleteApiHttp(data)
    fetchApiHttpTree()
    ElMessage.success('操作成功')
  }
  catch (error) {
    console.log(error)
  }
}
// 取消批量删除事件
function handleBatchDeleteCancel() {
  ElMessage.warning('取消操作')
}

// 删除取消事件
function handleCancelEvent() {
  ElMessage.warning('取消操作')
}

// 状态变更
async function handleStatusChange(row: ApiHttp) {
  const data = {
    id: row.id,
    status: row.status,
  }
  try {
    await updateApiHttpStatus(data)
    fetchApiHttpTree()
    ElMessage.success('操作成功')
  }
  catch (error) {
    console.log(error)
  }
}
// 全部展开/全部折叠 事件
function handleExpandAllEvent(value: boolean) {
  toggleRowExpansionAll(tableData.value, value)
}

// 全部展开/全部折叠
function toggleRowExpansionAll(dataList: ApiHttp[] | undefined, value: boolean) {
  if (!dataList) {
    return
  }
  dataList.forEach((v) => {
    tableRef.value?.toggleRowExpansion(v, value)
    if (v.children !== undefined && v.children !== null) {
      toggleRowExpansionAll(v.children, value)
    }
  })
}
</script>

<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('apiAuth:apiHttp:list')" class="filter">
      <el-input
        v-model="listQuery.name"
        class="filter-name"
        clearable
        :disabled="isDisabledButton('apiAuth:apiHttp:list')"
        placeholder="筛选接口名称"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.uri"
        class="filter-name"
        :disabled="isDisabledButton('apiAuth:apiHttp:list')"
        placeholder="筛选URI资源"
        @keyup.enter.native="handleFilter"
      />
      <el-select
        v-model="listQuery.method"
        class="filter-name"
        :disabled="isDisabledButton('apiAuth:apiHttp:list')"
        placeholder="请筛选请求类型"
        @change="handleFilter"
      >
        <el-option label="GET" value="GET" />
        <el-option label="POST" value="POST" />
        <el-option label="PUT" value="PUT" />
        <el-option label="DELETE" value="DELETE" />
        <el-option label="OPTIONS" value="OPTIONS" />
      </el-select>
      <el-select
        v-model="listQuery.status"
        class="filter-name"
        clearable
        placeholder="筛选状态"
        @change="handleChangeStatus"
      >
        <el-option
          v-for="item in statusOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('apiAuth:apiHttp:list')"
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
          permission="apiAuth:apiHttp:add"
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
              permission="apiAuth:apiHttp:delall"
              type="danger"
              :icon="Delete"
            >
              批量删除
            </ButtonPermission>
          </template>
        </el-popconfirm>
        <ButtonPermission
          permission="apiAuth:apiHttp:expand"
          type=""
          @click="handleExpandAllEvent(true)"
        >
          全部展开
        </ButtonPermission>
        <ButtonPermission
          permission="apiAuth:apiHttp:collapse"
          type=""
          @click="handleExpandAllEvent(false)"
        >
          全部折叠
        </ButtonPermission>
      </div>
      <div class="right-button">
        <ConvenienTools
          v-model:size="tableSize"
          v-model:checked-dict="checkedDict"
          screen-full-element="el-table-full"
          :check-all-list="checkAllList"
          @refresh-event="fetchApiHttpTree"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <ApiHttpForm
      v-if="state.form.visible"
      v-model:data="state.form.data"
      v-model:visible="state.form.visible"
      :type="state.form.type"
      @refresh="fetchApiHttpTree"
    />

    <el-table
      ref="tableRef"
      class="el-table-full"
      :data="tableData"
      :size="tableSize"
      row-key="id"
      style="width: 100%; margin-top: 10px"
      :default-expand-all="tableExpandAll"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column
        v-if="checkedDict.id"
        prop="id"
        label="自增ID"
        width="80"
      />
      <el-table-column
        v-if="checkedDict.name"
        prop="name"
        label="接口名称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.method"
        prop="method"
        label="请求类型"
        show-overflow-tooltip
        width="100"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-for="(item, _) in methodOptions.filter(
              (v) => v.value === scope.row.method,
            )"
            :key="item.value"
            size="small"
            :type="item.type"
          >
            {{ scope.row.method }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.uri"
        prop="uri"
        label="URI资源"
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
            :disabled="isDisabledButton('apiAuth:apiHttp:status')"
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
        width="210"
      >
        <template #default="scope">
          <ButtonPermission
            permission="apiAuth:apiHttp:addchild"
            link
            type="primary"
            size="small"
            :icon="Plus"
            @click="handleAddById(scope.row)"
          >
            添加
          </ButtonPermission>
          <ButtonPermission
            permission="apiAuth:apiHttp:update"
            link
            type="primary"
            size="small"
            :icon="DocumentCopy"
            @click="handleCopy(scope.row)"
          >
            拷贝
          </ButtonPermission>
          <ButtonPermission
            permission="apiAuth:apiHttp:update"
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
                permission="apiAuth:apiHttp:delete"
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
      v-model:current-page="listQuery.page"
      v-model:page-size="listQuery.page_size"
      :total="tableDataTotal"
      @pagination="fetchApiHttpTree"
    />
  </el-card>
</template>

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
