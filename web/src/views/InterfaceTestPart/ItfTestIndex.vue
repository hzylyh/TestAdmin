<template>
  <el-row class="sl-itf-main">
    <el-col :span="4"
            class="sl-itf-nav">
      <el-scrollbar style="height: 100%">
        <el-menu
                default-active="2"
                class="sl-itf-nav-menu"
                @open="handleOpen"
                @close="handleClose"
                background-color="#545c64"
                text-color="#fff"
                active-text-color="#ffd04b"
                style="height: 100%;">
          <el-submenu index="1">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>导航一</span>
            </template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="1-1">选项1</el-menu-item>
              <el-menu-item index="1-2">选项2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="1-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="1-4">
              <template slot="title">选项4</template>
              <el-menu-item index="1-4-1">选项1</el-menu-item>
            </el-submenu>
          </el-submenu>
          <el-menu-item index="2">
            <i class="el-icon-menu"></i>
            <span slot="title">导航二</span>
          </el-menu-item>
          <el-menu-item index="3" disabled>
            <i class="el-icon-document"></i>
            <span slot="title">导航三</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-setting"></i>
            <span slot="title">导航四</span>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-col>
    <el-col :span="20">
      <router-view />
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
    /deep/.el-scrollbar__wrap {
      height: 100% !important;
      overflow-x: auto !important;
    }
  }
}

</style>
