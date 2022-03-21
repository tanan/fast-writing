<template>
  <div>
    <v-row no-gutters class="question">
    {{ index }}. {{ question }}
    </v-row>
    <v-row no-gutters class="answer">
      {{ getAnswer() }}
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'QuizText',
  props: {
    index: Number,
    question: String,
    answer: String,
  },
  methods: {
    getAnswer() {
      if (this.answer === "") {
        return ""
      }
      return "A. " + this.answer
    }
  }
}
</script>

<style lang="scss">
  .answer {
    color: red;
    // padding-left: 4%;
  }
</style>