<template>
  <div id="app">
    <LogViewer msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import LogViewer from './components/LogViewer.vue'

export default {
  name: 'App',
  components: {
    LogViewer: LogViewer
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  font-size: 14px;
  margin-top: 0px;
}
</style>
