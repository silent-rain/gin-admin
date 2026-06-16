<script setup>
import { QuestionFilled } from '@element-plus/icons-vue'
import { driver } from 'driver.js'
import { useI18n } from 'vue-i18n'
import steps from './steps'
import 'driver.js/dist/driver.css'

const { t } = useI18n()

function guide() {
  const driverObj = driver({
    showProgress: true,
    steps,
  })
  driverObj.drive()
}
</script>

<template>
  <div class="app-container">
    <aside class="font-bold">
      {{ t('guide.description') }}
      <a href="https://github.com/kamranahmedse/driver.js" target="_blank">
        driver.js.
      </a>
    </aside>
    <el-button
      class="mt-20px"
      :icon="QuestionFilled"
      type="primary"
      @click.prevent.stop="guide"
    >
      {{ t('guide.button') }}
    </el-button>
  </div>
</template>
