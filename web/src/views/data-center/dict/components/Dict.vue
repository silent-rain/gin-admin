<script setup lang="ts">
import type { Dict, DictListRsp } from '~/api/data-center/dict'
import { ArrowRight, InfoFilled } from '@element-plus/icons-vue'
import { ElMessage, ElTable } from 'element-plus'
import { ref } from 'vue'
import { deleteDict, getDictList } from '@/api/data-center/dict'
import ButtonPermission from '@/components/ButtonPermission.vue'
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission'
import DictForm from './DictForm.vue'

const emits = defineEmits(['update:dictId'])

// 筛选过滤条件
const listQuery = ref({
  page: 1,
  page_size: 10,
  name: null,
  code: null,
})

const state = reactive({
  form: {
    data: {} as Dict,
    visible: false,
    type: '',
  },
})

const currentRow = ref()
const singleTableRef = ref<InstanceType<typeof ElTable>>()
const tableData = ref<Dict[]>([])
const tableDataTotal = ref<number>(0)

onBeforeMount(() => {
  fetchDictList()
})

// 获取角色列表
async function fetchDictList() {
  try {
    const resp = (await getDictList(listQuery.value)).data as DictListRsp
    tableData.value = resp.data_list
    tableDataTotal.value = resp.tatol

    // 默认设置选中第一行
    if (tableData.value.length > 0) {
      setCurrent(tableData.value[0])
    }
  }
  catch (error) {
    console.log(error)
  }
}

// 过滤事件
function handleFilter() {
  fetchDictList()
}

// 设置选中行
function setCurrent(row?: Dict) {
  singleTableRef.value!.setCurrentRow(row)
  emits('update:dictId', row?.id)
}
// 清空过滤条件
function handleCleanFilter() {
  listQuery.value = {} as any
  fetchDictList()
}

// 选择所在的行
function handleCurrentChange(val: Dict | undefined) {
  currentRow.value = val
  emits('update:dictId', val?.id)
}

// 添加
function handleAdd() {
  state.form.type = 'add'
  state.form.visible = true
}
// 编辑
function handleEdit() {
  if (!currentRow.value) {
    return
  }
  state.form.data = { ...currentRow.value }
  state.form.type = 'edit'
  state.form.visible = true
}
// 删除
async function handleDel() {
  if (!currentRow.value) {
    return
  }
  const data = {
    id: currentRow.value.id,
  }
  try {
    await deleteDict(data)
    fetchDictList()
    ElMessage.success('操作成功')
  }
  catch (error) {
    console.log(error)
  }
}
// 取消删除
function handleDelCancel() {
  ElMessage.warning('取消操作')
}
</script>

<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('dataCenter:dict:list')" class="filter">
      <label>字典筛选:</label>
      <el-input
        v-model="listQuery.name"
        class="filter-name"
        clearable
        :disabled="isDisabledButton('dataCenter:dict:list')"
        placeholder="筛选字典名称"
        @keyup.enter.native="handleFilter"
        @clear="handleCleanFilter"
      />
      <el-input
        v-model="listQuery.code"
        class="filter-name"
        clearable
        :disabled="isDisabledButton('dataCenter:dict:list')"
        placeholder="筛选字典编码"
        @keyup.enter.native="handleFilter"
        @clear="handleCleanFilter"
      />
      <!-- <div class="filter-button">
        <el-button-group>
          <el-button
            type="primary"
            :icon="Search"
            :disabled="isDisabledButton('dataCenter:dict:list')"
            @click="handleFilter"
          >
            查询
          </el-button>
          <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
        </el-button-group>
      </div> -->
    </div>

    <el-divider />

    <!-- 操作按钮 -->
    <div class="operation-button">
      <ButtonPermission
        permission="dataCenter:dict:add"
        type="primary"
        @click="handleAdd"
      >
        添加
      </ButtonPermission>
      <ButtonPermission
        permission="dataCenter:dict:update"
        type="warning"
        @click="handleEdit"
      >
        修改
      </ButtonPermission>
      <el-popconfirm
        confirm-button-text="确认"
        cancel-button-text="取消"
        :icon="InfoFilled"
        icon-color="#E6A23C"
        title="确定删除吗?"
        @confirm="handleDel"
        @cancel="handleDelCancel"
      >
        <template #reference>
          <ButtonPermission
            permission="dataCenter:dict:delall"
            type="danger"
            icon=""
          >
            删除
          </ButtonPermission>
        </template>
      </el-popconfirm>
    </div>

    <!-- 表单 -->
    <DictForm
      v-if="state.form.visible"
      v-model:data="state.form.data"
      v-model:visible="state.form.visible"
      :type="state.form.type"
      @refresh="fetchDictList"
    />

    <ElTable
      ref="singleTableRef"
      class="content"
      highlight-current-row
      border
      :data="tableData"
      @current-change="handleCurrentChange"
    >
      <el-table-column type="index" width="50" />
      <el-table-column prop="name" label="字典名称">
        <template #default="scope">
          <div class="dict-name-item">
            <span>{{ scope.row.name }}</span>
            <span>
              <el-icon><ArrowRight /></el-icon>
            </span>
          </div>
        </template>
      </el-table-column>
    </ElTable>
    <Pagination
      v-model:current-page="listQuery.page"
      v-model:page-size="listQuery.page_size"
      :total="tableDataTotal"
      layout="sizes, prev, pager, next"
      @pagination="fetchDictList"
    />
  </el-card>
</template>

<style scoped lang="scss">
.filter {
  .filter-button {
    margin-top: 5px;
    display: flex;
    justify-content: flex-end;
  }
}
.operation-button {
  margin-top: 10px;
}
.dict-name-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.content {
  margin-top: 10px;
}
</style>
