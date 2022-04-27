<template>
  <div class="flex flex-column">
    <div class="flex mt-4 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3" v-for="(question, index) in questions" :key="index">
      <QuizText :index="getIndex(index)" :question="question" :answer="getAnswer(index)" />
    </div>
    <!-- <Button class="lg:col-offset-3 col-1 mt-2 mb-4 lg:mt-4" label="解答を表示" icon="pi pi-plus" @click="removeTimer()" /> -->
    <div class="flex h-30rem"></div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import QuizText from '@/components/molecules/QuizText.vue'
// import Button from 'primevue/button'
export default defineComponent ({
  name: 'QuizList',
  components: {
    QuizText,
    // Button,
  },
  props: {
    indecies: Array,
    questions: Array,
    answers: Array,
    timerId: Number,
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

    const removeTimer = () => {
      console.log(props.timerId)
      window.clearTimeout(props.timerId)
    }

    return {
      getIndex,
      getAnswer,
      removeTimer,
    }
  }
})
</script>
