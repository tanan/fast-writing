<template>
  <div class="header">
    <v-container class="ma-0 pa-2">
      <v-row no-gutters>
        <v-col cols="12" sm="3"><MainLogo /></v-col>
        <v-col cols="10" offset="1" sm="4" class="d-flex align-center"><HeaderForm /></v-col>
      </v-row>
    </v-container>
  </div>    
</template>

<script>
import MainLogo from '@/components/atoms/MainLogo.vue'
import HeaderForm from '@/components/molecules/HeaderForm.vue'

export default {
  name: 'MainHeader',
  components: {
    MainLogo,
    HeaderForm,
  },
  props: {
    page: String
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header {
  background-color: #55c500;
}
</style>