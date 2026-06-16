<script setup>
const data = ref('')
function setData() {
  data.value = '执行长按指令'
}
</script>

<template>
  <div>
    <el-button v-longpress="setData">
      长按指令
    </el-button>
    <p>{{ data }}</p>
  </div>
</template>

<style lang="scss" scoped></style>
