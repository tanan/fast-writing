<template>
  <v-card class="contents-card" elevation="2">
    <v-card-title>{{ contents.title }}</v-card-title>
    <v-card-text class="creator">作成者: {{ contents.creator }}</v-card-text>
    <v-card-text class="last-updated">更新日時: {{ getUpdatedDate(contents.lastUpdated) }}</v-card-text>
    <v-card-text class="description">買い物で使う英文を中心に練習します</v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'ContentsCard',
  props: {
    contents: Object
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

<style lang="scss" scoped>
  .contents-card {
    margin: 8px 16px;

    .creator {
      padding: 8px 16px;
      padding-bottom: 4px;
    }
    .last-updated {
      padding: 4px 16px;
    }
  }
</style>