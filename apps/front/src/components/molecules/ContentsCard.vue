<template>
  <Card class="card">
    <template #title>
      {{ contents.title }}
    </template>
    <template #subtitle>
      <div class="flex">
        <div class="creator">作成者: {{ contents.creator }}</div>
        <div class="last-updated ml-4">更新日時: {{ getUpdatedDate(contents.lastUpdated) }}</div>
      </div>
    </template>
    <template #content>
        <div class="description">{{ contents.description }}</div>
    </template>
    <template #footer>
      <a :href="`/contents/${contents.id}`" rel="noopener" class="link">
        <Button icon="pi pi-check" label="Start" to="`/contents/${contents.id}`" />
      </a>
      <a v-show="isPrivate()" :href="`/contents/edit/${contents.id}`" rel="noopener" class="link">
        <Button icon="pi pi-pencil" label="Edit" class="p-button-secondary" style="margin-left: .5em" />
      </a>
    </template>
  </Card>
</template>

<script>
import { defineComponent } from 'vue'
import Card from 'primevue/card'
import Button from 'primevue/button'

export default defineComponent({
  name: 'ContentsCard',
  components: {
    Card,
    Button,
  },
  props: {
    contents: Object
  },
  setup (props) {
    const isPrivate = () => {
      console.log(props.contents.scope)
      if (props.contents.scope === "private") {
        return true
      }
      return false
    }

    return {
      isPrivate,
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

<style lang="scss" scoped>
  .link {
    text-decoration: none;
  }
  .card {
    width: 100%;
  }
  .description {
    min-height: 3rem;
  }
</style>