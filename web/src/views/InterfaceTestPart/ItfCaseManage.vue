<!--
 * @Description: 接口测试部分用例管理
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
<div class="sl-case-main">
 <el-container style="height: 100%">
    <el-aside class="sl-suit-tree white-back">
      <el-scrollbar style="height: 100%">
<!--        <el-input></el-input>-->
        <el-button-group class="sl-btn-group">
          <el-button @click="addModule"
                     type="primary"
                     size="mini">新增</el-button>
          <el-button @click="runCase"
                     type="primary"
                     size="mini">运行</el-button>
        </el-button-group>

        <el-tree :props="props"
                 ref="caseTree"
                 :load="loadNode"
                 :data ='dataList'
                 lazy
                 show-checkbox
                 node-key="id"
                 :default-expanded-keys="[1]"
                 @node-click="handleNodeClick"
                 @check-change="handleCheckChange">
            <span class="tree-line"
                  slot-scope="{ node, data}">
              <el-badge is-dot class="item">
                {{ data.label }}
              </el-badge>
              <el-button-group>
                <el-button type="text"
                           class="add-btn"
                           v-if="node.level === 2"
                           size="mini"
                           @click.stop="addChildDialogVisible(data)">新增</el-button>
              <el-button type="text"
                         class="add-btn"
                         v-if="node.level === 2"
                         size="mini"
                         @click.stop="addChildDialogVisible(data)">删除</el-button>
              </el-button-group>
            </span>
        </el-tree>
      </el-scrollbar>
    </el-aside>
    <el-main class="sl-case-list">
      <el-row class="sl-case-handle">
        <div class="sl-handle-in white-back">
          <el-button type="primary"
                     size="mini"
                     @click="addCaseStep">新增</el-button>
        </div>
      </el-row>
      <el-row class="sl-case-table white-back">
        <el-table :data="tableData"
                  stripe
                  class="white-back">
          <el-table-column type="index"
                           align="center"
                           label="序号">
          </el-table-column>
          <el-table-column prop="stepName"
                           label="步骤名称">
          </el-table-column>
          <el-table-column prop="stepDesc"
                           show-overflow-tooltip
                           label="步骤描述">
          </el-table-column>
          <el-table-column show-overflow-tooltip
                           label="执行结果">
            <template slot-scope="scope">
              <el-tag :type="scope.row.stepStatus | statusFilter"
                      style="cursor: pointer"
                      effect="plain"
                      @click="showStepResDialog"
                      disable-transitions>{{ scope.row.stepStatus }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作"
                           width="350px">
            <template slot-scope="scope">
              <div style="float:left">
                <el-button type="primary"
                           @click.stop="handleChange(scope.row.caseId)"
                           icon="el-icon-s-promotion"
                           size="mini">运行</el-button>
                <el-button type="primary"
                           @click.stop="handleChange(scope.row.caseId)"
                           icon="el-icon-document-copy"
                           size="mini">复制</el-button>
                <el-button type="primary"
                           @click.stop="editCaseStep(scope.row)"
                           icon="el-icon-edit"
                           size="mini">修改</el-button>
                <el-button type="danger"
                           @click.stop="handleCaseStepDel(scope.row)"
                           icon="el-icon-delete"
                           size="mini">删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
        <div style="position: absolute; bottom: 5px; right: 5px">
          <el-pagination
                  @size-change="handleSizeChange"
                  @current-change="handleCurrentChange"
                  :current-page="currentPage4"
                  background
                  :page-size="pageSize"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="total">
          </el-pagination>
        </div>
      </el-row>

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
             width="40%">
    <el-form ref="form"
             :model="caseStepForm"
             label-width="80px">
      <el-row>
        <el-tabs tab-position="top" style="height: 100%;">
          <el-tab-pane label="基本信息">
            <el-row>
              <el-col :span="12">
                <el-form-item label="步骤名称">
                  <el-input v-model="caseStepForm.stepName"></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="18">
                <el-form-item label="步骤描述">
                  <el-input v-model="caseStepForm.stepName"
                            type="textarea"></el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-tab-pane>
          <el-tab-pane label="请求信息">
            <el-row>

              <el-col :span="12">
                <el-form-item label="接口">
                  <el-select v-model="caseStepForm.interfaceId"
                             @change="changeInterface"
                             placeholder="请选择">
                    <el-option
                            v-for="item in interfaceOptions"
                            :key="item.value"
                            :label="item.name"
                            :value="item.interfaceId">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="24">
                <el-form-item label="请求体">
                  <el-input v-model="caseStepForm.reqData"></el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-tab-pane>
          <el-tab-pane label="接口断言">
            <el-row :gutter="10">
              <el-col :span="10">
                比对字段
              </el-col>
              <el-col :span="4">
                关系
              </el-col>
              <el-col :span="10">
                期望
              </el-col>
            </el-row>
            <el-row v-for="(item, index) in caseStepForm.assertInfos"
                    :gutter="10"
                    :key="index">
              <el-col :span="10">
                <el-input v-model="item.assertCol"></el-input>
              </el-col>
              <el-col :span="4">
                <el-input v-model="item.method"></el-input>
              </el-col>
              <el-col :span="10">
                <el-input v-model="item.expValue"></el-input>
              </el-col>
            </el-row>
          </el-tab-pane>
          <el-tab-pane label="变量提取">
            <el-row>
              <el-col :span="6">
                <el-form-item label="提取变量">
                  <el-switch v-model="caseStepForm.isCollectVar"></el-switch>
                </el-form-item>
              </el-col>
            </el-row>
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
          </el-tab-pane>
        </el-tabs>
      </el-row>
<!--      <el-form-item label="用例描述">-->
<!--        <el-input v-model="caseStepForm.caseDesc"></el-input>-->
<!--      </el-form-item>-->

    </el-form>
    <span slot="footer" class="dialog-footer">
        <el-button @click="childDialogVisible = false">取 消</el-button>
        <el-button v-if="actionFlag === 'add'" type="primary" @click="addCaseStepAction">确 定</el-button>
        <el-button v-if="actionFlag === 'edit'" type="primary" @click="editCaseStepAction">确 定</el-button>
    </span>
  </el-dialog>

  <el-dialog :visible="stepResDialogVisible">
    <json-format v-if="stepResDialogVisible" :json-cont="testObj" :col="testCol"></json-format>
    <span slot="footer" class="dialog-footer">
      <el-button @click="stepResDialogVisible = false">取 消</el-button>
    </span>
  </el-dialog>
</div>

</template>

<script>
import {
  addModule,
  getList,
  addCase,
  getCaseList,
  getCaseTree,
  getCaseStepList,
  addCaseStep,
  runCase,
  getItfSelectOptions,
  editCaseStep,
  delCaseStep,
  getCaseStepDetail
} from 'api/case.js'
import JsonFormat from '@/components/JsonFormat'
export default {
  name: 'ItfCaseManage',
  components: { JsonFormat },
  data () {
    return {
      caseId: '',
      actionFlag: '',
      dialogVisible: false,
      childDialogVisible: false,
      caseStepDialogVisible: false,
      stepResDialogVisible: false,
      tableData: [],
      dataList: [],
      pageNum: 1,
      pageSize: 10,
      total: 0,
      interfaceOptions: [],
      testCol: 'b.d',
      testObj: {
        a: 'd',
        b: {
          d: 'd'
        }
      },
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
        caseId: '',
        stepId: '',
        stepName: '', // 步骤名称
        stepDesc: '', // 步骤描述
        interfaceId: '', // 接口ID
        reqData: '', // 请求体
        expRes: '', // 预期响应
        assertInfos: [
          {
            assertCol: '',
            method: '',
            expValue: ''
          }
        ], // 断言
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
        children: 'zones',
        isLeaf: 'leaf'
      },
      count: 1
    }
  },
  filters: {
    statusFilter (val) {
      let statusMap = {
        '成功': 'success',
        '失败': 'danger'
      }
      return statusMap[val]
    }
  },
  created () {
    this.getModuleList()
  },
  mounted () {
    this.getItfSelectOptions()
  },
  methods: {
    editCaseStep (row) {
      let reqInfo = {
        stepId: row.stepId
      }
      getCaseStepDetail(reqInfo).then(response => {
        this.caseStepForm = response
        if (response.assertInfos.length === 0) {
          this.caseStepForm.assertInfos = [{
            assertCol: '',
            method: '',
            expValue: ''
          }]
        }
        if (response.variables.length === 0) {
          this.caseStepForm.variables = [{
            collectCol: '', // 需要提取的字段，jsonpath提取
            collectColAlias: '' // 提取字段别名
          }]
        }
        this.actionFlag = 'edit'
        this.caseStepDialogVisible = true
      })
    },
    editCaseStepAction () {
      editCaseStep(this.caseStepForm).then(response => {
        this.caseStepDialogVisible = false
        this.getCaseStepList(this.caseId)
      })
    },
    getCaseStepList (caseId) {
      let reqInfo = {
        caseId: caseId,
        pageNum: 1,
        pageSize: 10
      }
      getCaseStepList(reqInfo).then((res) => {
        this.tableData = res.list
        this.pageNum = res.pageNum
        this.pageSize = res.pageSize
        this.total = res.total
      })
    },
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
    runCase () {
      // 目前逻辑有点问题，模块用例运行顺序可能存在问题
      let caseMap = {}
      let currNode = null
      let regionList = this.$refs.caseTree.getCheckedNodes(false, true)
      for (let item of regionList) {
        currNode = this.$refs.caseTree.getNode(item)
        if (currNode.level !== 1 && currNode.isLeaf === false) {
          caseMap[currNode.data.value] = []
        } else if (currNode.isLeaf === true) {
          caseMap[currNode.parent.data.value].push(currNode.data.value)
        }
      }
      // console.log(caseMap)
      // console.log(currNode)
      // console.log(this.caseList)
      let reqInfo = {
        'cases': caseMap
      }
      runCase(reqInfo).then(res => {
        this.$message({
          message: '运行成功',
          type: 'success'
        })
      })
    },
    handleCheckChange (data, checked, indeterminate) {
      console.log(data, checked, indeterminate)

    },
    handleNodeClick (data, node, last) {
      console.log(data.value)
      if (node.level === 1) { // 是一级节点，获取用例列表
        // let reqInfo = {
        //   moduleId: data.value,
        //   pageNum: 1,
        //   pageSize: 10
        // }
        // getCaseList(reqInfo).then((res) => {
        //   this.tableData = res.list
        // })
      } else if (node.level === 3) { // 是二级节点，获取用例步骤列表
        this.caseId = data.value
        this.getCaseStepList(data.value)
      }
    },
    loadNode (node, resolve) {
      if (node.level === 1) {
        getList({ projectId: 'd244862701204b0e8467ec5f5a666b32' }).then((res) => {
          return resolve(res)
        })
      } else if (node.level === 2) {
        let reqInfo = {
          moduleId: node.data.value,
          pageNum: 1,
          pageSize: 10
        }
        getCaseTree(reqInfo).then((res) => {
          console.log(res.length)
          if (res !== null && res.length !== undefined) {
            for (let index in res) {
              res[index].leaf = true
            }
            return resolve(res)
          } else {
            return resolve([])
          }
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
        // this.dataList = res
        this.dataList = [{
          id: 1,
          label: '项目名',
          children: res
        }]
      })
    },
    addCaseStep () {
      this.caseStepDialogVisible = true
      this.actionFlag = 'add'
    },
    addCaseStepAction () {
      this.caseStepForm.caseId = this.caseId
      addCaseStep(this.caseStepForm).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.caseStepDialogVisible = false
        this.getCaseStepList(this.caseId)
      })
    },
    getItfSelectOptions () {
      getItfSelectOptions().then(res => {
        this.interfaceOptions = res
      })
    },
    /**
     * @name: handleStepChange
     * @description: 修改用例步骤
     * @param {type}: 用例步骤ID
     * @return {type}: 默认类型
     */
    handleCaseStepChange (caseStepInfo) {
      let reqInfo = {
        'caseStepId': caseStepInfo.stepId
      }
      editCaseStep(reqInfo).then(response => {
        this.getCaseStepList(caseStepInfo.caseId)
      })
    },
    /**
     * @name: handleStepChange
     * @description: 删除用例步骤
     * @param {type}: 用例步骤ID
     * @return {type}: 默认类型
     */
    handleCaseStepDel (caseStepInfo) {
      let reqInfo = {
        'caseStepId': caseStepInfo.stepId
      }
      delCaseStep(reqInfo).then(response => {
        this.getCaseStepList(caseStepInfo.caseId)
      })
    },
    showStepResDialog () {
      this.stepResDialogVisible = true
    }
  }
}
</script>

<style lang="scss" scoped>
.sl-case-main {
  height: 100%;
  /*padding: 5px;*/
  /*background: rgb(240, 242, 245);*/
  .sl-suit-tree {
    height: 100%;
    .sl-btn-group {
      margin-top: 15px;
      margin-bottom: 10px;
    }
    .tree-line {
      flex: 1;
      display: flex;
      align-items: center;
      justify-content: space-between;
      /*justify-content: space-around;*/
      font-size: 14px;
      padding-right: 18px;
      .add-btn {
        /*display: none;*/
        padding-right: 10px;
      }
    }
    .tree-line:hover {
      .add-btn {
        display: inline;
      }
    }
  }
  .sl-case-list {
    height: 100%;
    padding: 0px 5px 0px 5px;
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
