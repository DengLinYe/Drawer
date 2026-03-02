import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path:'/',
    name:'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path:'/bills',
    name:'Bills',
    component: () => import('../views/Bills.vue')
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router