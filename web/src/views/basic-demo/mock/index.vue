<script setup lang="ts">
import axios from 'axios'

function listReq() {
  axios.get('/getMapInfo').then(({ data }) => {
    elMessage(data)
  })
}
</script>

<template>
  <div class="scroll-y">
    <div class="font-bold mb-20px">
      mock使用示例
    </div>
    <el-button type="primary" @click="listReq">
      点击发送mock请求
    </el-button>
  </div>
</template>

<style scoped lang="scss"></style>
