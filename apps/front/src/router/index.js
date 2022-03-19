import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '../components/pages/TopPage.vue'

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  }
  // {
  //   path: '/article/:id',
  //   name: 'Article',
  //   component: Article
  // },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
