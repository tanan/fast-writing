import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '../components/pages/TopPage.vue'
import WritingQuizPage from '../components/pages/WritingQuizPage.vue'
import CreateQuizPage from '../components/pages/CreateQuizPage.vue'

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
  {
    path: '/contents/create',
    name: 'CreateQuizPage',
    component: CreateQuizPage
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
