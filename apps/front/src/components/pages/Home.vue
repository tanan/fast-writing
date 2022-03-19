<template>
  <div class="home">
    <MainHeader page="home" />
    <div class="home-container">
      <ContentsList />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import MainHeader from '@/components/organisms/MainHeader.vue'
import ContentsList from '@/components/organisms/ContentsList.vue'

export default {
  name: 'Home',
  components: {
    MainHeader,
    ContentsList,
  },
 }
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .home-container {
    display: flex;
  }
</style>