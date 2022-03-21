<template>
  <div class="header">
    <div class="top"><MainLogo /></div>
  </div>    
</template>

<script>
import MainLogo from '@/components/atoms/MainLogo.vue'

export default {
  name: 'MainHeader',
  components: {
    MainLogo,
  },
  props: {
    page: String
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header {
  height: 96px;
  background-color: #55c500;

  .top {
    height: 48px;
  }

  .content {
    height: 48px;
  }
}
</style>