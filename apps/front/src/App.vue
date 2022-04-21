<template>
  <div id="wrapper">
    <MainHeader />
    <router-view/>
    <footer>
      <MainFooter />
    </footer>
  </div>

</template>

<script>
import { defineComponent } from 'vue';
import Store from '@/store/index.js'
import MainFooter from '@/components/organisms/MainFooter.vue'
import MainHeader from '@/components/organisms/MainHeader.vue'

export default defineComponent({
  name: 'App',
  components: {
    MainFooter,
    MainHeader,
  },
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