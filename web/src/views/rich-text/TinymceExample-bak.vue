<script setup>
import Tinymce from './tinymce/index.vue'
/* tinymce操作 */
const refTinymce = ref(null)
function setTinyContent() {
  if (refTinymce?.value) {
    refTinymce.value.setContent('tinymce设置的内容')
  }
}
function getTinyContent() {
  if (refTinymce?.value) {
    elMessage(refTinymce.value.getContent())
  }
}
function clearTinyContent() {
  if (refTinymce?.value) {
    refTinymce.value.setContent('')
  }
}
</script>

<template>
  <div class="scroll-y">
    <div class="mb-10px">
      <el-button @click="setTinyContent">
        设置文本消息
      </el-button>
      <el-button @click="getTinyContent">
        获取文本消息
      </el-button>
      <el-button @click="clearTinyContent">
        清空内容
      </el-button>
    </div>
    <Tinymce ref="refTinymce" />
  </div>
</template>

<style scoped lang="scss"></style>
