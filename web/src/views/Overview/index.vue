<template>
  <div class="overView">
    <el-tabs type="border-card"
             class="full">
      <el-tab-pane class="full">
        <span slot="label"><i class="el-icon-folder-opened"></i>项目</span>
        <!-- 项目内容列表 -->
        <ul class="project-list">
          <li v-for="(item,index) in projectList"
              :key="index"
              class="child-item">
            <i class="el-icon-plus addicon"></i>
          </li>
        </ul>
      </el-tab-pane>
      <el-tab-pane class="full"
                   label="我的收藏"> <span slot="label"><i class="el-icon-star-off"></i>我的收藏</span>
        <undeveloped></undeveloped>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import undeveloped from '@/components/Dev/undeveloped.vue'
export default {
  name: 'index',
  components: {
    undeveloped
  },
  data () {
    return {
      projectList: [1, 2, 3, 4, 4, 4, 4, 44, 4, 44, 5]
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.overView {
  height: 100%;
  padding-top: 40px;
  box-sizing: border-box;
  background: white;
  .full {
    height: 100%;
    box-sizing: border-box;
  }
  /deep/.el-tabs__content {
    height: calc(100% - 39px);
    box-sizing: border-box;
    padding: 0px;
  }
  /deep/.content {
    background-color: white;
  }
  .project-list {
    height: 100%;
    // background: red;
    padding: 40px;
    box-sizing: border-box;
    display: flex;
    flex-wrap: wrap;
    overflow-y: auto;
    .child-item {
      margin: 10px;
      width: calc(25% - 20px);
      // width: 25px;
      height: 200px;
      background: green;
      border-radius: 10px;
      cursor: pointer;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}
</style>
