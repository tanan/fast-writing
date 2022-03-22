<template>
  <MainHeader />
  <div class="create-quiz-page">
    <ContentsCreator />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'

export default defineComponent({
  name: 'CreateQuizPage',

  components: {
    MainHeader,
    ContentsCreator,
  },
});
</script>
