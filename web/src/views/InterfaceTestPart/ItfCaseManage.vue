<!--
 * @Description: 接口测试部分用例管理
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
<div class="sl-case-main">
 <el-container >
    <el-aside
            class="sl-suit-tree white-back">
      <el-scrollbar>
        <el-input></el-input>
        <el-button @click="addMoudel">新增</el-button>
        <el-tree
                :props="props"
                :load="loadNode"
                :data ='dataList'
                show-checkbox
                @check-change="handleCheckChange">
<span class="custom-tree-node"
              slot-scope="{ data}">
{{ data.moduleName }} <el-button type="primary"
                     size="mini" @click="addChidlDialogVisible(data)">新增</el-button></span>
        </el-tree>
      </el-scrollbar>
    </el-aside>
    <el-main
            class="sl-case-list">
      <el-row class="sl-case-handle">
        <div class="sl-handle-in white-back">
          <el-button type="primary"
                     size="mini">新增</el-button>
        </div>
      </el-row>
      <el-table :data="tableData"
                stripe
                class="sl-case-table white-back">
        <el-table-column
                prop="date"
                label="日期"
                width="180">
        </el-table-column>
        <el-table-column
                prop="name"
                label="姓名"
                width="180">
        </el-table-column>
        <el-table-column
                prop="address"
                label="地址">
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
  <!-- 弹出框 -->
  <el-dialog
      title="新增模块"
      :visible.sync="dialogVisible"
      width="30%"
      >
      <el-form ref="form"
               :model="caseForm"
               label-width="80px">
        <el-form-item label="模块名称">
          <el-input v-model="caseForm.caseName"></el-input>
        </el-form-item>
         <el-form-item label="模块描述">
          <el-input v-model="caseForm.caseDesc"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addChildModule">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 子模块新增 -->
    <el-dialog
      title="新增模块"
      :visible.sync="chidlDialogVisible"
      width="30%"
      >
      <el-form ref="form"
               :model="addModuleform"
               label-width="80px">
        <el-form-item label="模块名称">
          <el-input v-model="addModuleform.moduleName"></el-input>
        </el-form-item>
         <el-form-item label="模块描述">
          <el-input v-model="addModuleform.moduleDesc"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addChildModule">确 定</el-button>
      </span>
    </el-dialog>
</div>

</template>

<script>
import { add, getList, addModule } from 'api/case.js'
export default {
  name: 'ItfCaseManage',
  data () {
    return {
      dialogVisible: false,
      chidlDialogVisible: false,
      tableData: [],
      dataList: [],
      addModuleform: {
        moduleName: '',
        moduleDesc: '',
        projectId: 1
      },
      caseForm: {
        'moduleId': '',
        'caseName': 'test',
        'caseDesc': 'test'
      },
      props: {
        label: 'name',
        children: 'zones'
      },
      count: 1
    }
  },
  methods: {
    /**
     * @name: addMoudel
     * @description: 新增接口模块
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addMoudel () {
      this.dialogVisible = true
    },
    /**
     * @name: addModuleAction
     * @description: 新增模块接口调用
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addModuleAction () {
      add(this.addModuleform).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.dialogVisible = false
      })
    },
    addChidlDialogVisible (data) {
      this.chidlDialogVisible = true
      this.cashData = data
    },
    /**
     * @name: addChildModule
     * @description: 添加子模块
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addChildModule (data) {
      this.caseForm.moduleId = this.cashData.id
      addModule(this.caseForm).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.chidlDialogVisible = false
      })
    },
    handleCheckChange (data, checked, indeterminate) {
      console.log(data, checked, indeterminate)
    },
    handleNodeClick (data) {
      console.log(data)
    },
    loadNode (node, resolve) {
      if (node.level === 0) {
        getList({ projectId: 1 }).then((res) => {
          return resolve(res)
        })
      }
      if (node.level > 3) return resolve([])
    },
    getModuleList () {
      // this.$router.projectId
      getList({ projectId: 1 }).then((res) => {
        console.log(res)
        this.dataList = res
      })
    }
  },
  created () {
    this.getModuleList()
  }
}
</script>

<style lang="scss" scoped>
.sl-case-main {
  height: 100%;
  padding: 5px;
  /*background: rgb(240, 242, 245);*/
  .sl-suit-tree {
    height: 100%;
  }
  .sl-case-list {
    height: 100%;
    padding-left: 5px;
    .sl-case-handle {
      text-align: right;
      padding-bottom: 8px;
      height: 56px;
      .sl-handle-in {
        padding: 10px;
      }
    }
    .sl-case-table {
      width: 100%;
      height: calc(100% - 56px);
    }
  }
}
.white-back {
  background: #fff;
}
</style>
