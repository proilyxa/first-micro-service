import { createRouter, createWebHistory } from 'vue-router'
import { loadLayoutMiddleware } from '@/router/middlewares/loadLayout.middleware'
import { AppLayoutsEnum } from '@/layouts/layouts.types'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/pages/Dashboard/Index.vue')
    },
    {
      path: '/login',
      name: 'login',
      meta: {
        layout: AppLayoutsEnum.auth
      },
      component: () => import('@/pages/Auth/Login.vue')
    },
    {
      path: '/register',
      name: 'register',
      meta: {
        layout: AppLayoutsEnum.auth
      },
      component: () => import('@/pages/Auth/Register.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/pages/Errors/NotFound.vue')
    }
  ]
})

router.beforeEach(loadLayoutMiddleware)

export default router
