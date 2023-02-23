<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:systemlog:list')" class="filter">
      <!-- <el-input
        v-model="listQuery.user_id"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请输入用户ID"
        @keyup.enter.native="handleFilter"
      /> -->
      <el-input
        v-model="listQuery.trace_id"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请输入 Trace Id"
        @keyup.enter.native="handleFilter"
      />
      <el-select
        v-model="listQuery.level"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请选择日志级别"
        @change="handleFilter"
      >
        <el-option
          v-for="(item, _) in levelOptions"
          :key="item.level"
          :label="item.lebal"
          :value="item.level"
        />
      </el-select>
      <!-- <el-input
        v-model="listQuery.error_code"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请输入业务错误码"
        @keyup.enter.native="handleFilter"
      /> -->
      <el-input
        v-model="listQuery.error_msg"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请输入业务错误信息"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.msg"
        class="filter-name"
        :disabled="isDisabledButton('sys:systemlog:list')"
        placeholder="请输入日志消息"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:systemlog:list')"
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
          @refreshEvent="fetchSystemLogList"
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
        v-if="checkedDict.nickname"
        prop="nickname"
        label="用户昵称"
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
        v-if="checkedDict.level"
        prop="level"
        label="日志级别"
        show-overflow-tooltip
        width="100"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-for="(item, _) in levelOptions.filter(
              (v) => v.level === scope.row.level,
            )"
            size="small"
            :key="item.level"
            :type="item.type"
          >
            {{ scope.row.level }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.caller_line"
        prop="caller_line"
        label="日志位置"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.error_code"
        prop="error_code"
        label="业务码"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.error_msg"
        prop="error_msg"
        label="业务码信息"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.msg"
        prop="msg"
        label="日志消息"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.stack"
        prop="stack"
        label="堆栈信息"
        show-overflow-tooltip
        width="100"
        align="center"
      >
        <template #default="scope">
          <el-button
            v-if="scope.row.stack"
            type="primary"
            text
            @click="handleShowStack(scope.row.stack)"
          >
            查看
          </el-button>
          <span v-else></span>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.extend"
        prop="extend"
        label="扩展信息"
        show-overflow-tooltip
      >
        <template #default="scope">
          <el-button
            v-if="scope.row.extend !== '{}'"
            type="primary"
            text
            @click="handleShowExtend(scope.row.extend)"
          >
            查看
          </el-button>
          <span v-else></span>
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
      @pagination="fetchSystemLogList"
    />

    <!-- 扩展信息 -->
    <LogDrawer
      v-if="state.extend.visible"
      v-model="state.extend.visible"
      :data="state.extend.data"
      :key="state.extend.key"
      language="json"
    ></LogDrawer>

    <!-- 堆栈信息 -->
    <LogDrawer
      v-if="state.extend.visible"
      v-model="state.extend.visible"
      :data="state.extend.data"
      :key="state.extend.key"
      language="text"
      size="800px"
    ></LogDrawer>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { storeToRefs } from 'pinia/dist/pinia';
import { Search, Delete } from '@element-plus/icons-vue';
import { useBasicStore } from '@/store/basic';
import { getSystemLogList } from '@/api/system/log';
import { SystemLog, SystemLogListRsp } from '~/api/permission/log';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';
import LogDrawer from './LogDrawer.vue';

const { settings } = storeToRefs(useBasicStore());
const route = useRoute();
const router = useRouter();

// 筛选过滤条件
const listQuery = ref<any>({
  page: 1,
  page_size: 10,
  user_id: null,
  trace_id: null,
  level: null,
  error_code: null,
  error_msg: null,
  msg: null,
});

const state = reactive({
  extend: {
    visible: false,
    data: '',
    key: new Date().getMilliseconds(),
  },
  stack: {
    visible: false,
    data: '',
    key: new Date().getMilliseconds(),
  },
});

// 日志级别
const levelOptions = [
  {
    lebal: '调试',
    level: 'DEBUG',
    type: 'info',
  },
  {
    lebal: '信息',
    level: 'INFO',
    type: '',
  },
  {
    lebal: '警告',
    level: 'WARN',
    type: 'warning',
  },
  {
    lebal: '错误',
    level: 'ERROR',
    type: 'danger',
  },
  {
    lebal: '恐慌',
    level: 'PANIC',
    type: 'danger',
  },
];

// 过滤事件
const handleFilter = () => {
  router.push({
    path: route.path,
    query: listQuery.value,
  });
  fetchSystemLogList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {} as any;
};

const checkAllList = [
  { label: '日志ID', value: 'id', disabled: false, enabled: false },
  { label: '用户ID', value: 'user_id', disabled: false, enabled: false },
  { label: '用户昵称', value: 'nickname', disabled: true, enabled: true },
  { label: 'Trace ID', value: 'trace_id', disabled: true, enabled: true },
  { label: 'Span ID', value: 'span_id', disabled: true, enabled: true },
  { label: '日志级别', value: 'level', disabled: true, enabled: true },
  { label: '日志位置', value: 'caller_line', disabled: false, enabled: true },
  { label: '业务码', value: 'error_code', disabled: true, enabled: true },
  {
    label: '业务码信息',
    value: 'error_msg',
    disabled: true,
    enabled: true,
  },
  { label: '日志消息', value: 'msg', disabled: false, enabled: true },
  {
    label: '堆栈信息',
    value: 'stack',
    disabled: false,
    enabled: true,
  },
  { label: '扩展信息', value: 'extend', disabled: false, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: false },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});
const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<SystemLog[]>([]);
const tableDataTotal = ref<number>(0);

onBeforeMount(() => {
  defaultQuery();
  fetchSystemLogList();
});

// 默认请求参数
const defaultQuery = () => {
  listQuery.value.trace_id = route.query.trace_id;
};

// 获取网络请求日志列表
const fetchSystemLogList = async () => {
  try {
    const resp = (await getSystemLogList(listQuery.value))
      .data as SystemLogListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

const handleShowExtend = (v: string) => {
  state.extend.key = new Date().getMilliseconds();
  if (v) {
    state.extend.data = JSON.stringify(JSON.parse(v), null, 2);
  }
  state.extend.visible = true;
};

const handleShowStack = (v: string) => {
  state.extend.key = new Date().getMilliseconds();
  state.extend.data = v;
  state.extend.visible = true;
};
</script>

<style scoped lang="scss">
.filter {
  .filter-name {
    width: 180px;
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
