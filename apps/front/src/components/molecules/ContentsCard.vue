<template>
  <v-card class="contents-card" elevation="2">
    <v-card-title>{{ contents.title }}</v-card-title>
    <v-card-subtitle class="info">
      <div class="creator">作成者: {{ contents.creator }}</div>
      <div class="last-updated">更新日時: {{ getUpdatedDate(contents.lastUpdated) }}</div>
    </v-card-subtitle>
    <v-card-text class="description">{{ contents.description }}</v-card-text>
    <v-card-actions>
      <v-btn color="blue" text>
        <a :href="`/contents/${contents.id}`" rel="noopener" class="link">Go to Lesson</a>
      </v-btn>
    </v-card-actions>
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

    .info {
      display: flex;
      padding: 4px 16px;
      font-weight: 0.8rem;
      .last-updated {
        padding-left: 24px;
      }
    }

    .link {
      text-decoration: none;
    }
  }
</style>