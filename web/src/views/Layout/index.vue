<template>
  <el-row class="sl-main">
    <el-row class="sl-header">
      <el-row>
        <el-col :span="4"
                class="sl-logo">
          <div>logoArea</div>
        </el-col>
        <el-col :span="20">
          <el-menu
                  :default-active="activeIndex2"
                  mode="horizontal"
                  @select="handleSelect"
                  background-color="#545c64"
                  text-color="#fff"
                  router
                  active-text-color="#ffd04b">
            <el-menu-item index="1" :route="{name: 'OverviewIndex'}">总览</el-menu-item>
            <el-menu-item index="2" :route="{name: 'ItfDashboard'}">接口自动化</el-menu-item>
            <el-menu-item index="3">性能测试</el-menu-item>
            <el-menu-item index="4">云真机</el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
    </el-row>
    <el-row class="sl-each-part">
      <router-view/>
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
      line-height: 61px;
      margin: 0;
      background-color: rgb(84, 92, 100);
    }
  }
  .sl-each-part {
    padding: 0px;
    height: calc(100% - 61px);
  }
}
</style>
