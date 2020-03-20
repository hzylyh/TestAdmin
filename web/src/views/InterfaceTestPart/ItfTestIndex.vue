<!--
 * @Description: 接口测试部分首页
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
  <el-container class="sl-itf-main">
    <el-aside width="200px"
              class="sl-itf-nav">
      <el-scrollbar style="height: 100%">
        <el-menu default-active="activeIndex"
                 router
                 class="sl-itf-nav-menu"
                 background-color="#1C2B36"
                 text-color="white"
                 active-text-color="#1b867b"
                 style="height: 100%;">
          <el-menu-item index="1"
                        :route="{ name: 'ItfDashboard' }">
            <i class="el-icon-odometer"></i>
            <span slot="title"
                  class="title">仪表盘</span>
          </el-menu-item>
          <el-menu-item index="2"
                        :route="{ name: 'apiList' }">
            <i class="el-icon-document"></i>
            <span slot="title"
                  class="title">接口管理</span>
          </el-menu-item>
          <el-menu-item index="3"
                        :route="{ name: 'ItfCaseManage' }">
            <i class="el-icon-menu"></i>
            <span slot="title"
                  class="title">用例管理</span>
          </el-menu-item>

          <!-- <el-menu-item index="4">
            <i class="el-icon-setting"></i>
            <span slot="title"
                  class="title">变量管理</span>
          </el-menu-item> -->
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-main :span="20"
             style="height: 100%;padding:15px">
      <router-view />
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'ItfTestIndex',
  data () {
    return {
      activeIndex: '1'
    }
  }
}
</script>

<style lang="scss" scoped>
.sl-itf-main {
  height: 100%;
  .sl-itf-nav {
    height: 100%;
    background-color: #1c2b36;
    .sl-itf-nav-menu {
      height: 100%;
      border-right: none;
    }
    /deep/ .el-scrollbar__wrap {
      height: 100% !important;
      overflow-x: auto !important;
    }
  }
}
.title {
  width: 56px;
  text-align: left;
}
</style>
