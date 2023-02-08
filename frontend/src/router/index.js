import { createRouter, createWebHistory } from 'vue-router';
import AppLayout from '@/layout/AppLayout.vue';
import HomeView from '../views/HomeView.vue';
import Landing from '../views/Landing.vue';
import Login from '../views/Login.vue';
import Chart from '@/views/Chart.vue';
import OsCrud from '@/views/OsCrud.vue';
import UserCrud from '@/views/UserCrud.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            component: Landing
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/home',
            component: AppLayout,
            children: [
                {
                    path: '/home',
                    name: 'Home',
                    component: HomeView
                },
                {
                    path: '/charts',
                    name: 'charts',
                    component: Chart
                },
                {
                    path: '/os',
                    name: 'OS',
                    component: OsCrud
                },
                {
                    path: '/user',
                    name: 'User',
                    component: UserCrud
                }

                /* ,
                {
                    path: '/about',
                    name: 'about',
                    // route level code-splitting
                    // this generates a separate chunk (About.[hash].js) for this route
                    // which is lazy-loaded when the route is visited.

                    component: () => import('../views/AboutView.vue')
                } */
            ]
        }
    ]
});

export default router;
