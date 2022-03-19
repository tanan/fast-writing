<template>
  <MainHeader />
  <ContentsList />
</template>

<script>
import { defineComponent } from 'vue';
import MainHeader from '@/components/organisms/MainHeader.vue'
import ContentsList from '@/components/organisms/ContentsList.vue'

export default defineComponent({
  name: 'TopPage',

  components: {
    MainHeader,
    ContentsList,
  },
});
</script>
