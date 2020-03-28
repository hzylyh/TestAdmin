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
        <div style="text-align: right; margin-right: 10px">
          <el-button-group class="sl-btn-group">
            <!--          <el-button @click="addModule"-->
            <!--                     type="primary"-->
            <!--                     size="mini">新增</el-button>-->
            <el-button @click="runCaseNew"
                       type="primary"
                       size="mini">运行</el-button>
          </el-button-group>
        </div>


        <el-tree ref="caseTree"
                 :data ='dataList'
                 show-checkbox
                 node-key="id"
                 @node-click="handleNodeClick"
                 @check-change="handleCheckChange">
            <span class="tree-line"
                  slot-scope="{ node, data}">
              <div>
                <i v-if="data.nodeType === '模块'" class="el-icon-folder"></i>
                <i v-if="data.nodeType === '用例'" class="el-icon-document"></i>
              {{ data.nodeName }}
              </div>

              <el-button-group>
<!--                <el-button type="text"-->
<!--                           class="add-btn"-->
<!--                           v-if="node.level === 1"-->
<!--                           size="mini"-->
<!--                           @click.stop="addModule">新增</el-button>-->
<!--                <el-button type="text"-->
<!--                           class="add-btn"-->
<!--                           v-if="node.level === 2"-->
<!--                           size="mini"-->
<!--                           @click.stop="addChildDialogVisible(data)">新增</el-button>-->
<!--                <el-button type="text"-->
<!--                         class="add-btn"-->
<!--                         v-if="node.level === 2"-->
<!--                         size="mini"-->
<!--                         @click.stop="delModule(data)">删除</el-button>-->
<!--                <el-button type="text"-->
<!--                           class="add-btn"-->
<!--                           v-if="node.level === 3"-->
<!--                           size="mini"-->
<!--                           @click.stop="showEditCase(data)">编辑</el-button>-->
<!--                <el-popover-->
<!--                        placement="top"-->
<!--                        width="160"-->
<!--                        v-model="visible">-->
<!--                  <p>这是一段内容这是一段内容确定删除吗？</p>-->
<!--                  <div style="text-align: right; margin: 0">-->
<!--                    <el-button size="mini" type="text" @click="visible = false">取消</el-button>-->
<!--                    <el-button type="primary" size="mini" @click="visible = false">确定</el-button>-->
<!--                  </div>-->
<!--                </el-popover>-->
<!--                <el-button type="text"-->
<!--                           class="add-btn"-->
<!--                           v-if="node.level === 3"-->
<!--                           size="mini"-->
<!--                           @click.stop="delCase(data)">删除</el-button>-->
                <el-button type="text"
                           class="add-btn"
                           size="mini"
                           @click.stop="addNode(data)">新增</el-button>
                <el-button type="text"
                           class="add-btn"
                           v-if="node.level !== 1"
                           size="mini"
                           @click.stop="editNode(data)">编辑</el-button>
                <el-button type="text"
                           class="add-btn"
                           v-if="node.level !== 1"
                           size="mini"
                           @click.stop="delNodeAction(data)">删除</el-button>
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
              <el-tag v-if="scope.row.stepStatus !== '成功'"
                      :type="scope.row.stepStatus | statusFilter"
                      style="cursor: pointer"
                      effect="plain"
                      @click="showStepResDialog(scope.row)"
                      disable-transitions>{{ scope.row.stepStatus }}</el-tag>
              <el-tag v-else
                      :type="scope.row.stepStatus | statusFilter"
                      effect="plain"
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
                  :current-page="pageNum"
                  background
                  :page-size="pageSize"
                  layout="total, sizes, prev, pager, next, jumper"
                  :total="total">
          </el-pagination>
        </div>
      </el-row>

    </el-main>
  </el-container>

    <el-dialog title="新增"
               :visible.sync="nodeDialogVisible"
               width="30%">
      <el-form ref="nodeForm"
               :rules="nodeFormRule"
               :model="nodeForm"
               label-width="80px">
        <el-form-item label="类型"
                      style="text-align: left"
                      prop="nodeType">
          <el-select v-model="nodeForm.nodeType" placeholder="请选择">
            <el-option
                    v-for="item in nodeOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="名称"
                      prop="nodeName">
          <el-input v-model="nodeForm.nodeName"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="nodeForm.nodeDesc"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button @click="nodeDialogVisible = false">取 消</el-button>
          <el-button type="primary"
                     v-if="nodeActionFlag === 'add'"
                     @click="addNodeAction">确 定</el-button>
          <el-button type="primary"
                     v-if="nodeActionFlag === 'edit'"
                     @click="editNodeAction">确 定</el-button>
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
            <el-row>
              <el-col :span="22">
                <el-col :span="10">
                  比对字段
                </el-col>
                <el-col :span="4">
                  关系
                </el-col>
                <el-col :span="10">
                  期望
                </el-col>
              </el-col>
            </el-row>
            <el-row v-for="(item, index) in caseStepForm.assertInfos"
                    style="margin-bottom: 5px"
                    :key="index">
              <el-col :span="22">
                <el-row :gutter="10">
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
              </el-col>
              <el-col :span="1">
                <i @click="assertAction(index,1)" v-if="index === 0" class="add" />
                <i @click="assertAction(index,0)" v-else class="reduce" />
              </el-col>
            </el-row>
          </el-tab-pane>
          <el-tab-pane label="变量提取">
