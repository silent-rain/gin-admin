<template>
  <el-card>
    <!-- 提示 -->
    <div class="tips">
      <p><label>提示:</label></p>
      <p>Token 令牌可用于 API 接口访问</p>
      <p>最多可申请 {{ state.api_auth_max_token_num }} 个令牌</p>
    </div>
    <!-- 表格全局按钮 -->
    <div class="operation-button">
      <div class="left-button">
        <el-button
          type="primary"
          :icon="Plus"
          :disabled="tableDataTotal >= state.api_auth_max_token_num"
          @click="handleAdd"
        >
          添加
        </el-button>
      </div>
      <div class="right-button"></div>
    </div>

    <!-- 添加/编辑表单 -->
    <ApiTokenForm
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
    >
      <el-table-column prop="token" label="令牌" show-overflow-tooltip />
      <el-table-column prop="passphrase" label="口令" show-overflow-tooltip />
      <el-table-column
        prop="permission"
        label="权限标识"
        show-overflow-tooltip
      />
      <el-table-column prop="status" label="状态" align="center" width="90">
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :active-value="1"
            :inactive-value="0"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" align="center" width="120">
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            :icon="EditPen"
            @click="handleEdit(scope.row)"
          >
            修改
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
              <el-button link type="danger" size="small" :icon="Delete">
                删除
              </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup lang="ts">
import { reactive, ref, onBeforeMount } from 'vue';
import { EditPen, Delete, InfoFilled, Plus } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { useUserStore } from '@/store/user';
import {
  getUserApiTokenList,
  updateUserApiTokenStatus,
  deleteUserApiToken,
} from '@/api/permission/user-api-token';
import {
  UserApiTokenListRsp,
  UserApiToken,
} from '~/api/permission/user-api-token';
import { getConfigInfo } from '@/api/data-center/config';
import { ConfigRsp } from '~/api/data-center/config';
import ApiTokenForm from './ApiTokenForm.vue';

const userStore = useUserStore();

const state = reactive({
  form: {
    data: {} as UserApiToken,
    visible: false,
    type: '',
  },
  api_auth_max_token_num: 5,
});

const tableData = ref<UserApiToken[]>();
const tableDataTotal = ref<number>(0);

onBeforeMount(() => {
  fetchUserApiTokenList();
  fetchConfigInfo();
});

// 获取令牌列表
const fetchUserApiTokenList = async () => {
  if (!userStore.userId) {
    return;
  }
  const data = {
    page: 1,
    page_size: 20,
    user_id: userStore.userId,
    nickname: null,
    status: null,
  };
  try {
    const resp = (await getUserApiTokenList(data)).data as UserApiTokenListRsp;
    tableData.value = resp.data_list;
    tableDataTotal.value = resp.data_list.length;
  } catch (error) {
    console.log(error);
  }
};

// 通过 key 获取可申请最大令牌数
const fetchConfigInfo = async () => {
  try {
    const resp = (await getConfigInfo({
      key: 'api_auth_max_token_num',
    })) as ConfigRsp;
    state.api_auth_max_token_num = resp.data.value;
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
  state.form.data.user_id = userStore.userId;
};
// 添加
const handleAdd = async () => {
  state.form.type = 'add';
  state.form.visible = true;
  state.form.data.status = 1;
  state.form.data.user_id = userStore.userId;
  state.form.data.permission = '';
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
.tips {
  border: 1px #dcdfe6 dashed;

  label {
    color: orange;
  }
  p {
    color: #9ea0a6;
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
