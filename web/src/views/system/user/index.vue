<template>
  <el-card id="role-manage">
    <!-- 过滤条件 -->
    <div class="filter">
      <label>用户昵称: </label>
      <el-input
        v-model="listQuery.nickname"
        class="filter-name"
        placeholder="请输入用户昵称"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <el-button type="primary" :icon="Search" @click="handleFilter"
          >查询
        </el-button>
        <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
      </el-button-group>
    </div>

    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button">
        <ConvenienButtons
          :buttonList="['add', 'batchDelete']"
          @add-event="handleAdd"
          @batch-delete-event="handleBatchDelete"
        />
      </div>
      <div class="right-button">
        <ConvenienTools
          @refreshEvent="fetchUserList"
          v-model:size="tableSize"
          :screenFullElement="'el-table-user'"
          :checkAllList="checkAllList"
          v-model:checkedDict="checkedDict"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <RoleForm
      v-model:data="state.roleForm.data"
      v-model:visible="state.roleForm.visible"
      :type="state.roleForm.type"
      :width="state.roleForm.width"
      @refresh="fetchUserList"
    />

    <el-table
      class="el-table-user"
      :data="tableData"
      style="width: 100%; margin-top: 10px"
      :size="tableSize"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column
        v-if="checkedDict.id"
        property="id"
        label="用户ID"
        width="80"
      />
      <!-- <el-table-column
        v-if="checkedDict.realname"
        property="realname"
        label="真实姓名"
        show-overflow-tooltip
      /> -->
      <el-table-column
        v-if="checkedDict.nickname"
        property="nickname"
        label="昵称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.gender"
        property="gender"
        label="性别"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-tag v-if="scope.row.gender === 0" type="info">保密</el-tag>
          <el-tag v-else-if="scope.row.gender === 1" type="success">女</el-tag>
          <el-tag v-else>男</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.age"
        property="age"
        label="年龄"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.birthday"
        property="birthday"
        label="出生日期"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.avatar"
        property="avatar"
        label="用户头像URL"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.phone"
        property="phone"
        label="手机号码"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.email"
        property="email"
        label="邮箱"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.intro"
        property="intro"
        label="介绍"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.note"
        property="note"
        label="备注"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.sort"
        property="sort"
        label="排序"
        show-overflow-tooltip
        width="80"
      />
      <el-table-column
        v-if="checkedDict.roles"
        property="roles"
        label="角色"
        align="center"
        width="90"
      >
        <template #default="scope">
          {{ scope.row.roles }}
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.status"
        property="status"
        label="状态"
        align="center"
        width="90"
      >
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column
        v-if="checkedDict.note"
        property="note"
        label="备注"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.created_at"
        property="created_at"
        label="创建时间"
        width="165"
      />
      <el-table-column
        v-if="checkedDict.updated_at"
        property="updated_at"
        label="更新时间"
        width="165"
      />
      <el-table-column
        v-if="checkedDict.operation"
        fixed="right"
        label="操作"
        align="center"
        width="186"
      >
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            :icon="EditPen"
            @click="handleEdit(scope.row)"
            >修改
          </el-button>
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
              <el-button link type="danger" size="small" :icon="Delete"
                >删除
              </el-button>
            </template>
          </el-popconfirm>
          <el-button
            link
            type="warning"
            size="small"
            :icon="Finished"
            @click="handleResetUserPwd(scope.row.id)"
            >重置密码
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination
      v-model:currentPage="listQuery.page"
      v-model:pageSize="listQuery.page_size"
      :total="tableDataTotal"
      @pagination="fetchUserList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { useBasicStore } from '@/store/basic';
import {
  EditPen,
  Search,
  Delete,
  Finished,
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import {
  getUserList,
  updateUserStatus,
  deleteUser,
  batchDeleteUser,
  resetUserPwd,
} from '@/api/system/user';
import { RoleListRsp, Role } from '~/api/system/role';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ConvenienButtons from '@/components/ConvenienButtons/index.vue';
import RoleForm from './components/RoleForm.vue';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const listQuery = reactive({
  nickname: '',
  page: 1,
  page_size: 10,
});
// 过滤事件
const handleFilter = () => {
  fetchUserList();
};
// 清空过滤条件
const handleCleanFilter = () => {
  listQuery.nickname = '';
};

const state = reactive({
  roleForm: {
    data: {} as Role,
    visible: false,
    type: '',
    width: '500px',
  },
});

const checkAllList = [
  {
    value: 'id',
    disabled: false,
    enabled: true,
  },
  {
    value: 'realname',
    disabled: true,
    enabled: true,
  },
  {
    value: 'nickname',
    disabled: true,
    enabled: true,
  },
  {
    value: 'gender',
    disabled: false,
    enabled: true,
  },
  {
    value: 'age',
    disabled: false,
    enabled: false,
  },
  {
    value: 'birthday',
    disabled: false,
    enabled: false,
  },
  {
    value: 'avatar',
    disabled: false,
    enabled: true,
  },
  {
    value: 'phone',
    disabled: false,
    enabled: true,
  },
  {
    value: 'email',
    disabled: false,
    enabled: false,
  },
  {
    value: 'intro',
    disabled: false,
    enabled: false,
  },
  {
    value: 'note',
    disabled: false,
    enabled: true,
  },
  {
    value: 'roles',
    disabled: false,
    enabled: true,
  },
  {
    value: 'sort',
    disabled: false,
    enabled: false,
  },
  {
    value: 'status',
    disabled: true,
    enabled: true,
  },
  {
    value: 'created_at',
    disabled: false,
    enabled: false,
  },
  {
    value: 'updated_at',
    disabled: false,
    enabled: true,
  },
  {
    value: 'operation',
    disabled: false,
    enabled: true,
  },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(settings.value.defaultSize);
const tableData = ref<Role[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<Role[]>([]);

onBeforeMount(() => {
  fetchUserList();
});

// 获取用户列表
const fetchUserList = async () => {
  try {
    const resp = (await getUserList(listQuery)).data as RoleListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: Role) => {
  const data = {
    id: row.id,
  };
  try {
    await deleteUser(data);
    fetchUserList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
// 编辑
const handleEdit = async (row: Role) => {
  state.roleForm.data = row;
  state.roleForm.type = 'edit';
  state.roleForm.visible = true;
};
// 添加
const handleAdd = async () => {
  state.roleForm.data.sort = 1;
  state.roleForm.type = 'add';
  state.roleForm.visible = true;
};
// 多选事件
const handleSelectionChange = (val: Role[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: Role) => {
      return v.id;
    }),
  };
  try {
    await batchDeleteUser(data);
    fetchUserList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 删除取消事件
const handleCancelEvent = () => {
  ElMessage.warning('取消操作');
};

// 状态变更
const handleStatusChange = async (row: Role) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    const resp = (await updateUserStatus(data)).data as RoleListRsp;
    tableData.value = resp.data_list;
    fetchUserList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 重置密码
const handleResetUserPwd = async (id: number) => {
  const data = {
    id: id,
  };
  try {
    const resp = (await resetUserPwd(data)).data as RoleListRsp;
    tableData.value = resp.data_list;
    fetchUserList();
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
