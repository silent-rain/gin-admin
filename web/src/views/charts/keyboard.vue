<script setup>
import Chart from './components/Keyboard.vue'
</script>

<template>
  <div class="chart-container">
    <Chart height="100%" width="100%" />
  </div>
</template>

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}
</style>
