<template>
  <div class="grid">
    <div class="mt-6 col-offset-1 col-10 md:col-10 lg:col-6 lg:col-offset-3">
      <QuizInfo :contents="contents" :interval="interval" />
    </div>
  </div>
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
