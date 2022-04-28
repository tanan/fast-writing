<template>
  <div class="flex flex-column">
    <div class="flex mt-4 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3" v-for="quiz in display.quizList" :key="quiz.index">
      <QuizText :quiz="quiz" />
    </div>
    <Button v-show="display.showAddButton" class="lg:col-offset-3 col-1 mt-2 mb-4 lg:mt-4" :label="getLabel()" icon="pi pi-plus" @click="addNewQuiz()" />
    <div class="flex h-30rem"></div>
  </div>
</template>

<script>
import { defineComponent, reactive } from "vue"
import Store from '@/store/index.js'
import QuizText from '@/components/molecules/QuizText.vue'
import Button   from 'primevue/button'
export default defineComponent ({
  name: 'QuizList',
  components: {
    QuizText,
    Button,
  },
  setup (props) {
    const display = reactive({
      showAddButton: true,
      quizList: []
    })

    const getLabel = () => {
      return display.quizList.length === 0 ? "クイズ開始" : "次の問題へ"
    }

    const addNewQuiz = () => {
      let quiz = Store.getters.getNextQuiz
      quiz.isDisplayAnswer = false
      display.quizList.push(quiz)
      Store.dispatch('shiftQuizList')
      display.showAddButton = Store.getters.hasQuizElement ? true : false
    }

    const removeTimer = () => {
      console.log(props.timerId)
      window.clearTimeout(props.timerId)
    }
    return {
      display,
      getLabel,
      addNewQuiz,
      removeTimer,
    }
  }
})
</script>
