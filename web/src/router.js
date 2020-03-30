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
    {
      path: '/',
      name: 'Login',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login')
    },
    {
      path: '/overview',
      name: 'Overview',
      component: Layout,
      children: [
        {
          // 总览首页
          path: 'index',
          name: 'OverviewIndex',
          component: () => import('@/views/Overview')
        }
      ]
    },
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
            {
              // 接口部分仪表盘
              path: 'dashboard',
              name: 'ItfDashboard',
              component: () =>
                import('@/views/InterfaceTestPart/ItfDashboard.vue')
            },
            {
              // 接口部分用例管理
              path: 'caseManage',
              name: 'ItfCaseManage',
              component: () =>
                import('@/views/InterfaceTestPart/ItfCaseManage.vue')
            },
            {
              // 接口管理
              path: 'apiList',
              name: 'apiList',
              component: () =>
                import('@/views/InterfaceTestPart/ItfInterfaceManage/index.vue')
            }
          ]
        }
      ]
    },
    {
      path: '/performanceTesting',
      name: 'performanceTesting',
      redirect: '/performanceTesting/index',
      component: Layout,
      children: [
        {
          // 总览首页
          path: 'index',
          name: 'performanceTestingIndex',
          component: () => import('@/components/Dev/undeveloped.vue')
        }
      ]
    },
    {
      path: '/realMachine',
      name: 'realMachine',
      redirect: '/realMachine/index',
      component: Layout,
      children: [
        {
          // 总览首页
          path: 'index',
          name: 'realMachineIndex',
          component: () => import('@/components/Dev/undeveloped.vue')
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
