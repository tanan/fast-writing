<template>
  <v-app>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
  beforeCreate () {
    this.$store.dispatch('loadState')
    this.$store.dispatch('loadInterval')
  }
}
</script>
