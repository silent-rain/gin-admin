<template>
  <el-card>
    <!-- 过滤条件 -->
    <div v-if="hasButtonPermission('sys:user:list')" class="filter">
      <el-input
        v-model="listQuery.nickname"
        class="filter-name"
        :disabled="isDisabledButton('sys:user:list')"
        placeholder="筛选用户昵称"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.phone"
        class="filter-name"
        :disabled="isDisabledButton('sys:user:list')"
        placeholder="筛选手机号码"
        @keyup.enter.native="handleFilter"
      />
      <el-input
        v-model="listQuery.email"
        class="filter-name"
        :disabled="isDisabledButton('sys:user:list')"
        placeholder="筛选邮箱"
        @keyup.enter.native="handleFilter"
      />
      <el-button-group>
        <ButtonPermission
          permission="sys:user:list"
          type="primary"
          :icon="Search"
          @click="handleFilter"
        >
          查询
        </ButtonPermission>
        <el-button type="primary" :icon="Delete" @click="handleCleanFilter" />
      </el-button-group>
    </div>

    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button">
        <ButtonPermission
          permission="sys:user:add"
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
              permission="sys:user:delall"
              type="danger"
              :icon="Delete"
            >
              批量删除
            </ButtonPermission>
          </template>
        </el-popconfirm>

        <ButtonPermission
          permission="sys:user:import"
          type=""
          @click="handleImportEvent"
        >
          导入
        </ButtonPermission>
        <ButtonPermission
          permission="sys:user:export"
          type=""
          @click="handleExportEvent"
        >
          导出
        </ButtonPermission>
      </div>
      <div class="right-button">
        <ConvenienTools
          v-model:size="tableSize"
          v-model:checkedDict="checkedDict"
          :screen-full-element="'el-table-user'"
          :check-all-list="checkAllList"
          @refreshEvent="fetchUserList"
        />
      </div>
    </div>

    <!-- 添加/编辑表单 -->
    <UserForm
      v-if="state.userForm.visible"
      v-model:data="state.userForm.data"
      v-model:visible="state.userForm.visible"
      :type="state.userForm.type"
      :width="state.userForm.width"
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
        prop="id"
        label="用户ID"
        width="80"
      />
      <el-table-column
        v-if="checkedDict.realname"
        prop="realname"
        label="姓名"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.nickname"
        prop="nickname"
        label="昵称"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.gender"
        prop="gender"
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
        prop="age"
        label="年龄"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.birthday"
        prop="birthday"
        label="出生日期"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.avatar"
        prop="avatar"
        label="头像"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.phone"
        prop="phone"
        label="手机号码"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.email"
        prop="email"
        label="邮箱"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.intro"
        prop="intro"
        label="介绍"
        show-overflow-tooltip
      />
      <el-table-column
        v-if="checkedDict.sort"
        prop="sort"
        label="排序"
        show-overflow-tooltip
        width="80"
      />
      <el-table-column
        v-if="checkedDict.roles"
        prop="roles"
        label="角色"
        align="center"
      >
        <template #default="scope">
          <el-tag
            v-for="(item, _) in scope.row.roles"
            :key="item.id"
            size="small"
          >
            {{ item.name }}
          </el-tag>
        </template>
      </el-table-column>
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
            :disabled="isDisabledButton('sys:user:status')"
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
        width="186"
      >
        <template #default="scope">
          <ButtonPermission
            permission="sys:user:update"
            type="primary"
            size="small"
            :icon="EditPen"
            link
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
                permission="sys:user:delete"
                type="danger"
                size="small"
                :icon="Delete"
                link
              >
                删除
              </ButtonPermission>
            </template>
          </el-popconfirm>

          <el-popconfirm
            confirm-button-text="确认"
            cancel-button-text="取消"
            :icon="InfoFilled"
            icon-color="#E6A23C"
            title="确定重置吗?"
            @confirm="handleResetUserPwd(scope.row.id)"
            @cancel="handleCancelEvent"
          >
            <template #reference>
              <ButtonPermission
                permission="sys:user:resetPwd"
                type="warning"
                size="small"
                :icon="Finished"
                link
              >
                重置密码
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
      @pagination="fetchUserList"
    />
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import {
  EditPen,
  Search,
  Finished,
  Delete,
  Plus,
  InfoFilled,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { useBasicStore } from '@/store/basic';
import {
  getUserList,
  updateUserStatus,
  deleteUser,
  batchDeleteUser,
  resetUserPwd,
} from '@/api/permission/user';
import { UserListRsp, User } from '@/typings/api/permission/user';
import Pagination from '@/components/Pagination.vue';
import ConvenienTools from '@/components/ConvenienTools/index.vue';
import ButtonPermission from '@/components/ButtonPermission.vue';
import UserForm from './components/UserForm.vue';
import { aoaToSheetXlsx } from '@/utils/excel';
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission';

const basicStore = useBasicStore();

// 筛选过滤条件
const listQuery = reactive({
  nickname: '',
  phone: '',
  email: '',
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
  listQuery.phone = '';
  listQuery.email = '';
};

const state = reactive({
  userForm: {
    data: {} as User,
    visible: false,
    type: '',
    width: '750px',
  },
});

const checkAllList = [
  { label: '用户ID', value: 'id', disabled: false, enabled: true },
  { label: '姓名', value: 'realname', disabled: false, enabled: true },
  { label: '昵称', value: 'nickname', disabled: true, enabled: true },
  { label: '性别', value: 'gender', disabled: false, enabled: true },
  { label: '年龄', value: 'age', disabled: false, enabled: false },
  { label: '出生日期', value: 'birthday', disabled: false, enabled: false },
  { label: '头像', value: 'avatar', disabled: false, enabled: false },
  { label: '手机号码', value: 'phone', disabled: false, enabled: true },
  { label: '邮箱', value: 'email', disabled: false, enabled: true },
  { label: '介绍', value: 'intro', disabled: false, enabled: false },
  { label: '备注', value: 'note', disabled: false, enabled: true },
  { label: '角色', value: 'roles', disabled: false, enabled: true },
  { label: '排序', value: 'sort', disabled: false, enabled: false },
  { label: '状态', value: 'status', disabled: true, enabled: true },
  { label: '创建时间', value: 'created_at', disabled: false, enabled: false },
  { label: '更新时间', value: 'updated_at', disabled: false, enabled: true },
  { label: '操作', value: 'operation', disabled: false, enabled: true },
];
const checkedDict = ref<any>({});

const tableSize = ref<string>(basicStore.settings.defaultSize);
const tableData = ref<User[]>();
const tableDataTotal = ref<number>(0);
const multipleSelection = ref<User[]>([]);

onBeforeMount(() => {
  fetchUserList();
});

// 获取用户列表
const fetchUserList = async () => {
  try {
    const resp = (await getUserList(listQuery)).data as UserListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.tatol;
  } catch (error) {
    console.log(error);
  }
};

// 删除
const handleDelete = async (row: User) => {
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
const handleEdit = async (row: User) => {
  state.userForm.data = { ...row };
  state.userForm.type = 'edit';
  state.userForm.visible = true;
};
// 添加
const handleAdd = async () => {
  state.userForm.data.age = 0;
  state.userForm.data.sort = 1;
  state.userForm.data.gender = 0;
  state.userForm.data.avatar = '';
  state.userForm.data.password = basicStore.settings.defaultPassword;
  state.userForm.data.role_ids = [];
  state.userForm.type = 'add';
  state.userForm.visible = true;
};
// 多选事件
const handleSelectionChange = (val: User[]) => {
  multipleSelection.value = val;
};

// 批量删除
const handleBatchDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据');
    return;
  }
  const data = {
    ids: multipleSelection.value.map((v: User) => {
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
// 取消批量删除事件
const handleBatchDeleteCancel = () => {
  ElMessage.warning('取消操作');
};

// 删除取消事件
const handleCancelEvent = () => {
  ElMessage.warning('取消操作');
};

// 状态变更
const handleStatusChange = async (row: User) => {
  const data = {
    id: row.id,
    status: row.status,
  };
  try {
    await updateUserStatus(data);
    fetchUserList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 重置密码
const handleResetUserPwd = async (id: number) => {
  const data = {
    id,
  };
  try {
    await resetUserPwd(data);
    fetchUserList();
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};

// 导入
const handleImportEvent = () => {
  console.log('导入');
};

// 导出
const handleExportEvent = async () => {
  const fileName = '用户列表';
  const header = [
    '用户ID',
    '姓名',
    '昵称',
    '性别',
    '年龄',
    '出生日期',
    '手机号码',
    '邮箱',
    '介绍',
    '备注',
    '角色',
    '排序',
    '状态',
    '创建时间',
    '更新时间',
  ];
  const data = tableData.value?.map((item, _index) => {
    const {
      id,
      realname,
      nickname,
      gender,
      age,
      birthday,
      phone,
      email,
      intro,
      note,
      roles,
      sort,
      status,
      created_at,
      updated_at,
    } = item;
    const genderMap = {
      0: '保密',
      1: '女',
      2: '男',
    };
    const genderZh = genderMap[gender];
    const statusZh = status === 1 ? '启用' : '禁用';
    const rolesList: string[] = [];
    roles.forEach((v) => {
      rolesList.push(v.name);
    });

    return [
      id,
      realname,
      nickname,
      genderZh,
      age,
      birthday,
      phone,
      email,
      intro,
      note,
      rolesList.join(';'),
      sort,
      statusZh,
      created_at,
      updated_at,
    ];
  });
  aoaToSheetXlsx({
    data,
    header,
    filename: `${unref(fileName)}.xlsx`,
  });
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
