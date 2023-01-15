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
          <el-button :icon="RefreshRight" @click="handleFilter" />
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

        <el-button :icon="Setting" @click="handleFilter" />
        <el-tooltip content="全屏" placement="top">
          <el-button :icon="FullScreen" @click="handleScreenFull" />
        </el-tooltip>
      </div>
    </div>
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
import { reactive, ref } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { useBasicStore } from '@/store/basic';

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

  .el-button + .el-button {
    margin-left: 0px;
  }
}
</style>
