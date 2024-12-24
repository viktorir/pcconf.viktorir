<template>
  <n-layout style="min-height: 100vh; display: flex; flex-direction: column;">
    <Header/>

    <n-layout-content style="height: calc(100vh - 74px); display: flex;" > 
      <router-view/>
    </n-layout-content>

    <n-layout-footer bordered style="display: flex; justify-content: center;">
      © 2024-2025 sso.viktorir. All rights reserved.
    </n-layout-footer>
  </n-layout>
</template>

<script>
import { NLayout, NLayoutHeader, NLayoutFooter, NLayoutContent, NButtonGroup, useNotification } from 'naive-ui';
import { MoonOutline, SunnyOutline } from '@vicons/ionicons5';
import Header from './components/Header.vue'


export default {
  components: {
    NLayout, 
    NLayoutFooter,
    NLayoutHeader,
    NLayoutContent,
    NButtonGroup, 
    MoonOutline, 
    SunnyOutline,
    Header,
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  
}

a:active,
a:hover,  
a {
  text-decoration: inherit;
  color: inherit;
}

.n-layout-scroll-container {
  display: flex;
  flex-direction: column;
  flex: 1;
  height: 100%; 
}
</style>