<template>
  <div class="quiz-text fadein animation-duration-500">
    <div class="question" @click="speak('ja')">{{ getIndex(quiz.index) }}. {{ quiz.question }}</div>
    <div class="answer mt-2">A. 
      <Button v-if="!isDisplayAnswer" label="解答を表示" icon="pi pi-plus" class="button p-button-sm p-button-danger ml-2" @click="click()" />
      <transition name="answer">
        <span v-if="isDisplayAnswer" class="answer-text" @click="speak('en')">{{ quiz.answer }}</span>
      </transition>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue"
import Store from '@/store/index.js'
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

    const click = async () => {
      isDisplayAnswer.value = true
      await new Promise((resolve) => setTimeout(resolve, 1000))
      Store.dispatch('updateDispalyAnswer', true)
    }

    const getIndex = (i) => {
      return parseInt(i) + 1
    }

    const speak = (lang) => {
      let uttr = new SpeechSynthesisUtterance()
      if(lang == 'ja') {
        uttr.text = props.quiz.question
        uttr.lang = 'ja'
      } else {
        uttr.text = props.answer
        uttr.lang = 'en-US'
      }
      speechSynthesis.speak(uttr)
    }

    return {
      isDisplayAnswer,
      showButton,
      click,
      getIndex,
      getAnswer,
      speak,
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

  .answer-enter-active, .answer-leave-active {
    transition: opacity .8s ease;
  }

  .answer-enter-from, .answer-leave-to {
    opacity: 0;
  }
</style>