<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:weblog:list')" class="filter">
      <el-input
        v-model="listQuery.nickname"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
        placeholder="请输入用户昵称"
        @keyup.enter.native="handleFilter"
      />
      <!-- <el-select
        v-model="listQuery.os_type"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
        placeholder="请选择终端类型"
        @change="handleFilter"
      >
        <el-option
          v-for="(item, _) in osTypeOptions"
          :key="item.value"
          :label="item.lebal"
          :value="item.value"
        />
      </el-select> -->
      <el-select
        v-model="listQuery.error_type"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
        placeholder="请选择错误类型"
        @change="handleFilter"
      >
        <el-option
          v-for="(item, _) in errorTypeOptions"
          :key="item.value"
          :label="item.lebal"
          :value="item.value"
        />
      </el-select>
      <el-select
        v-model="listQuery.level"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
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
      <el-input
        v-model="listQuery.url"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
        placeholder="请输入页面链接"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.msg"
        class="filter-name"
        :disabled="isDisabledButton('sys:weblog:list')"
        placeholder="请输入日志消息"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:weblog:list')"
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
          @refreshEvent="fetchWebLogList"
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
        v-if="checkedDict.os_type"
        prop="os_type"
        label="终端类型"
        show-overflow-tooltip
        width="100"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-for="(item, _) in osTypeOptions.filter(
              (v) => v.value === scope.row.os_type,
            )"
            size="small"
            :key="item.value"
            :type="item.type"
          >
            {{ item.lebal }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.error_type"
        prop="error_type"
        label="错误类型"
        show-overflow-tooltip
        width="100"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-for="(item, _) in errorTypeOptions.filter(
              (v) => v.value === scope.row.error_type,
            )"
            size="small"
            :key="item.value"
            :type="item.type"
          >
            {{ item.lebal }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.level"
        prop="level"
        label="日志级别"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.caller_line"
        prop="caller_line"
        label="日志位置"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.url"
        prop="url"
        label="错误页面"
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
      @pagination="fetchWebLogList"
    />

    <!-- 堆栈信息 -->
    <LogDrawer
      v-if="state.stack.visible"
      v-model="state.stack.visible"
      :data="state.stack.data"
      :key="state.stack.key"
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
import { getWebList } from '@/api/system/log';
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
  nickname: null,
  os_type: null,
  error_type: null,
  level: null,
  url: null,
  msg: null,
});

const state = reactive({
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
// 终端类型: 0: 未知,1: 安卓,2 :ios,3 :web
const osTypeOptions = [
  {
    lebal: '未知',
    value: 0,
    type: 'info',
  },
  {
    lebal: '安卓',
    value: 1,
    type: '',
  },
  {
    lebal: 'IOS',
    value: 2,
    type: '',
  },
  {
    lebal: 'WEB',
    value: 3,
    type: '',
  },
];
// 错误类型: 1:接口报错,2:代码报错
const errorTypeOptions = [
  {
    lebal: '代码报错',
    value: 2,
    type: 'success',
  },
  {
    lebal: '接口报错',
    value: 1,
    type: '',
  },
];

// 过滤事件
const handleFilter = () => {
  router.push({
    path: route.path,
    query: listQuery.value,
  });
  fetchWebLogList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {} as any;
};

const checkAllList = [
  { label: '日志ID', value: 'id', disabled: false, enabled: false },
  { label: '用户ID', value: 'user_id', disabled: false, enabled: false },
  { label: '用户昵称', value: 'nickname', disabled: true, enabled: true },
  { label: 'Trace ID', value: 'trace_id', disabled: false, enabled: false },
  { label: '终端类型', value: 'os_type', disabled: true, enabled: true },
  { label: '错误类型', value: 'error_type', disabled: true, enabled: true },
  { label: '日志级别', value: 'level', disabled: true, enabled: true },
  { label: '日志位置', value: 'caller_line', disabled: false, enabled: true },
  { label: '错误页面', value: 'url', disabled: true, enabled: true },
  { label: '日志消息', value: 'msg', disabled: false, enabled: true },
  {
    label: '堆栈信息',
    value: 'stack',
    disabled: false,
    enabled: true,
  },
  { label: '备注', value: 'note', disabled: false, enabled: false },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});
const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<SystemLog[]>([]);
const tableDataTotal = ref<number>(0);

onBeforeMount(() => {
  defaultQuery();
  fetchWebLogList();
});

// 默认请求参数
const defaultQuery = () => {
  listQuery.value.trace_id = route.query.trace_id;
};

// 获取 WEB 日志列表
const fetchWebLogList = async () => {
  try {
    const resp = (await getWebList(listQuery.value)).data as SystemLogListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 显示堆栈信息
const handleShowStack = (v: string) => {
  state.stack.key = new Date().getMilliseconds();
  state.stack.data = v;
  state.stack.visible = true;
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
