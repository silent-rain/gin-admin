<template>
  <component :is="Component" />
</template>

<script setup lang="ts">
import LoginAlt from './login-alt.vue';
import LoginBase from './login-basic.vue';
import LoginLighting from './login-lighting.vue';

const renderNum = Math.floor(Math.random() * 10) + 1;

const Component = computed(() => {
  if (renderNum <= 2) {
    return LoginBase;
  }
  if (renderNum <= 4 && renderNum > 2) {
    return LoginLighting;
  }
  return LoginAlt;
});
</script>

<style scoped lang="scss"></style>
