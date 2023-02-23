<template>
  <el-drawer
    :model-value="modelValue"
    title="详情"
    :before-close="handleClose"
    style="width: 1200px"
    :size="size"
  >
    <div class="details">
      <highlightjs class="details" key:="key" :language="language"
      :code="getData" />
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    modelValue: boolean;
    data: string;
    language: string;
    key?: number;
    size?: string;
  }>(),
  {
    key: 0,
    size: '30%',
  },
);

const hljs = ref(null);

const emit = defineEmits(['update:modelValue']);

const getData = computed(() => {
  return props.data;
});

const handleClose = (done: () => void) => {
  emit('update:modelValue', false);
  done();
};
</script>

<style scoped lang="scss">
.details {
  height: 80vh;
}
</style>
