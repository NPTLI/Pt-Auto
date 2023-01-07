<script setup lang="ts">
import { watch } from 'vue'
import HelloWorld from './components/HelloWorld.vue'


const switchMode = () => {
  console.log(localStorage.theme);
  // Whenever the user explicitly chooses dark mode
  localStorage.theme = localStorage.theme == 'dark' ? 'light' : 'dark'
  if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }


}

</script>

<template>
  <div class="page-content dark:bg-black">
    <div @click="switchMode">切换模式</div>
    <div class="flex items-center justify-center">
      <a href="https://vitejs.dev" target="_blank">
        <img src="/vite.svg" class="logo" alt="Vite logo" />
      </a>
      <a href="https://vuejs.org/" target="_blank">
        <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
    </div>
    <HelloWorld msg="Vite + Vue" />
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

function watch(arg0: [any, {}]) {
  throw new Error('Function not implemented.')
}
