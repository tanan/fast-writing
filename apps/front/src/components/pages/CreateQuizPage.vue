<template>
  <MainHeader :isLoggedIn="this.isLoggedIn()" />
  <div class="create-quiz-page">
    <ContentsCreator />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'
import Store from '@/store/index.js'

export default defineComponent({
  name: 'CreateQuizPage',

  components: {
    MainHeader,
    ContentsCreator,
  },
  methods: {
    isLoggedIn () {
      if (Store.state.userToken === "") {
        return false
      }
      return true
    }
  },
});
</script>
