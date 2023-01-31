<template>
  <el-pagination
    :class="{
      hidden: props.hidden,
      paginationContainer: true,
    }"
    v-model:current-page="props.currentPage"
    v-model:page-size="props.pageSize"
    :page-sizes="props.pageSizes"
    :small="props.small"
    :background="props.background"
    :disabled="props.disabled"
    :layout="props.layout"
    :total="props.total"
    :hide-on-single-page="props.hideOnSinglePage"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = withDefaults(
  defineProps<{
    currentPage: number;
    total: number;
    pageSize?: number;
    pageSizes?: number[];
    small?: boolean;
    background?: boolean;
    disabled?: boolean;
    layout?: string;
    autoScroll?: boolean;
    hidden?: boolean;
    hideOnSinglePage?: boolean;
  }>(),
  {
    currentPage: 1,
    pageSize: 10,
    pageSizes: () => {
      return [10, 20, 50, 100, 200, 300, 400];
    },
    small: false,
    background: true,
    disabled: false,
    layout: 'total, sizes, prev, pager, next, jumper',
    autoScroll: false,
    hidden: false,
    hideOnSinglePage: true,
  },
);

const emit = defineEmits([
  'update:currentPage',
  'update:pageSize',
  'pagination',
]);

const handleSizeChange = (_val: number) => {
  emit('pagination', null);
  if (ref(props.autoScroll).value) {
    scrollTo(0, 800);
  }
};
const handleCurrentChange = (_val: number) => {
  emit('pagination', null);
  if (ref(props.autoScroll).value) {
    scrollTo(0, 800);
  }
};
</script>

<style scoped lang="scss">
.paginationContainer {
  background: #fff;
  padding: 16px 16px;
  float: right;
  //   text-align: right;
}
.pagination-container.hidden {
  display: none;
}
</style>
