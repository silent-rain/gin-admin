<script setup name="ThirdChildren">
const testInput = ref('')
</script>

<template>
  <div class="mt-20px mb-10px font-bold">
    I Am ThirdChildren.vue
  </div>
  <el-input
    v-model="testInput"
    placeholder="input test keepAlive"
    style="width: 400px"
  />
</template>

<style scoped lang="scss"></style>
