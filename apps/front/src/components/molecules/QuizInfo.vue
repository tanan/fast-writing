<template>
  <v-row>
    <v-col>
      <div class="text-h4">{{ contents.title }}</div>
      <div class="quiz-info text-subtitle-2">
        <div class="creator">作成者：{{ contents.creator }}</div>
        <div class="last-updated">更新日時：{{ getUpdatedDate(contents.lastUpdated) }}</div>
      </div>
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="4" sm="2">
      <div class="interval text-subtitle-1">回答時間：</div>
    </v-col>
    <v-col cols="4" sm="2">
      <v-select v-model="select" :items="items" dense solo label="秒数"></v-select>
    </v-col>
  </v-row>
</template>

<script>
import { defineComponent } from "vue"

export default defineComponent ({
  name: 'QuizInfo',
  data () {
    return {
      select: this.interval,
      items: [1,2,3,4,5],
    }
  },
  props: {
    contents: Object,
    interval: Number
  },
  watch: {
    select: function(newVal, oldVal) {
      console.log("change interval from " + oldVal + " to " + newVal)
      this.$store.dispatch("interval", {
        interval: newVal
      })
      this.$router.go()
    }
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
})
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