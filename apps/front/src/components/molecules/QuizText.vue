<template>
  <div class="quiz-text">
    <div class="question">{{ getIndex(quiz.index) }}. {{ quiz.question }}</div>
    <div class="answer mt-2">A. 
      <Button v-if="!isDisplayAnswer" label="解答を表示" icon="pi pi-plus" @click="click()" />
      <span v-if="isDisplayAnswer">{{ quiz.answer }}</span>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue"
import Button   from 'primevue/button'

export default defineComponent ({
  name: 'QuizText',
  props: {
    quiz: Object
  },
  components: {
    Button
  },
  setup (props) {
    const isDisplayAnswer = ref(props.isDisplayAnswer)
    const getAnswer = () => {
      if (props.answer === "") {
        return ""
      }
      return "A. " + props.answer
    }

    const showButton = () => {
      return !isDisplayAnswer.value
    }

    const click = () => {
      isDisplayAnswer.value = true
    }

    const getIndex = (i) => {
      return parseInt(i) + 1
    }


    return {
      isDisplayAnswer,
      showButton,
      click,
      getIndex,
      getAnswer,
    }
  }
})
</script>

<style lang="scss">
  .quiz-text {
    font-size: 1.2rem;
    .answer {
      color: red;
      icon: {
        padding-left: 0.6rem;
      }
    }
  }
</style>