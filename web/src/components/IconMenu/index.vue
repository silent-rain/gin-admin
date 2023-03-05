<template>
  <!-- Icon 菜单 -->
  <el-dropdown split-button type="" trigger="click" @command="handleCommand">
    <span class="el-dropdown-link">
      <template v-if="props.el_svg_icon">
        <el-icon class="icon">
          <component :is="ElSvg[props.el_svg_icon]" />
        </el-icon>
        <span>{{ props.el_svg_icon }}</span>
      </template>
      <template v-else-if="props.icon">
        <svg-icon :icon-class="props.icon" class="icon" />
        <span>{{ props.icon }}</span>
      </template>
      <template v-else>
        <span>请选择菜单图标</span>
      </template>
    </span>
    <template #dropdown>
      <el-input
        class="filter"
        v-model="query"
        placeholder="图标搜索"
        :suffix-icon="ElSvg.Search"
      />
      <el-dropdown-menu>
        <el-tabs v-model="activeName" class="icon-menus">
          <el-scrollbar>
            <ElementIcon :query="query" />
            <CustomIcon :query="query" />
          </el-scrollbar>
        </el-tabs>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import * as ElSvg from '@element-plus/icons-vue';
import ElementIcon from './ElementIcon.vue';
import CustomIcon from './CustomIcon.vue';

const props = withDefaults(
  defineProps<{
    icon?: string;
    el_svg_icon?: string;
  }>(),
  {},
);

const emits = defineEmits(['update:icon', 'update:el_svg_icon']);

const activeName = ref('element');
const query = ref('');

const handleCommand = (command: string | number | object) => {
  if (activeName.value == 'element') {
    emits('update:el_svg_icon', command);
    emits('update:icon', '');
  } else {
    emits('update:icon', command);
    emits('update:el_svg_icon', '');
  }
};
</script>

<style scoped lang="scss">
.el-dropdown-link {
  display: flex;
  align-items: center;
}
.icon-menus {
  margin: 10px;
  width: 380px;
}
.el-scrollbar {
  height: 350px;
}
.filter {
  padding: 10px;
}
</style>
