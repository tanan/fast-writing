import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '../components/pages/TopPage.vue'
import WritingQuizPage from '../components/pages/WritingQuizPage.vue'

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  },
  {
    path: '/contents/:id',
    name: 'WritingQuizPage',
    component: WritingQuizPage
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