<!--            <el-row>-->
<!--              <el-col :span="6">-->
<!--                <el-form-item label="提取变量">-->
<!--                  <el-switch v-model="caseStepForm.isCollectVar"></el-switch>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--            </el-row>-->
<!--            <el-row v-if="caseStepForm.isCollectVar">-->
            <el-row>
              <el-row v-for="(item, index) in caseStepForm.variables"
                      :key="index">
                <el-col :span="22">
                  <el-row>
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
                </el-col>
                <el-col :span="1">
                  <i @click="variableAction(index,1)" v-if="index === 0" class="add" />
                  <i @click="variableAction(index,0)" v-else class="reduce" />
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
<!--    <json-format v-if="stepResDialogVisible" :json-cont="testObj" :col="testCol"></json-format>-->
    <el-row>
      {{ stepLog }}
    </el-row>
    <span slot="footer" class="dialog-footer">
      <el-button @click="stepResDialogVisible = false">取 消</el-button>
    </span>
  </el-dialog>
</div>

</template>

<script>
import {
  addNode,
  editNode,
  delNode,
  getTree,
  getCaseStepList,
  addCaseStep,
  runCase,
  getItfSelectOptions,
  editCaseStep,
  delCaseStep,
  getCaseStepDetail,
  getSingleNodeInfo
} from 'api/case.js'
// import JsonFormat from '@/components/JsonFormat'
export default {
  name: 'ItfCaseManage',
  // components: { JsonFormat },
  data () {
    return {
      projectId: localStorage.getItem('projectId'),
      stepLog: '',
      caseId: '',
      actionFlag: '',
      caseActionFlag: '',
      nodeActionFlag: '',
      dialogVisible: false,
      childDialogVisible: false,
      caseStepDialogVisible: false,
      stepResDialogVisible: false,
      nodeDialogVisible: false,
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
      nodeForm: {
        projectId: '',
        nodeType: '',
        nodeId: '',
        pNodeId: '',
        nodeName: '',
        nodeDesc: ''
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
      count: 1,

      nodeFormRule: {
        nodeType: [{
          required: true,
          message: '请选择类型',
          trigger: 'change'
        }, {
          required: true,
          message: '请选择类型',
          trigger: 'blur'
        }],
        nodeName: [{
          required: true,
          message: '请输入名称',
          trigger: 'change'
        }, {
          required: true,
          message: '请输入名称',
          trigger: 'blur'
        }]
      },
      nodeOptions: [
        {
          label: '模块',
          value: '模块'
        },
        {
          label: '用例',
          value: '用例'
        }
      ]
    }
  },
  filters: {
    statusFilter (val) {
      let statusMap = {
        '成功': 'success',
        '失败': 'danger',
        '未运行': 'info'
      }
      return statusMap[val]
    }
  },
  // created () {
  //   this.getModuleList()
  // },
  mounted () {
    this.getItfSelectOptions()
    this.getNodeList()
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
     * @name: addNode
     * @description: 新增节点
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addNode (data) {
      this.nodeDialogVisible = true
      this.nodeActionFlag = 'add'
      this.nodeForm.pNodeId = data.nodeId
    },

    /**
     * @name: editNode
     * @description: 编辑节点
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    editNode (data) {
      this.nodeActionFlag = 'edit'
      this.nodeForm.nodeId = data.nodeId
      let reqInfo = {
        nodeId: data.nodeId
      }
      getSingleNodeInfo(reqInfo).then(response => {
        this.nodeForm = response
        this.nodeDialogVisible = true
      })
    },

    /**
     * @name: delNodeAction
     * @description: 新增节点
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    delNodeAction (data) {
      let reqInfo = {
        nodeId: data.nodeId
      }
      delNode(reqInfo).then((res) => {
        this.$message({
          message: '恭喜你,删除成功',
          type: 'success'
        })
        this.getNodeList()
      })
    },

    /**
     * @name: addNodeAction
     * @description: 新增节点接口调用
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addNodeAction () {
      this.nodeForm.projectId = this.projectId
      console.log(this.nodeForm)
      this.$refs['nodeForm'].validate((valid) => {
        if (valid) {
          addNode(this.nodeForm).then((res) => {
            this.$message({
              message: '恭喜你,新增成功',
              type: 'success'
            })
            this.nodeDialogVisible = false
            this.getNodeList()
          })
        } else {
          return false
        }
      })
    },

    /**
     * @name: editNodeAction
     * @description: 编辑节点接口调用
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    editNodeAction () {
      this.nodeForm.projectId = this.projectId
      console.log(this.nodeForm)
      this.$refs['nodeForm'].validate((valid) => {
        if (valid) {
          editNode(this.nodeForm).then((res) => {
            this.$message({
              message: '恭喜你,编辑成功',
              type: 'success'
            })
            this.nodeDialogVisible = false
            this.getNodeList()
          })
        } else {
          return false
        }
      })
    },

    runCaseNew () {
      console.log(this.dataList)
      console.log(this.$refs.caseTree.getCheckedNodes())
      let reqInfo = {
        nodeList: this.$refs.caseTree.getCheckedNodes()
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
      console.log(data.nodeName)
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
      // this.loadNode(node)
    },
    getNodeList () {
      // this.$router.projectId
      // getList({ projectId: this.projectId }).then((res) => {
      //   console.log(res)
      //   // this.dataList = res
      //   this.dataList = [{
      //     id: 1,
      //     label: '项目名',
      //     children: res
      //   }]
      // })
      let reqInfo = {
        projectId: localStorage.getItem('projectId')
      }
      getTree(reqInfo).then(response => {
        console.log(response)
        this.dataList = [response]
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
    showStepResDialog (row) {
      console.log(row.stepLog)
      this.stepLog = row.stepLog
      this.stepResDialogVisible = true
    },
    assertAction (index, type) {
      // type: 0 减；1 加
      if (type === 0) {
        this.caseStepForm.assertInfos.splice(index, 1)
      } else if (type === 1) {
        let item = {
          assertCol: '',
          method: '',
          expValue: ''
        }
        this.caseStepForm.assertInfos.push(item)
      }
    },
    variableAction (index, type) {
      // type: 0 减；1 加
      if (type === 0) {
        this.caseStepForm.variables.splice(index, 1)
      } else if (type === 1) {
        let item = {
          assertCol: '',
          method: '',
          expValue: ''
        }
        this.caseStepForm.variables.push(item)
      }
    },
    handleSizeChange () {

    },
    handleCurrentChange () {

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
.add {
  // display: inline-block;
  display: block;
  height: 28px;
  width: 28px;
  margin-left: 5px;
  background: url("../../assets/img/line_add.png") center center;
  background-size: contain;
  float: right;
  margin-right: -5px;
}
.reduce {
  // display: inline-block;
  display: block;
  height: 28px;
  width: 28px;
  background: url("../../assets/img/line_reduce.png") center center;
  background-size: contain;
  float: right;
  margin-right: -5px;
}
</style>
