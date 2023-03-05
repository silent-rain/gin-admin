<template>
  <!-- 自定义 Icon -->
  <el-tab-pane label="自定义" name="custom">
    <el-row>
      <el-col :span="4" v-for="(item, _) in getOptions">
        <el-dropdown-item
          :command="item.value"
          :class="`custom-icon-${item.value}-item`"
        >
          <svg-icon :icon-class="item.value" class="custom-icon" />
        </el-dropdown-item>
      </el-col>
    </el-row>
  </el-tab-pane>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    query: string;
  }>(),
  {},
);

const options = [
  // common
  { label: '404', value: '404' },
  { label: 'bug', value: 'bug' },
  { label: 'chart', value: 'chart' },
  { label: 'clipboard', value: 'clipboard' },
  { label: 'component', value: 'component' },
  { label: 'dashboard', value: 'dashboard' },
  { label: 'demo', value: 'demo' },
  { label: 'documentation', value: 'documentation' },
  { label: 'drag', value: 'drag' },
  { label: 'edit', value: 'edit' },
  { label: 'education', value: 'education' },
  { label: 'email', value: 'email' },
  { label: 'example', value: 'example' },
  { label: 'excel', value: 'excel' },
  { label: 'exit-fullscreen', value: 'exit-fullscreen' },
  { label: 'eye-open', value: 'eye-open' },
  { label: 'eye', value: 'eye' },
  { label: 'form', value: 'form' },
  { label: 'fullscreen', value: 'fullscreen' },
  { label: 'guide', value: 'guide' },
  { label: 'hamburger', value: 'hamburger' },
  { label: 'icon', value: 'icon' },
  { label: 'international', value: 'international' },
  { label: 'language', value: 'language' },
  { label: 'link', value: 'link' },
  { label: 'list', value: 'list' },
  { label: 'lock', value: 'lock' },
  { label: 'message', value: 'message' },
  { label: 'money', value: 'money' },
  { label: 'nested', value: 'nested' },
  { label: 'password', value: 'password' },
  { label: 'pdf', value: 'pdf' },
  { label: 'peoples', value: 'peoples' },
  { label: 'people', value: 'people' },
  { label: 'qq', value: 'qq' },
  { label: 'search', value: 'search' },
  { label: 'shopping', value: 'shopping' },
  { label: 'sidebar-logo', value: 'sidebar-logo' },
  { label: 'size', value: 'size' },
  { label: 'skill', value: 'skill' },
  { label: 'star', value: 'star' },
  { label: 'table', value: 'table' },
  { label: 'tab', value: 'tab' },
  { label: 'theme', value: 'theme' },
  { label: 'tree', value: 'tree' },
  { label: 'tree-table', value: 'tree-table' },
  { label: 'user', value: 'user' },
  { label: 'wechat', value: 'wechat' },
  { label: 'zip', value: 'zip' },
  // nav-bar
  { label: 'dashboard', value: 'dashboard' },
  { label: 'example', value: 'example' },
  { label: 'eye-open', value: 'eye-open' },
  { label: 'eye', value: 'eye' },
  { label: 'form', value: 'form' },
  { label: 'language', value: 'language' },
  { label: 'link', value: 'link' },
  { label: 'nested', value: 'nested' },
  { label: 'password', value: 'password' },
  { label: 'table', value: 'table' },
  { label: 'theme-icon', value: 'theme-icon' },
  { label: 'tree', value: 'tree' },
  { label: 'user', value: 'user' },
];

const getOptions = computed(() => {
  if (props.query === '') {
    return options;
  }
  const re = new RegExp(`${props.query}`,"i");
  return options.filter((v) => re.test(v.value));
});
</script>

<style scoped lang="scss">
.custom-icon {
  height: 45px;
  width: 45px;
  text-align: center;
  font-size: 25px;
  &:hover {
    transform: scale(1.2);
  }
}
</style>
