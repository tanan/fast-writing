<template>
  <MainHeader :isLoggedIn="isLoggedIn" />
  <div class="edit-quiz-page">
    <ContentsCreator :id="id" />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router'
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
    const route = useRoute()
    const id = route.params.id
    const isLoggedIn = Store.getters.isLoggedIn
    return {
      isLoggedIn,
      id
    }
  },
});
</script>
