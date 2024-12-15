import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
  {
    path: "/solid_career",
    component: () => import('@/view/solidcareer/layOut.vue'),
    children:[
      {
        path:"",
        name: 'indexPage',
        component: () => import('@/view/solidcareer/homeIndex.vue'),
        redirect: '/solid_career/home'
      },
      {
        path:"/solid_career/home",
        name: 'homeIndex',
        component: () => import('@/view/solidcareer/homeIndex.vue'),
        meta: {
          title: '首页'
        }
      },
      {
        path: '/solid_career/login',
        name: 'loginIndex',
        component: () =>import('@/view/solidcareer/loginIndex.vue'),
        meta: {
          title: '登录页'
        }
      },
      {
        path: '/solid_career/profile',
        name: 'profileIndex',
        component: () => import('@/view/solidcareer/profileIndex.vue'),
        meta: {
          title: '个人主页'
        }
      },
      {
        path:"company",
        name: 'companyIndex',
        component: () =>  import('@/view/solidcareer/companyIndex.vue'),
        meta: {
          title: '公司'
        }
      },
    ]
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
