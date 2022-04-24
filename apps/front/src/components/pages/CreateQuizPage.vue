<template>
  <div class="create-quiz-page">
    <MainHeader />
    <ContentsCreator :id="undefined" />
    <MainFooter />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import MainFooter from '@/components/organisms/MainFooter.vue'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'

export default defineComponent({
  name: 'CreateQuizPage',

  components: {
    MainHeader,
    MainFooter,
    ContentsCreator,
  },
});
</script>
