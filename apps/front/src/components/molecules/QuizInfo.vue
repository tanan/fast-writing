<template>
  <div>
    <div class="title">
      <h2>{{ contents.title }}</h2>
    </div>
    <div class="quiz-info mt-2">
      <p class="">作成者：{{ contents.creator }}</p>
      <p class="ml-4">更新日時：{{ getUpdatedDate(contents.lastUpdated) }}</p>
    </div>
    <div class="flex mt-4">
      <h3 class="mr-2 pt-2">回答時間:</h3>
      <Dropdown v-model="select" :options="items" />
      <p class="ml-2 pt-2">秒</p>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import Dropdown from 'primevue/dropdown';

export default defineComponent ({
  name: 'QuizInfo',
  components: {
    Dropdown,
  },
  data () {
    return {
      select: this.interval,
      items: [4,5,6,7,8,9,10],
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