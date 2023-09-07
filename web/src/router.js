import {createRouter, createWebHistory} from 'vue-router'
import HomeView from './views/Home.vue'

const routes = [
    {
        path: '/',
        alias: ['/index.html'],
        name: 'home',
        component: HomeView
    },

    {
        path: '/login',
        name: 'login',
        component: () => import('./views/auth/Index.vue')
    },
    {
        path: '/about',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import('./views/About.vue')
    },
    // {
    //     path: '/:pathMatch(.*)*',
    //     name: 'NotFound',
    //     component: () => import('./views/NotFound.vue')
    // }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router