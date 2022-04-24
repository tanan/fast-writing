<template>
  <div class="edit-quiz-page">
    <MainHeader />
    <ContentsCreator :id="id" />
    <MainFooter />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router'
import MainHeader from '@/components/organisms/MainHeader.vue'
import MainFooter from '@/components/organisms/MainFooter.vue'
import ContentsCreator from '@/components/organisms/ContentsCreator.vue'

export default defineComponent({
  name: 'EditQuizPage',
  components: {
    ContentsCreator,
    MainHeader,
    MainFooter
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
