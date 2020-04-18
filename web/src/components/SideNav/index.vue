<template>
  <el-scrollbar style="height: 100%">
    <el-menu :default-active="activeMenu"
             router
             background-color="#1C2B36"
             text-color="white"
             :unique-opened="false"
             active-text-color="#1b867b"
             :collapse-transition="false"
             mode="vertical">
      <side-nav-item v-for="route in routes"
                     :item="route"
                     :key="route.routeName"
                     :route-name="route.routeName"></side-nav-item>
    </el-menu>
  </el-scrollbar>
</template>

<script>
import SideNavItem from './SideNavItem'
export default {
  name: 'index',
  components: {
    SideNavItem
  },
  data () {
    return {
      activeMenu: '1',
      routes: [
        {
          code: '1',
          logo: 'el-icon-odometer',
          name: '仪表盘',
          routeName: 'ItfDashboard',
          children: []
        },
        {
          code: '2',
          logo: 'el-icon-menu',
          name: '接口管理',
          routeName: 'apiList'
        },
        {
          code: '3',
          logo: 'el-icon-document',
          name: '用例管理',
          routeName: 'ItfCaseManage'
        }
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
.el-menu {
  border-right: none;
  text-align: left;
}
</style>
