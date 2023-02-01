<template>
  <div class="convenient-tools">
    <el-tooltip content="刷新" placement="top">
      <el-button :icon="RefreshRight" @click="refreshEvent" />
    </el-tooltip>
    <el-tooltip content="密度" placement="top">
      <el-dropdown @command="handleSizeCommand">
        <el-button class="el-dropdown-link" :icon="Expand" trigger="click" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item
              v-for="(item, _) in sizeOptions"
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
            <el-button :icon="Setting" />
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
        <el-button type="primary" text @click="handleCheckdReset"
          >重置</el-button
        >
      </div>

      <el-divider style="margin: 4px 0" />

      <el-checkbox-group v-model="checkedList" @change="handleCheckedChange">
        <el-checkbox
          v-for="item in checkAllList"
          :key="item.value"
          :label="item.value"
          :disabled="item.disabled"
          style="width: 100%"
          >{{ item.value }}
        </el-checkbox>
      </el-checkbox-group>
    </el-popover>
    <el-tooltip content="全屏" placement="top">
      <el-button :icon="FullScreen" @click="handleScreenFull" />
    </el-tooltip>
  </div>
</template>

<script setup lang="ts">
import {
  Setting,
  RefreshRight,
  Expand,
  FullScreen,
} from '@element-plus/icons-vue';
import defaultSettings from '@/settings';

const props = withDefaults(
  defineProps<{
    size?: string; // UI 尺寸
    screenFullElement?: string; // 全屏元素 class-name
    checkAllList?: any[];
    checkedDict?: any;
  }>(),
  {
    size: defaultSettings.defaultSize,
    screenFullElement: '',
    checkAllList: () => [],
    checkedDict: {},
  },
);

const emit = defineEmits(['refreshEvent', 'update:size', 'update:checkedDict']);

onBeforeMount(() => {
  checkedList.value = props.checkAllList
    .filter((v: any) => {
      if (v.enabled) {
        return v;
      }
    })
    .map((v) => v.value);
  isIndeterminate.value = !(
    checkedList.value.length === props.checkAllList.length
  );
  checkedListToMap();
});

// 刷新事件
const refreshEvent = () => {
  emit('refreshEvent', null);
};

// UI 尺寸列表
const sizeOptions = [
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
// 尺寸选择事件
const handleSizeCommand = (value: string) => {
  emit('update:size', value);
};

// 多选框组
const checkAll = ref(true);
const isIndeterminate = ref(true);
const checkedList = ref<string[]>([]);
// 全选
const handleCheckAllChange = (val: boolean) => {
  checkedList.value = val
    ? props.checkAllList.map((v) => v.value)
    : props.checkAllList
        .filter((v: any) => {
          if (v.disabled) {
            return v;
          }
        })
        .map((v) => v.value);

  const checkedCount = checkedList.value.length;
  isIndeterminate.value =
    checkedCount > 0 && checkedCount < props.checkAllList.length;
  checkedListToMap();
};
// 重置
const handleCheckdReset = () => {
  checkAll.value = true;
  checkedList.value = props.checkAllList.map((v) => v.value);

  const checkedCount = checkedList.value.length;
  isIndeterminate.value =
    checkedCount > 0 && checkedCount < props.checkAllList.length;
};
// 筛选框点击事件
const handleCheckedChange = (value: string[]) => {
  const checkedCount = value.length;
  checkAll.value = checkedCount === props.checkAllList.length;
  isIndeterminate.value =
    checkedCount > 0 && checkedCount < props.checkAllList.length;
  checkedListToMap();
};
// 多选列表转map
const checkedListToMap = () => {
  const m = {};
  checkedList.value.forEach((v) => {
    m[v] = true;
  });
  emit('update:checkedDict', m);
};

// 全屏
const screenFullFlag = ref(false);
const handleScreenFull = () => {
  let element = document.getElementsByClassName(
    props.screenFullElement as string,
  )[0];
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
:deep(.operation-settings-show) {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
