<template>
  <MainHeader :isLoggedIn="this.isLoggedIn()" />
  <div class="top-page">
    <ContentsList />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import Store from '@/store/index.js'
import ContentsList from '@/components/organisms/ContentsList.vue'

export default defineComponent({
  name: 'TopPage',

  components: {
    MainHeader,
    ContentsList,
  },
  data () {
    return {
      state: {
        count: 0
      }
    }
  },
  methods: {
    isLoggedIn () {
      if (Store.state.userToken === "") {
        return false
      }
      return true
    }
  }
});
</script>
