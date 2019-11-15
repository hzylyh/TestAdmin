<!--
 * @Description: 接口测试部分首页
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
  <el-row class="sl-itf-main">
    <el-col :span="4"
            class="sl-itf-nav">
      <el-scrollbar style="height: 100%">
        <el-menu default-active="2"
                 router
                 class="sl-itf-nav-menu"
                 background-color="#545c64"
                 text-color="#fff"
                 active-text-color="#ffd04b"
                 style="height: 100%;">
          <el-menu-item index="1"
                        :route="{ name: 'ItfDashboard' }">
            <i class="el-icon-odometer"></i>
            <span slot="title">仪表盘</span>
          </el-menu-item>
          <el-menu-item index="2"
                        :route="{ name: 'ItfCaseManage' }">
            <i class="el-icon-menu"></i>
            <span slot="title">用例管理</span>
          </el-menu-item>
          <el-menu-item index="3">
            <i class="el-icon-document"></i>
            <span slot="title">接口管理</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-setting"></i>
            <span slot="title">变量管理</span>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-col>
    <el-col :span="20" style="height: 100%">
      <router-view/>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'ItfTestIndex'
}
</script>

<style lang="scss" scoped>
  .sl-itf-main {
    height: 100%;
    .sl-itf-nav {
      height: 100%;
      background-color: rgb(84, 92, 100);
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

</style>
