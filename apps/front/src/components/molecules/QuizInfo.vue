<template>
  <div>
    <div class="title">
      <h2 class="main">{{ contents.title }}</h2>
      <p class="sub">{{ contents.description }}</p>
    </div>
    <div class="flex">
      <div class="tags mt-2">
        <Tag class="mr-1" severity="success" v-for="tag in getTags(contents.tags)" :key="tag" :value="tag"></Tag>
      </div>
    </div>
    <div class="quiz-info mt-3">
      <p class="">作成者：{{ contents.creator }}</p>
      <p class="ml-4">更新日時：{{ getUpdatedDate(contents.lastUpdated) }}</p>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue"
import Tag from 'primevue/tag'

export default defineComponent ({
  name: 'QuizInfo',
  components: {
    Tag,
  },
  props: {
    contents: Object,
    interval: Number
  },
  setup () {
    const getUpdatedDate = (v) => {
      if (!v) {
        return
      }
      let date = new Date(v.seconds*1000)
      return date.getFullYear() + "/" + (date.getMonth() + 1) + "/" + date.getDate()
    }

    const getTags = (tags) => {
      if (!tags) {
        return
      }
      return tags.split(',')
    }

    return {
      getUpdatedDate,
      getTags,
    }
  },
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