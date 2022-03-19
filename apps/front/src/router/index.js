import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/pages/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
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
