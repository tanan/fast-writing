<template>
  <div class="flex flex-column">
    <div class="flex mt-4 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3" v-for="quiz in display.quizList" :key="quiz.index">
      <QuizText :quiz="quiz" />
    </div>
    <div v-if="showAddButton()"
      class="fadein animation-duration-400 bg-orange-500 text-white
        mt-2 mb-4 col-offset-1 col-4 lg:col-offset-3 lg:col-1 lg:mt-4
        flex align-items-center justify-content-center font-bold border-round cursor-pointer"
      @click="addNewQuiz()">{{ getLabel() }}</div>
    <div class="flex h-30rem"></div>
  </div>
</template>

<script>
import { defineComponent, reactive } from "vue"
import Store from '@/store/index.js'
import QuizText from '@/components/molecules/QuizText.vue'
export default defineComponent ({
  name: 'QuizList',
  components: {
    QuizText,
  },
  setup (props) {
    const display = reactive({
      quizList: []
    })

    const getLabel = () => {
      return display.quizList.length === 0 ? "クイズ開始" : "次の問題へ"
    }

    const showAddButton = () => {
      return Store.state.isDisplayAnswer && Store.getters.hasQuizElement
    }

    const addNewQuiz = () => {
      let quiz = Store.getters.getNextQuiz
      quiz.isDisplayAnswer = false
      display.quizList.push(quiz)
      Store.dispatch('shiftQuizList')
      Store.dispatch('updateDispalyAnswer', false)
    }

    const removeTimer = () => {
      console.log(props.timerId)
      window.clearTimeout(props.timerId)
    }
    return {
      display,
      getLabel,
      showAddButton,
      addNewQuiz,
      removeTimer,
    }
  }
})
</script>
