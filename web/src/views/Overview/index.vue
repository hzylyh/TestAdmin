<template>
  <div class="overView">
    <!-- 头部说明部分 -->
    <div class="overViewHead">
      一个轻量级的企业测试平台
    </div>
    <el-tabs type="border-card"
             class="tabcontent">
      <el-tab-pane class="full">
        <span slot="label"><i class="el-icon-folder-opened"></i>我的项目</span>
        <!-- 项目内容列表 -->
        <ul class="project-list">
          <li class="child-item"
              @click="addPojectClick">
            <i class="el-icon-plus addicon"></i>
          </li>
          <li v-for="(item,index) in projectList"
              :key="index"
              class="child-item"
              @click="goToProject(item)">
<!--            <div v-if="!item.isAdd">-->
            <div>
              <div>{{item.projectName}}</div>
              <div>{{item.projetcDesc}}</div>
            </div>
<!--            <i v-else-->
<!--               class="el-icon-plus addicon"-->
<!--               @click.stop="addPojectClick"></i>-->
          </li>
        </ul>
      </el-tab-pane>
      <el-tab-pane class="full"
                   label="我的收藏"> <span slot="label"><i class="el-icon-star-off"></i>我的收藏</span>
        <undeveloped></undeveloped>
      </el-tab-pane>
    </el-tabs>
    <el-dialog title="新建项目"
               :visible.sync="
               dialogVisible"
               width="30%">
      <el-form ref="form"
               :model="form"
               label-width="80px">
        <el-form-item label="项目名称">
          <el-input v-model="form.projectName"></el-input>
        </el-form-item>
        <el-form-item label="项目描述">
          <el-input v-model="form.projetcDesc"
                    type="textarea"></el-input>
        </el-form-item>
        <!-- <el-form-item label="图片">
          <el-upload action="https://jsonplaceholder.typicode.com/posts/"
                     :show-file-list="false"
                     :on-success="handleAvatarSuccess"
                     :before-upload="beforeAvatarUpload">
            <img v-if="imageUrl"
                 :src="imageUrl"
                 class="avatar">
            <i v-else
               class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item> -->
      </el-form>
      <span slot="footer"
            class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary"
                   @click="addPojectAction">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import undeveloped from '@/components/Dev/undeveloped.vue'
import { addProject, getProjectList } from '@/api/overView.js'
export default {
  name: 'index',
  components: {
    undeveloped
  },
  data () {
    return {
      projectList: [],
      dialogVisible: false,
      form: {
        projectName: '',
        projetcDesc: ''
      }
    }
  },
  methods: {
    addPojectClick () {
      this.dialogVisible = true
    },
    /**
     * @name: addPojectAction
     * @description: 新增项目
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addPojectAction () {
      addProject(this.form).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.dialogVisible = false
        this.$router.push({
          name: 'apiList'
        })
      })
    },
    queryProjectList () {
      var self = this
      getProjectList().then((res) => {
        // res.unshift({
        //   isAdd: true
        // })
        self.projectList = res
      })
    },
    /**
     * @name: goToProject
     * @description: 去往项目
     * @param {type}: 项目信息
     * @return {type}: 默认类型
     */
    goToProject (projectItem) {
      // 存储项目ID，方便后续请求使用
      localStorage.setItem('projectId', projectItem.projectId)
      this.$router.push({
        name: 'ItfDashboard'
      })
    }
  },
  created () {
    this.queryProjectList()
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.overView {
  height: 100%;
  // padding-top: 40px;
  box-sizing: border-box;
  // background: white;
  .overViewHead {
    height: 200px;
    font-size: 34px;
    font-weight: 200;
    line-height: 200px;
    color: white;
    background: url("../../assets/img/overViewHeadBg.jpg") no-repeat;
    background-size: cover;
  }
  .tabcontent {
    height: calc(100% - 200px);
    box-sizing: border-box;
  }
  .full {
    height: 100%;
    box-sizing: border-box;
  }
  /deep/.el-tabs__content {
    height: calc(100% - 39px);
    box-sizing: border-box;
    padding: 0px;
    overflow-y: auto;
  }
  /deep/.content {
    background-color: white;
  }
  .project-list {
    // height: 100%;
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
      // background: green;
      border-radius: 10px;
      border: 1px solid #dcdfe6;
      box-sizing: border-box;
      box-shadow: 0 5px 5px 0 rgba(0, 0, 0, 0.25);
      cursor: pointer;
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .child-item:hover {
      border: 1px solid #409eff;
    }
    .child-item:nth-child(4n) {
      // margin-right: 0px;
    }
  }
  // .avatar-uploader .el-upload {
  //   border: 1px dashed #d9d9d9;
  //   border-radius: 6px;
  //   cursor: pointer;
  //   position: relative;
  //   overflow: hidden;
  // }
  // .avatar-uploader .el-upload:hover {
  //   border-color: #409eff;
  // }
  // .avatar-uploader-icon {
  //   font-size: 28px;
  //   color: #8c939d;
  //   width: 178px;
  //   height: 178px;
  //   line-height: 178px;
  //   text-align: center;
  // }
  // .avatar {
  //   width: 178px;
  //   height: 178px;
  //   display: block;
  // }
}
</style>
