import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '@/pages/Home.vue'
import {loadLayoutMiddleware} from '@/router/middlewares/loadLayout.middleware'
import {AppLayoutsEnum} from "@/layouts/layouts.types";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            meta: {
                layout: AppLayoutsEnum.auth
            },

            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('@/pages/Auth/Login.vue')
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
