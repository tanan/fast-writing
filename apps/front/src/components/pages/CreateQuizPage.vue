<template>
  <MainHeader :isLoggedIn="isLoggedIn" />
  <div class="create-quiz-page">
    <ContentsCreator :id="undefined" />
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
  setup () {
    const isLoggedIn = Store.getters.isLoggedIn
    return {
      isLoggedIn
    }
  },
});
</script>
