import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Devices from '../views/Devices.vue'
import Login from '../views/Login.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/devices',
      name: 'devices',
      component: Devices
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
  ]
})

export default router
