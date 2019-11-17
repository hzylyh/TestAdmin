<template>
  <el-row class="sl-main">
    <div class="header">
      <div style="width:200px"
           class="sl-logo">
        <span class="first">Test</span>
        Admin
      </div>
      <div>
        <el-menu :default-active="activeIndex2"
                 mode="horizontal"
                 @select="handleSelect"
                 background-color="#1C2B36"
                 text-color="white"
                 router
                 active-text-color="#1b867b">
          <el-menu-item index="1"
                        :route="{name: 'OverviewIndex'}">总览</el-menu-item>
          <!-- <el-menu-item index="2" :route="{name: 'ItfDashboard'}">接口自动化</el-menu-item> -->
          <el-menu-item index="2"
                        :route="{name: 'performanceTestingIndex'}">性能测试</el-menu-item>
          <el-menu-item index="4"
                        :route="{name: 'realMachineIndex'}">云真机</el-menu-item>
        </el-menu>
      </div>
    </div>
    <el-row class="sl-each-part">
      <router-view />
    </el-row>
  </el-row>
</template>

<script>
export default {
  name: 'index',
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect () {
      console.log('select')
    }
  }
}
</script>

<style lang="scss" scoped>
.sl-main {
  height: 100%;
  /*background: rgb(240, 242, 245);*/
  .sl-header {
    padding: 0px;
    .sl-logo {
      float: left;
      height: 61px;
      line-height: 60px;
      color: white;
      margin: 0;
      // background-color: #1c2b36;
      cursor: pointer;
    }
  }
  .sl-each-part {
    padding: 0px;
    height: calc(100% - 60px);
    overflow-y: auto;
    background: rgb(240, 243, 244);
  }
  .header {
    display: flex;
    background: #1c2b36;
    height: 60px;
    .sl-logo {
      color: #fff;
      display: block;
      font-size: 24px;
      font-family: Roboto, sans-serif;
      white-space: nowrap;
      float: left;
      line-height: 60px;
      .first {
        color: #209e91;
      }
    }
    .sl-logo:hover {
      color: #1b867b;
    }
  }
}
</style>
