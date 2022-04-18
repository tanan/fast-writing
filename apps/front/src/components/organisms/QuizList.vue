<template>
  <div class="grid">
    <div class="mt-4 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3" v-for="(question, index) in questions" :key="index">
      <QuizText :index="getIndex(index)" :question="question" :answer="getAnswer(index)" />
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import QuizText from '@/components/molecules/QuizText.vue'
export default defineComponent ({
  name: 'QuizList',
  components: {
    QuizText
  },
  props: {
    indecies: Array,
    questions: Array,
    answers: Array,
  },
  setup (props) {
    const getIndex = (i) => {
      return parseInt(i) + 1
    }
    const getAnswer = (i) => {
      let answer = props.answers[i]
      if (!answer) {
        return ""
      } else {
        return answer
      }
    }

    return {
      getIndex,
      getAnswer,
    }
  }
})
</script>
