<script setup lang="ts">
import { TransitionPresets, useTransition } from '@/utils/use-transition'

const baseNumber = ref(0)
const duration = 2500

function toggle() {
  baseNumber.value += 100
}
const value = useTransition(baseNumber, {
  duration,
  // 动画效果库
  transition: TransitionPresets.easeOutExpo,
})
</script>

<template>
  <div>
    <span>
      {{ value.toFixed(2) }}
    </span>
    <el-button class="ml-1" type="primary" @click="toggle">
      开始
    </el-button>
  </div>
</template>
