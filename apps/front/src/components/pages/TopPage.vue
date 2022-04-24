<template>
  <div class="top-page">
    <MainHeader />
    <ContentsList />
    <MainFooter />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import ContentsList from '@/components/organisms/ContentsList.vue'
import MainHeader from '@/components/organisms/MainHeader.vue'
import MainFooter from '@/components/organisms/MainFooter.vue'

export default defineComponent({
  name: 'TopPage',

  components: {
    ContentsList,
    MainHeader,
    MainFooter,
  },
});
</script>
