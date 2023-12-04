// src/router/index.js

import { createRouter, createWebHistory } from 'vue-router'
import App from '@/App.vue'
import Layout from '@/views/Layout.vue'
import Home from '@/views/Home.vue'
const routes = [
  {
    path: '/',
    name: 'Layout',
    component: Layout,
    redirect: '/home',
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
      },
      {
        path: '/home',
        name: 'Home',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        query: {
          flag:false
        },
        components: {
          Home
        }
      },
      {
        path: '/ranking',
        name: 'ranking',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Ranking.vue')
      },
      {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        components:{
          About: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
        }
        // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      }
    ]
  },
  {
    path: '/votes',
    name: 'Votes',
    component: () => import(/* webpackChunkName: "about" */ '../views/Votes.vue')
  },
  {
    path: '/history',
    name: 'History',
    component: () => import(/* webpackChunkName: "about" */ '../views/History.vue')
  }
]

const router = createRouter({
  history: createWebHistory('/'),
  base: '/',
  routes
})

export default router
