<template>
  <v-container class="quiz-header">
    <QuizInfo :contents="contents" :interval="interval" />
  </v-container>    
</template>

<script>
import QuizInfo from '@/components/molecules/QuizInfo.vue'

export default {
  name: 'QuizHeader',
  components: {
    QuizInfo,
  },
  props: {
    contents: Object,
    interval: Number
  }
}
</script>
