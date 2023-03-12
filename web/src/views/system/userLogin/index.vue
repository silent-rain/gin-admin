<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:userlogin:list')" class="filter">
      <el-input
        v-model="listQuery.nickname"
        class="filter-name"
        :disabled="isDisabledButton('sys:userlogin:list')"
        placeholder="请输入用户昵称"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.remote_addr"
        class="filter-name"
        :disabled="isDisabledButton('sys:userlogin:list')"
        placeholder="请输入登录IP"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button
          type="primary"
          :icon="Search"
          :disabled="isDisabledButton('sys:userlogin:list')"
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
          :screen-full-element="'el-table-full'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchUserLoginList"
        />
      </div>
    </div>

    <el-table
      class="el-table-full"
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
    >
      <el-table-column
        v-if="checkedDict.id"
        prop="id"
        label="自增ID"
        width="80"
      />
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
        width="150"
      />
      <el-table-column
        v-if="checkedDict.remote_addr"
        prop="remote_addr"
        label="登录IP"
        show-overflow-tooltip
        width="150"
      />
      <el-table-column
        v-if="checkedDict.user_agent"
        prop="user_agent"
        label="用户代理"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.status"
        prop="status"
        label="登录状态"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            :disabled="isDisabledButton('sys:userlogin:status')"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
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
    </el-table>
    <Pagination
      v-model:currentPage="listQuery.page"
      v-model:pageSize="listQuery.page_size"
      :total="tableDataTotal"
      @pagination="fetchUserLoginList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { Search, Delete } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getUserLoginList,
  updateUserLoginStatus,
} from '@/api/system/user-login';
import { UserLoginListRsp, UserLogin } from '~/api/system/user-login';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = ref<any>({
  page: 1,
  page_size: 10,
  nickname: '',
  remote_addr: '',
});
// 过滤事件
const handleFilter = () => {
  fetchUserLoginList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {};
};

const checkAllList = [
  { label: '自增ID', value: 'id', disabled: false, enabled: false },
  { label: '用户ID', value: 'user_id', disabled: false, enabled: false },
  { label: '用户昵称', value: 'nickname', disabled: true, enabled: true },
  { label: '登录IP', value: 'remote_addr', disabled: false, enabled: true },
  { label: '用户代理', value: 'user_agent', disabled: false, enabled: true },
  { label: '登录状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: true },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<UserLogin[]>();
const tableDataTotal = ref<number>(0);

onBeforeMount(() => {
  fetchUserLoginList();
});

// 获取用户登录信息列表
const fetchUserLoginList = async () => {
  try {
    const resp = (await getUserLoginList(listQuery.value))
      .data as UserLoginListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 状态变更
const handleStatusChange = async (row: UserLogin) => {
  const data = {
    id: row.id,
    user_id: row.user_id,
    status: row.status,
  };
  try {
    await updateUserLoginStatus(data);
    fetchUserLoginList();
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
