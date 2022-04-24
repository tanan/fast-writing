<template>
  <div id="wrapper">
    <router-view/>
  </div>

</template>

<script>
import { defineComponent } from 'vue';
import Store from '@/store/index.js'

export default defineComponent({
  name: 'App',
  setup () {
    Store.dispatch('loadState')
    Store.dispatch('loadInterval')
  }
})
</script>

<style lang="scss">
  body,
  #wrapper {
    display: flex;
    flex-direction: column;
    min-height: calc(100vh - 46px);
  }
  footer {
    margin-top: auto;
  }
</style>