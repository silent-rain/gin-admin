<template>
  <el-card id="role-manage">
    <!-- 过滤条件 -->
    <div class="filter">
      <label>角色名称: </label>
      <el-input
        v-model="queryList.name"
        class="filter-name"
        placeholder="请输入角色名称"
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
        <el-button type="primary" :icon="Plus" @click="handleFilter"
          >添加
        </el-button>
        <el-button type="danger" :icon="Delete" @click="handleFilter"
          >删除
        </el-button>
      </div>
      <div class="right-button">
        <el-tooltip content="刷新" placement="top">
          <el-button :icon="RefreshRight" @click="fetchRoleList" />
        </el-tooltip>
        <el-tooltip content="密度" placement="top">
          <el-dropdown @command="handleTableSizeCommand">
            <el-button
              class="el-dropdown-link"
              :icon="Expand"
              trigger="click"
            />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  v-for="(item, _) in tableSizeOptions"
                  :command="item.value"
                  >{{ item.label }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tooltip>

        <el-popover placement="bottom" :width="180" trigger="hover">
          <template #reference>
            <span>
              <el-tooltip content="设置" placement="top">
                <el-button :icon="Setting" @click="handleFilter" />
              </el-tooltip>
            </span>
          </template>
          <div
            class="operation-settings-show"
            style="
              display: flex;
              justify-content: space-between;
              align-items: center;
            "
          >
            <el-checkbox
              v-model="checkAll"
              :indeterminate="isIndeterminate"
              @change="handleCheckAllChange"
              >列展示
            </el-checkbox>
            <el-button type="primary" text>重置</el-button>
          </div>

          <el-divider style="margin: 4px 0" />

          <el-checkbox-group v-model="tableColsCheckList">
            <el-checkbox label="Option A" />
            <el-checkbox label="Option B" />
            <el-checkbox label="Option C" />
            <el-checkbox label="disabled" disabled />
            <el-checkbox label="selected and disabled" disabled />
          </el-checkbox-group>
        </el-popover>
        <el-tooltip content="全屏" placement="top">
          <el-button :icon="FullScreen" @click="handleScreenFull" />
        </el-tooltip>
      </div>
    </div>

    <el-table :data="tableData" style="width: 100%" :size="tableSize">
      <el-table-column type="selection" width="55" />
      <el-table-column label="Date" width="120">
        <template #default="scope">{{ scope.row.date }}</template>
      </el-table-column>
      <el-table-column property="name" label="Name" width="120" />
      <el-table-column
        property="address"
        label="Address"
        show-overflow-tooltip
      />
    </el-table>
  </el-card>
</template>
<script setup lang="ts">
import {
  Search,
  Delete,
  Plus,
  Setting,
  RefreshRight,
  Expand,
  FullScreen,
} from '@element-plus/icons-vue';
import { reactive, ref, onBeforeMount } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { useBasicStore } from '@/store/basic';
import { getRoleList } from '@/api/system/role';

const { settings } = storeToRefs(useBasicStore());

// 筛选过滤条件
const queryList = reactive({
  name: '',
});
// 过滤事件
const handleFilter = () => {
  console.log(queryList);
};
// 清空过滤条件
const handleCleanFilter = () => {
  queryList.name = '';
};

// 表格尺寸列表
const tableSizeOptions = [
  {
    label: '宽松',
    value: 'large',
  },
  {
    label: '默认',
    value: 'default',
  },
  {
    label: '紧凑',
    value: 'small',
  },
];
// 表格尺寸
const tableSize = ref(settings.value.defaultSize);
// 表格尺寸选择事件
const handleTableSizeCommand = (data: string) => {
  tableSize.value = data;
};

const checkAll = ref(false);
const isIndeterminate = ref(true);
const tableColsCheckList = ref(['selected and disabled', 'Option A']);
const handleCheckAllChange = (val: boolean) => {
  // checkedCities.value = val ? cities : [];
  isIndeterminate.value = false;
};

// 全屏
const screenFullFlag = ref(false);
const handleScreenFull = () => {
  let element = document.getElementById('role-manage');
  if (!element) {
    return;
  }
  // 不全屏是null,返回false,
  screenFullFlag.value = document.fullscreenElement === null ? false : true;
  // false是进入全屏状态
  if (screenFullFlag.value) {
    // 退出全屏
    if (document.exitFullscreen) {
      document.exitFullscreen();
    }
  } else {
    // 全屏
    element.requestFullscreen();
  }
  // 切换文本状态（只是用在文本上，文本不是动态可以忽略）
  screenFullFlag.value = !screenFullFlag.value;
};

const tableData = ref([]);

onBeforeMount(() => {
  fetchRoleList();
});

const fetchRoleList = async () => {
  const resp = (await getRoleList(queryList)).data;
  console.log(resp);
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

:deep(.operation-settings-show) {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
