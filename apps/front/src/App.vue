<template>
  <div id="wrapper">
    <router-view/>
    <footer>
      <MainFooter />
    </footer>
  </div>

</template>

<script>
import MainFooter from '@/components/organisms/MainFooter.vue'
export default {
  name: 'App',
  components: {
    MainFooter,
  },
  beforeCreate () {
    this.$store.dispatch('loadState')
    this.$store.dispatch('loadInterval')
  }
}
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