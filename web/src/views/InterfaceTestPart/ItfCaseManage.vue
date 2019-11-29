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
        <el-button @click="addModule">新增</el-button>
        <el-tree :props="props"
                 :load="loadNode"
                 :data ='dataList'
                 lazy
                 show-checkbox
                 @node-click="handleNodeClick"
                 @check-change="handleCheckChange">
            <span class="custom-tree-node"
                  slot-scope="{ node, data}">
                {{ data.label }}
                <el-button type="text"
                           v-if="node.level === 1"
                           size="mini"
                           @click="addChildDialogVisible(data)">新增</el-button></span>
        </el-tree>
      </el-scrollbar>
    </el-aside>
    <el-main class="sl-case-list">
      <el-row class="sl-case-handle">
        <div class="sl-handle-in white-back">
          <el-button type="primary"
                     size="mini"
                     @click="caseStepDialogVisible = true">新增</el-button>
        </div>
      </el-row>
      <el-table :data="tableData"
                stripe
                class="sl-case-table white-back">
        <el-table-column prop="stepName"
                         label="步骤名称"
                         width="180">
        </el-table-column>
        <el-table-column prop="reqData"
                         label="请求数据">
        </el-table-column>
        <el-table-column label="操作"
                         min-width="120">
          <template slot-scope="scope">
            <div style="float:left">
              <el-button type="primary"
                         @click.stop="handleChange(scope.row)"
                         size="mini">修改</el-button>
              <el-button type="danger"
                         @click.stop="handleDelete(scope.$index,scope.row)"
                         size="mini">删除</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
  <!-- 弹出框 -->
  <el-dialog title="新增模块"
             :visible.sync="dialogVisible"
             width="30%">
      <el-form ref="form"
               :model="addModuleForm"
               label-width="80px">
        <el-form-item label="模块名称">
          <el-input v-model="addModuleForm.moduleName"></el-input>
        </el-form-item>
         <el-form-item label="模块描述">
          <el-input v-model="addModuleForm.moduleDesc"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addModuleAction">确 定</el-button>
      </span>
    </el-dialog>
  <!-- 用例新增 -->
  <el-dialog title="用例模块"
             :visible.sync="childDialogVisible"
             width="30%">
    <el-form ref="form"
             :model="caseForm"
             label-width="80px">
      <el-form-item label="用例名称">
        <el-input v-model="caseForm.caseName"></el-input>
      </el-form-item>
      <el-form-item label="用例描述">
        <el-input v-model="caseForm.caseDesc"></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
        <el-button @click="childDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addCase">确 定</el-button>
      </span>
  </el-dialog>
  <!-- 步骤新增 -->
  <el-dialog title="用例步骤"
             :visible.sync="caseStepDialogVisible"
             width="30%">
    <el-form ref="form"
             :model="caseStepForm"
             label-width="80px">
      <el-form-item label="步骤名称">
        <el-input v-model="caseStepForm.stepName"></el-input>
      </el-form-item>
      <el-form-item label="接口ID">
        <el-input v-model="caseStepForm.interfaceId"></el-input>
      </el-form-item>
      <el-form-item label="请求体">
        <el-input v-model="caseStepForm.reqData"></el-input>
      </el-form-item>
      <el-form-item label="预期结果">
        <el-input v-model="caseStepForm.expRes"></el-input>
      </el-form-item>
      <el-form-item label="提取变量">
        <el-switch v-model="caseStepForm.isCollectVar"></el-switch>
      </el-form-item>
      <el-row v-if="caseStepForm.isCollectVar">
        <el-row v-for="(item, index) in caseStepForm.variables"
                :key="index">
          <el-col :span="12">
            <el-form-item label="提取字段">
              <el-input v-model="item.collectCol"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="字段别名">
              <el-input v-model="item.collectColAlias"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-row>
<!--      <el-form-item label="用例描述">-->
<!--        <el-input v-model="caseStepForm.caseDesc"></el-input>-->
<!--      </el-form-item>-->

    </el-form>
    <span slot="footer" class="dialog-footer">
        <el-button @click="childDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addCaseStepAction">确 定</el-button>
      </span>
  </el-dialog>
</div>

</template>

<script>
import { addModule, getList, addCase, getCaseList, getCaseTree, getCaseStepList, addCaseStep } from 'api/case.js'
export default {
  name: 'ItfCaseManage',
  data () {
    return {
      dialogVisible: false,
      childDialogVisible: false,
      caseStepDialogVisible: false,
      tableData: [],
      dataList: [],
      addModuleForm: {
        moduleName: '',
        moduleDesc: '',
        projectId: 'd244862701204b0e8467ec5f5a666b32'
      },
      caseForm: {
        'moduleId': '',
        'caseName': '',
        'caseDesc': ''
      },
      caseStepForm: {
        caseId: '3b1adbe45b5842468cea1eb9d8766743',
        stepName: '', // 步骤名称
        interfaceId: '', // 接口ID
        reqData: '', // 请求体
        expRes: '', // 预期响应
        isCollectVar: false, // 是否提取变量
        variables: [
          {
            collectCol: '', // 需要提取的字段，jsonpath提取
            collectColAlias: '' // 提取字段别名
          }
        ]
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
     * @name: addModule
     * @description: 新增接口模块
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addModule () {
      this.dialogVisible = true
    },
    /**
     * @name: addModule
     * @description: 新增模块接口调用
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addModuleAction () {
      addModule(this.addModuleForm).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.dialogVisible = false
      })
    },
    addChildDialogVisible (data) {
      this.childDialogVisible = true
      this.cashData = data
    },
    /**
     * @name: addCase
     * @description: 添加用例
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addCase (data) {
      this.caseForm.moduleId = this.cashData.value
      addCase(this.caseForm).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.childDialogVisible = false
      })
    },
    handleCheckChange (data, checked, indeterminate) {
      console.log(data, checked, indeterminate)
    },
    handleNodeClick (data, node, last) {
      console.log('node', node)
      console.log('last', last)
      if (node.level === 1) { // 是一级节点，获取用例列表
        // let reqInfo = {
        //   moduleId: data.value,
        //   pageNum: 1,
        //   pageSize: 10
        // }
        // getCaseList(reqInfo).then((res) => {
        //   this.tableData = res.list
        // })
        console.log('是一级节点，获取用例列表')
      } else if (node.level === 2) { // 是二级节点，获取用例步骤列表
        let reqInfo = {
          caseId: data.value,
          pageNum: 1,
          pageSize: 10
        }
        getCaseStepList(reqInfo).then((res) => {
          this.tableData = res.list
        })
        console.log('是二级节点，获取用例步骤列表')
      }
    },
    loadNode (node, resolve) {
      if (node.level === 0) {
        getList({ projectId: 1 }).then((res) => {
          return resolve(res)
        })
      } else if (node.level === 1) {
        let reqInfo = {
          moduleId: node.data.value,
          pageNum: 1,
          pageSize: 10
        }
        getCaseTree(reqInfo).then((res) => {
          return resolve(res)
        })
      } else {
        return resolve([])
      }
      // if (node.level > 3) return resolve([])
    },
    getModuleList () {
      // this.$router.projectId
      getList({ projectId: 'd244862701204b0e8467ec5f5a666b32' }).then((res) => {
        console.log(res)
        this.dataList = res
      })
    },
    addCaseStepAction () {
      addCaseStep(this.caseStepForm).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.caseStepDialogVisible = false
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
