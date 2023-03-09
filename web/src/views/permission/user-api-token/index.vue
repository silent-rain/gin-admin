<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:role:list')" class="filter">
      <el-input
        v-model="listQuery.nickname"
        class="filter-name"
        clearable
        :disabled="isDisabledButton('sys:role:list')"
        placeholder="筛选用户昵称"
        @keyup.enter.native="handleFilter"
      />
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
          :disabled="isDisabledButton('sys:role:list')"
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
          permission="sys:role:add"
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
              permission="sys:role:delall"
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
          @refreshEvent="fetchUserApiTokenList"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <UserApiTokenForm
      v-if="state.form.visible"
      v-model:data="state.form.data"
      v-model:visible="state.form.visible"
      :type="state.form.type"
      @refresh="fetchUserApiTokenList"
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
      />
      <el-table-column
        v-if="checkedDict.permission"
        prop="permission"
        label="权限标识"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.token"
        prop="token"
        label="令牌"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.passphrase"
        prop="passphrase"
        label="口令"
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
            :disabled="isDisabledButton('sys:role:status')"
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
        fixed="right"
        label="操作"
        align="center"
        width="120"
      >
        <template #default="scope">
          <ButtonPermission
            permission="sys:role:update"
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
                permission="sys:role:delete"
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
      @pagination="fetchUserApiTokenList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import {
  EditPen,
  Search,
  Delete,
  InfoFilled,
  Plus,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getUserApiTokenList,
  updateUserApiTokenStatus,
  deleteUserApiToken,
  batchDeleteUserApiToken,
} from '@/api/permission/user-api-token';
import {
  UserApiTokenListRsp,
  UserApiToken,
} from '~/api/permission/user-api-token';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ButtonPermission from '@/components/ButtonPermission.vue';
import UserApiTokenForm from './components/UserApiTokenForm.vue';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = ref<any>({
  page: 1,
  page_size: 10,
  nickname: null,
  status: null,
});
const statusOptions = [
  {
    label: '启用',
    value: 1,
  },
  {
    label: '禁用',
    value: 0,
  },
];
// 过滤事件
const handleFilter = () => {
  fetchUserApiTokenList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.value = {};
};
// 状态变更事件
const handleChangeStatus = (value: any) => {
  if (!value) {
    listQuery.value.status = null;
  }
  handleFilter();
};

const state = reactive({
  form: {
    data: {} as UserApiToken,
    visible: false,
    type: '',
  },
});

const checkAllList = [
  { label: '自增ID', value: 'id', disabled: false, enabled: false },
  { label: '用户ID', value: 'user_id', disabled: true, enabled: true },
  { label: '用户昵称', value: 'nickname', disabled: true, enabled: true },
  { label: '权限标识', value: 'permission', disabled: true, enabled: true },
  { label: '令牌', value: 'token', disabled: true, enabled: true },
  { label: '口令', value: 'passphrase', disabled: true, enabled: true },
  { label: '备注', value: 'note', disabled: false, enabled: true },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<UserApiToken[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<UserApiToken[]>([]);

onBeforeMount(() => {
  fetchUserApiTokenList();
});

// 获取令牌列表
const fetchUserApiTokenList = async () => {
  try {
    const resp = (await getUserApiTokenList(listQuery.value))
      .data as UserApiTokenListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: UserApiToken) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteUserApiToken(data);
    fetchUserApiTokenList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: UserApiToken) => {
  state.form.data = { ...row };
  state.form.type = 'edit';
  state.form.visible = true;
};
// 添加
const handleAdd = async () => {
  state.form.type = 'add';
  state.form.visible = true;
  state.form.data.status = 1;
};
// 多选事件
const handleSelectionChange = (val: UserApiToken[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: UserApiToken) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteUserApiToken(data);
    fetchUserApiTokenList();
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
const handleStatusChange = async (row: UserApiToken) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    await updateUserApiTokenStatus(data);
    fetchUserApiTokenList();
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
