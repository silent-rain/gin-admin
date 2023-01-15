<!--
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:54
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 16:48:46
 * @company: 
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/views/other/drag-pane.vue
 * @Descripttion: 
-->
<template>
  <splitpanes class="default-theme" horizontal :push-other-panes="false">
    <pane>
      <div class="item" style="background-color: #fe5308">1</div>
    </pane>
    <pane>
      <splitpanes :push-other-panes="false">
        <pane>
          <div class="item" style="background-color: #fefe34">2</div>
        </pane>
        <pane>
          <div class="item" style="background-color: #cfea2b">3</div>
        </pane>
        <pane>
          <div class="item" style="background-color: #66b032">4</div>
        </pane>
      </splitpanes>
    </pane>
    <pane>
      <div class="item" style="background-color: #0492ce">5</div>
    </pane>
  </splitpanes>
</template>

<script setup lang="ts">
import { Pane, Splitpanes } from 'splitpanes';
import 'splitpanes/dist/splitpanes.css';
</script>

<style lang="scss" scoped>
.item {
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: Helvetica, Arial, sans-serif;
  color: #333;
  font-size: 5em;
  width: 100%;
  height: 100%;
}
</style>
