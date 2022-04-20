<template>
  <div class="edit-quiz-page">
    <ContentsCreator :id="id" />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'

export default defineComponent({
  name: 'EditQuizPage',
  components: {
    ContentsCreator,
  },
  setup () {
    const route = useRoute()
    const id = parseInt(route.params.id)
    return {
      id
    }
  },
});
</script>
