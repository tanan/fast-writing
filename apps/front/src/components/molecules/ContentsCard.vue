<template>
  <Card class="card">
    <template #title>
      {{ contents.title }}
    </template>
    <template #subtitle>
      <div class="flex">
        <img v-if="hasIcon" src="http://localhost:8888/images/blocks/avatars/circle/avatar-f-1.png" alt="user-icon" class="icon align-items-center">
        <div v-if="!hasIcon" class="icon flex align-items-center justify-content-center bg-gray-200 border-circle">
          <i class="pi pi-user text-gray-500 text-xl"></i>
        </div>
        <div class="flex flex-column align-items-start">
          <div class="creator align-items-center ml-2">{{ contents.creator }}</div>
          <div class="last-updated align-items-center ml-2">{{ getUpdatedDate(contents.lastUpdated) }}</div>
        </div>
      </div>
    </template>
    <template #content>
        <div class="description">{{ contents.description }}</div>
    </template>
    <template #footer>
      <a :href="`/contents/${contents.id}`" rel="noopener" class="link">
        <Button icon="pi pi-check" label="Start" to="`/contents/${contents.id}`" />
      </a>
      <a v-show="editable" :href="`/contents/edit/${contents.id}`" rel="noopener" class="link">
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
    contents: Object,
    editable: Boolean
  },
  setup () {
    const hasIcon = false
    const getUpdatedDate = (v) => {
      if (!v) {
        return
      }
      let date = new Date(v.seconds*1000)
      return date.getFullYear() + "/" + (date.getMonth() + 1) + "/" + date.getDate()
    }

    return {
      hasIcon,
      getUpdatedDate,
    }
  },
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
  .icon {
    width: 3rem;
    height: 3rem;
  }
</style>