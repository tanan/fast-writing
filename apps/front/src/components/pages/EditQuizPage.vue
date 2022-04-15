<template>
  <MainHeader :isLoggedIn="isLoggedIn" />
  <div class="edit-quiz-page">
    <ContentsCreator />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'
import Store from '@/store/index.js'

export default defineComponent({
  name: 'EditQuizPage',
  components: {
    MainHeader,
    ContentsCreator,
  },
  setup () {
    const isLoggedIn = Store.getters.isLoggedIn()
    return {
      isLoggedIn
    }
  },
});
</script>
