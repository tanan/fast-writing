<template>
  <div class="top-page">
    <ContentsList />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import ContentsList from '@/components/organisms/ContentsList.vue'

export default defineComponent({
  name: 'TopPage',

  components: {
    ContentsList,
  },
});
</script>
