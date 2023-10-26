import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Tracking from '../views/Tracking.vue'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import store from '../store';

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
      component: Tracking,
      meta: { requiresAuth: true},
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
  ]
});

router.beforeEach((to, from, next) => {
  const userLoggedIn = store.state.user !== null;
  if (to.meta.requiresAuth && !userLoggedIn) {
    next({name: 'login'});
  } else {
    next();
  }
})

export default router
