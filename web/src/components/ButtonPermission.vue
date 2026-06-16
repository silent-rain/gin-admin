<script setup lang="ts">
import { hasButtonPermission, isDisabledButton } from '@/hooks/use-permission'

const props = withDefaults(
  defineProps<{
    permission: string
    icon?: any
    type?: string
    link?: boolean
    size?: string
  }>(),
  {
    link: false,
  },
)
const emit = defineEmits(['click'])

function handleClick() {
  emit('click')
}
</script>

<template>
  <el-button
    v-if="hasButtonPermission(props.permission)"
    :type="props.type"
    :icon="props.icon"
    :link="props.link"
    :size="props.size"
    :disabled="isDisabledButton(props.permission)"
    @click="handleClick"
  >
    <slot />
  </el-button>
</template>

<style scoped lang="scss"></style>
