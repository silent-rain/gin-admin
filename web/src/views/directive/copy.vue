<script setup>
const input = ref('')
</script>

<template>
  <div class="box">
    <el-input
      v-model="input"
      style="width: 300px"
      placeholder="输入需要粘贴的值"
    />
    <el-button v-copy="input" type="primary">
      复制到剪切板
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.box {
  display: flex;
  .el-input {
    margin-right: 10px;
  }
}
</style>
