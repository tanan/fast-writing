import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '../components/pages/TopPage.vue'
import AccountPage from '../components/pages/AccountPage.vue'
import SigninPage from '../components/pages/SigninPage.vue'
import SignupPage from '../components/pages/SignupPage.vue'
import WritingQuizPage from '../components/pages/WritingQuizPage.vue'
import CreateQuizPage from '../components/pages/CreateQuizPage.vue'
import EditQuizPage from '../components/pages/EditQuizPage.vue'
import Store from '@/store/index.js'

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage,
  },
  {
    path: '/account',
    name: 'AccountPage',
    component: AccountPage
  },
  {
    path: '/signin',
    name: 'SigninPage',
    component: SigninPage
  },
  {
    path: '/signup',
    name: 'SignupPage',
    component: SignupPage
  },
  {
    path: '/contents/:id',
    name: 'WritingQuizPage',
    component: WritingQuizPage,
    // meta: { requiresAuth: true }
  },
  {
    path: '/contents/create',
    name: 'CreateQuizPage',
    component: CreateQuizPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/contents/edit/:id',
    name: 'EditQuizPage',
    component: EditQuizPage,
    meta: { requiresAuth: true }
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth) && Store.state.userToken === "") {
    next({ path: '/signin', query: { redirect: to.fullPath } });
  } else {
    next();
  }
});

export default router
