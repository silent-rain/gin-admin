<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    modelValue: boolean
    data: string
    language: string
    key?: number
    size?: string
  }>(),
  {
    key: 0,
    size: '30%',
  },
)

const emit = defineEmits(['update:modelValue'])

const hljs = ref(null)

const getData = computed(() => {
  return props.data
})

function handleClose(done: () => void) {
  emit('update:modelValue', false)
  done()
}
</script>

<template>
  <el-drawer
    :model-value="modelValue"
    title="详情"
    :before-close="handleClose"
    style="width: 1200px"
    :size="size"
  >
    <div class="details">
      <highlightjs
        :key="key"
        class="details"
        :language="language"
        :code="getData"
      />
    </div>
  </el-drawer>
</template>

<style scoped lang="scss">
.details {
  height: 80vh;
}
</style>
