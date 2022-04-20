<template>
  <div class="create-quiz-page">
    <ContentsCreator :id="undefined" />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'

export default defineComponent({
  name: 'CreateQuizPage',

  components: {
    ContentsCreator,
  },
});
</script>
