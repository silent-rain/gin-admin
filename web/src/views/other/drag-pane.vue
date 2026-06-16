<script setup lang="ts">
import { Pane, Splitpanes } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'
</script>

<template>
  <Splitpanes class="default-theme" horizontal :push-other-panes="false">
    <Pane>
      <div class="item" style="background-color: #fe5308">
        1
      </div>
    </Pane>
    <Pane>
      <Splitpanes :push-other-panes="false">
        <Pane>
          <div class="item" style="background-color: #fefe34">
            2
          </div>
        </Pane>
        <Pane>
          <div class="item" style="background-color: #cfea2b">
            3
          </div>
        </Pane>
        <Pane>
          <div class="item" style="background-color: #66b032">
            4
          </div>
        </Pane>
      </Splitpanes>
    </Pane>
    <Pane>
      <div class="item" style="background-color: #0492ce">
        5
      </div>
    </Pane>
  </Splitpanes>
</template>

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
