<template>
  <v-row>
    <v-col>
      <div class="text-h4">{{ contents.title }}</div>
      <div class="quiz-info text-subtitle-2">
        <div class="creator">作成者：{{ contents.creator }}</div>
        <div class="last-updated">更新日時：{{ getUpdatedDate(contents.lastUpdated) }}</div>
      </div>
      <div class="interval text-subtitle-1">回答時間：{{ interval }}秒</div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'QuizInfo',
  props: {
    contents: Object,
    interval: Number
  },
  methods: {
    getUpdatedDate(v) {
      if (!v) {
        return
      }
      let date = new Date(v.seconds*1000)
      return date.getFullYear() + "/" + date.getMonth() + "/" + date.getDate()
    }
  }
}
</script>

<style lang="scss">
  .quiz-info {
    display: flex;
    font-style: oblique;

    .last-updated {
      padding-left: 6%;
    }
  }


  .interval {
    margin-top: 2%;
  }
</style>