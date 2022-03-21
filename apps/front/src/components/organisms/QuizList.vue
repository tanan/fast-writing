<template>
  <v-container>
    <div class="quiz" v-for="(question, index) in questions" :key="index">
      <QuizText :index="getIndex(index)" :question="question" :answer="getAnswer(index)" />
    </div>
  </v-container>
</template>

<script>
import QuizText from '@/components/molecules/QuizText.vue'
export default {
  name: 'QuizList',
  components: {
    QuizText
  },
  props: {
    indecies: Array,
    questions: Array,
    answers: Array,
  },
  methods: {
    getIndex(i) {
      return parseInt(i) + 1
    },
    getAnswer(i) {
      console.log(i)
      let answer = this.answers[i]
      if (!answer) {
        return ""
      } else {
        return answer
      }
    }
  }
}
</script>

<style lang="scss">
  .quiz {
    margin: 16px 0;
  }
</style>

