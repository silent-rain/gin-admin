<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:httplog:list')" class="filter">
      <el-input
        v-model="listQuery.user_id"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请输入用户ID"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.trace_id"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请输入 Trace Id"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.status_code"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请输入状态码"
        @keyup.enter.native="handleFilter"
      />
      <el-select
        v-model="listQuery.method"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请选择请求方法"
        @change="handleFilter"
      >
        <el-option label="GET" value="GET" />
        <el-option label="POST" value="POST" />
        <el-option label="PUT" value="PUT" />
        <el-option label="DELETE" value="DELETE" />
        <el-option label="OPTIONS" value="OPTIONS" />
      </el-select>
      <el-input
        v-model="listQuery.path"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请输入请求地址路径"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.remote_addr"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请输入请求IP"
        @keyup.enter.native="handleFilter"
      />
      <el-select
        v-model="listQuery.htpp_type"
        class="filter-name"
        :disabled="isDisabledButton('sys:httplog:list')"
        placeholder="请选择请求类型"
        @change="handleFilter"
      >
        <el-option label="REQ" value="REQ" />
        <el-option label="RESP" value="RESP" />
      </el-select>
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:httplog:list')"
          @click="handleFilter"
        >
          查询
        </el-button>
        <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
      </el-button-group>
    </div>

    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button"></div>
      <div class="right-button">
        <ConvenienTools
          v-model:size="tableSize"
          v-model:checkedDict="checkedDict"
          :screen-full-element="'el-table-role'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchHttpLogList"
        />
      </div>
    </div>

    <el-table
      class="el-table-role"
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
    >
      <el-table-column v-if="checkedDict.id" prop="id" label="ID" width="80" />
      <el-table-column
        v-if="checkedDict.user_id"
        prop="user_id"
        label="用户ID"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.trace_id"
        prop="trace_id"
        label="Trace ID"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.span_id"
        prop="span_id"
        label="Span ID"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.status_code"
        prop="status_code"
        label="状态码"
        show-overflow-tooltip
        width="80"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-if="scope.row.status_code === 200"
            type="success"
            size="small"
          >
            {{ scope.row.status_code }}
          </el-tag>
          <el-tag v-else size="small" type="danger">
            {{ scope.row.status_code }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.method"
        prop="method"
        label="请求方法"
        show-overflow-tooltip
        width="80"
        align="center"
      />
      <el-table-column
        v-if="checkedDict.path"
        prop="path"
        label="路径"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.query"
        prop="query"
        label="请求参数"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.body"
        prop="body"
        label="请求体/响应体"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.remote_addr"
        prop="remote_addr"
        label="请求IP"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.user_agent"
        prop="user_agent"
        label="用户代理"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.cost"
        prop="cost"
        label="耗时(纳秒)"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.htpp_type"
        prop="htpp_type"
        label="请求类型"
        show-overflow-tooltip
        width="80"
        align="center"
      >
        <template #default="scope">
          <el-tag v-if="scope.row.htpp_type === 'REQ'" size="small">
            {{ scope.row.htpp_type }}
          </el-tag>
          <el-tag v-else size="small" type="success">
            {{ scope.row.htpp_type }}
          </el-tag>
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
    </el-table>
    <Pagination
      v-model:currentPage="listQuery.page"
      v-model:pageSize="listQuery.page_size"
      :total="tableDataTotal"
      @pagination="fetchHttpLogList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { Search, Delete } from '@element-plus/icons-vue';
import { useBasicStore } from '@/store/basic';
import { getHttpLogList, getSystemLogList } from '@/api/system/log';
import { RoleListRsp, Role } from '~/api/permission/role';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = ref<any>({
  page: 1,
  page_size: 10,
  user_id: null,
  trace_id: null,
  status_code: null,
  method: '',
  path: '',
  remote_addr: '',
  htpp_type: '',
});
// 过滤事件
const handleFilter = () => {
  fetchHttpLogList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {} as any;
};

const checkAllList = [
  { label: '日志ID', value: 'id', disabled: false, enabled: false },
  { label: '用户ID', value: 'user_id', disabled: true, enabled: true },
  { label: 'Trace ID', value: 'trace_id', disabled: false, enabled: true },
  { label: 'Span ID', value: 'span_id', disabled: false, enabled: true },
  { label: '状态码', value: 'status_code', disabled: false, enabled: true },
  { label: '路径', value: 'method', disabled: false, enabled: true },
  { label: '请求方法', value: 'path', disabled: false, enabled: true },
  { label: '请求参数', value: 'query', disabled: false, enabled: false },
  { label: '请求体/响应体', value: 'body', disabled: false, enabled: false },
  { label: '请求IP', value: 'remote_addr', disabled: false, enabled: false },
  { label: '用户代理', value: 'user_agent', disabled: false, enabled: true },
  { label: '耗时(纳秒)', value: 'cost', disabled: false, enabled: false },
  { label: '请求类型', value: 'htpp_type', disabled: false, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: false },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});
const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<Role[]>();
const tableDataTotal = ref<number>(0);

onBeforeMount(() => {
  fetchHttpLogList();
});

// 获取网络请求日志列表
const fetchHttpLogList = async () => {
  try {
    const resp = (await getHttpLogList(listQuery.value)).data as RoleListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};
</script>

<style scoped lang="scss">
.filter {
  .filter-name {
    width: 200px;
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
