import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/Layout/index.vue'

Vue.use(Router)

export default new Router({
  // mode: 'history',
  routes: [
    // {
    //     path: '404',
    //     component: () => import('@/views/404.vue')
    // },
    /* {
            path: '/login',
            name: 'Login',
            component: () => import('@/views/Login')
        }, */
    {
      // 接口测试部分
      path: '/itf',
      name: 'itf',
      component: Layout,
      children: [
        {
          // 接口自动化部分
          path: 'itfPart',
          name: 'itfPart',
          component: () => import('@/views/InterfaceTestPart/ItfTestIndex.vue'),
          children: [
            // {
            //   // 接口部分仪表盘
            //   path: 'dashboard',
            //   name: 'dashboard',
            //   component: () => import('@/views/InterfaceTestPart/Dashboard/index.vue')
            // },
            // {
            //   // 接口部分用例管理
            //   path: 'caseManage',
            //   name: 'caseManage',
            //   component: () => import('@/views/InterfaceTestPart/CaseManage/index.vue')
            // }
          ]
        }
      ]
    }
    /* {
            // 定时任务部分
            path: '/schedule',
            name: 'schedule',
            component: Layout,
            children: [
                {
                    //
                    path: 'test',
                    name: 'test',
                    component: () => import('@/views/SchedulePart/index.vue'),
                }
            ]
        } */
  ]
})
