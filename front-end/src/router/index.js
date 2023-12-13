import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'homeview',
      component: () => import('../views/HomeView.vue')
    },
    // {
    //   path: '/adminview',
    //   name: 'adminview',
    //   component: () => import('../views/AdminView.vue')
    // },
    {
      path: '/test',
      name: 'test',
      component: () => import('../views/AdminView.vue')
    },
  ]
})

export default router
