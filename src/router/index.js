// src/router/index.js

import { createRouter, createWebHistory } from 'vue-router'
import App from '@/App.vue'
import Home from '@/views/Home.vue'
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    children: [
      // Add your routes here
      {
        path: '/round1',
        name: 'Round1',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Round1.vue')
      },
      {
        path: '/round2',
        name: 'Round2',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Round2.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory('/voting'),
  base: '/voting',
  routes
})

export default router
