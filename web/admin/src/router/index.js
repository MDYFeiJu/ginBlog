import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '@/views/Login'
import Admin from '@/views/Admin'

import Index from '@/components/admin/Index'
import AddArt from '@/components/article/AddArt'
import ArtList from '@/components/article/ArtList'
import CateList from '@/components/category/CateList'
import UserList from '@/components/user/UserList'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
    children: [
      { path: 'index', component: Index },
      { path: 'addart', component: AddArt },
      { path: 'artlist', component: ArtList },
      { path: 'catelist', component: CateList },
      { path: 'userlist', component: UserList },
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/admin') {
    next('/login')
  }
  next()
})
export default router
