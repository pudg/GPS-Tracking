import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Tracking from '../views/Tracking.vue'
import Login from '../views/Login.vue'
import Products from '../views/Products.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/tracking',
      name: 'tracking',
      component: Tracking
    },
    {
      path: '/products',
      name: 'products',
      component: Products
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
  ]
})

export default router
