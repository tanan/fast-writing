<template>
  <div class="quiz-text">
    <div class="question">{{ index }}. {{ question }}</div>
    <div class="answer mt-2">{{ getAnswer() }}</div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
export default defineComponent ({
  name: 'QuizText',
  props: {
    index: Number,
    question: String,
    answer: String,
  },
  setup (props) {
    const getAnswer = () => {
      if (props.answer === "") {
        return ""
      }
      return "A. " + props.answer
    }

    return {
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