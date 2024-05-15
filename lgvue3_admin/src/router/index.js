import { createRouter, createWebHashHistory } from 'vue-router'
import layout from '@/layout/deflayout.vue'
// 2. 
const routes = [
    {
        path: '/',
        redirect: '/index',
        component: layout,
        children: [
          {
            path: 'index',
            name: 'index',
            meta: { title: 'Wrok', icon: 'mdi-home' },
            component: () => import('@/page/index.vue'),
          },
          {
            path: 'add',
            name: 'add',
            meta: { title: 'New', icon: 'mdi-home' },
            component: () => import('@/page/add.vue'),
          },
        ]
    }
]

// 3. 
const router = createRouter({
  history: createWebHashHistory(),  // 这里采用的是 hash 模式
  routes                            // 上面声明的路由规则
})

// 4. 
export default router