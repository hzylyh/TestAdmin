<template>
  <div class='apiList'>
    <!-- 接口过滤条件 -->
    <div class="panle">
      <el-form>
        <el-row type="flex"
                justify="start"
                align="middle"
                :gutter="10">
          <div class="fitterRow">
            <label style="font-size:14px;width:80px;color:#606266">接口名称</label>
            <el-input style="width:200px;margin-right:10px"
                      placeholder="请输入"></el-input>
            <el-button type="primary"
                       icon="el-icon-search">搜索</el-button>
            <el-button type="primary"
                       icon="el-icon-search">重置</el-button>
            <el-button type="primary"
                       icon="el-icon-plus"
                       @click="addApi">新增接口</el-button>
            <el-button type="primary"
                       icon="el-icon-plus"
                       @click="addSwagger">同步swagger</el-button>
          </div>

        </el-row>
        <span>
        </span>

      </el-form>
    </div>
    <!-- 接口列表 -->
    <div class="panle list">
      <el-table :data="tableData"
                stripe
                style="width: 100%">
        <el-table-column prop="name"
                         label="接口名称"
                         width="100">
        </el-table-column>
        <el-table-column prop="url"
                         label="接口地址"
                         show-overflow-tooltip
                         width="300">
        </el-table-column>
        <el-table-column prop="type"
                         label="请求方式"
                         width="100">
        </el-table-column>
        <el-table-column prop="mimeType"
                         label="mime类型"
                         width="250">
        </el-table-column>
        <el-table-column prop="dec"
                         label="接口描述">
        </el-table-column>
        <el-table-column label="操作"
                         width="130">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row)"
                       type="text"
                       size="small">查看</el-button>
            <el-button type="text"
                       size="small"
                       @click="showEditApi(scope.row)">编辑</el-button>
            <el-button type="text"
                       size="small"
                       @click="deleteApiAction(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background
                     @current-change="handleCurrentChange"
                     layout="prev, pager, next"
                     :current-page="pageNum"
                     :page-size="pageSize"
                     :total="total"
                     style="margin-top:20px">
      </el-pagination>
    </div>
    <el-dialog title="新增接口"
               :visible.sync="
               dialogVisible"
               width="40%">
      <el-form ref="form"
               :model="form"
               label-width="80px">
        <el-form-item label="接口名称">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="接口地址">
          <el-input v-model="form.url"></el-input>
        </el-form-item>
        <el-form-item label="请求方式">
          <el-select v-model="form.type"
                     placeholder="请选择">
            <el-option v-for="item in options"
                       :key="item.value"
                       :label="item.label"
                       :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="index === 0 ? '请求头' : ''"
                      v-for="(item, index) in form.headers"
                      :key="index">
          <!--          <el-select v-model="form.mimeType"-->
          <!--                     placeholder="请选择">-->
          <!--            <el-option v-for="item in mimeOptions"-->
          <!--                       :key="item.value"-->
          <!--                       :label="item.label"-->
          <!--                       :value="item.value">-->
          <!--            </el-option>-->
          <!--          </el-select>-->
          <el-row :gutter="10"
                  style="width: 100%">
            <el-col :span="11">
              <el-input v-model="item.headerName"></el-input>
            </el-col>
            <el-col :span="11">
              <el-input v-model="item.headerValue"></el-input>
            </el-col>
            <el-col :span="2">
              <i @click="HeaderRowAction(index,1)" v-if="index === 0" class="add" />
              <i @click="HeaderRowAction(index,0)" v-else class="reduce" />
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="接口描述">
          <el-input v-model="form.desc"
                    type="textarea"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer"
            class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary"
                   @click="addApiAction">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog title="同步swagger"
               :visible.sync="
               dialogVisibles"
               width="40%">
      <el-form ref="form"
               :model="form"
               label-width="150px">
        <el-form-item label="swagger接口地址">
          <el-input v-model="swaggerform.swaggerUrl"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer"
            class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button @click="dialogVisible = false">批量同步</el-button>
        <el-button type="primary"
                   @click="addSwaggerAction">全量同步</el-button>
      </span>
    </el-dialog>
    <el-dialog title="修改接口"
               :visible.sync="editDialogVisible"
               width="40%">
      <el-form ref="form"
               :model="form"
               label-width="80px">
        <el-form-item label="接口名称">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="接口地址">
          <el-input v-model="form.url"></el-input>
        </el-form-item>
        <el-form-item label="请求方式">
          <el-select v-model="form.type"
                     placeholder="请选择">
            <el-option v-for="item in options"
                       :key="item.value"
                       :label="item.label"
                       :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="index === 0 ? '请求头' : ''"
                      v-for="(item, index) in form.headers"
                      :key="index">
          <!--          <el-select v-model="form.mimeType"-->
          <!--                     placeholder="请选择">-->
          <!--            <el-option v-for="item in mimeOptions"-->
          <!--                       :key="item.value"-->
          <!--                       :label="item.label"-->
          <!--                       :value="item.value">-->
          <!--            </el-option>-->
          <!--          </el-select>-->
          <el-row :gutter="10">
            <el-col :span="11">
              <el-input v-model="item.headerName"></el-input>
            </el-col>
            <el-col :span="11">
              <el-input v-model="item.headerValue"></el-input>
            </el-col>
            <el-col :span="2">
              <i @click="HeaderRowAction(index,1)" v-if="index === 0" class="add" />
              <i @click="HeaderRowAction(index,0)" v-else class="reduce" />
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="接口描述">
          <el-input v-model="form.desc"
                    type="textarea"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer"
            class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary"
                   @click="editApiAction">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { addApi, editApi, deleteApi, getList, getSingleApi, importSwagger } from '@/api/apiList.js'
export default {
  name: 'itfInterfaceTable',
  data () {
    // 这里存放数据
    return {
      projectId: localStorage.getItem('projectId'),
      dialogVisible: false,
      dialogVisibles: false,
      editDialogVisible: false,
      pageSize: 10,
      pageNum: 1,
      total: 0,
      treeData: [
        {
          label: '全部接口',
          children: [
            {
              label: '接口1'
            }
          ]
        }
      ],
      options: [{
        value: 'GET',
        label: 'GET'
      },
      {
        value: 'POST',
        label: 'POST'
      }],
      mimeOptions: [{
        value: 'application/json',
        label: 'application/json'
      },
      {
        value: 'application/x-www-form-urlencoded',
        label: 'application/x-www-form-urlencoded'
      }],
      tableData: [],
      form: {
        projectId: '',
        name: '',
        url: '',
        desc: '',
        type: 'POST',
        headers: [
          {
            headerId: '',
            interfaceId: '',
            headerName: '',
            headerValue: ''
          }
        ]
      },
      swaggerform: {
        swaggerUrl: ''
      }
    }
  },
  created () {
    this.queryApiList()
  },
  methods: {
    addApi () {
      this.dialogVisible = true
      this.form.headers = [{
        headerId: '',
        interfaceId: '',
        headerName: '',
        headerValue: ''
      }]
    },
    /**
     * @name: showEditApi
     * @description: 显示编辑dialog
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    showEditApi (row) {
      console.log(row.interfaceId)
      let reqInfo = {
        interfaceId: row.interfaceId
      }
      getSingleApi(reqInfo).then(response => {
        console.log(response)
        this.form = response
        if (this.form.headers.length === 0) {
          this.form.headers = [{
            headerId: '',
            interfaceId: '',
            headerName: '',
            headerValue: ''
          }]
        }
        this.editDialogVisible = true
      })
    },

    /**
     * @name: editApiAction
     * @description: 编辑接口
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    editApiAction () {
      editApi(this.form).then(response => {
        console.log(response)
        this.form = response
        this.editDialogVisible = false
        this.$message({
          message: '恭喜你,编辑成功',
          type: 'success'
        })
        this.queryApiList()
      })
    },
    /**
     * @name: addApiAction
     * @description: 新增项目
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addApiAction () {
      this.form.projectId = this.projectId
      debugger
      addApi(this.form).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.dialogVisible = false
        this.queryApiList()
      })
    },

    /**
     * @name: deleteApiAction
     * @description: 删除接口
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    deleteApiAction (row) {
      let reqInfo = {
        id: row.id
      }
      deleteApi(reqInfo).then(response => {
        this.$message({
          message: '恭喜你,删除成功',
          type: 'success'
        })
        this.queryApiList()
      })
    },

    /**
     * 获取接口列表
     * @param
     * @returns
     */
    queryApiList () {
      let reqInfo = {
        projectId: this.projectId,
        pageNum: this.pageNum,
        pageSize: this.pageSize
      }
      getList(reqInfo).then((res) => {
        this.total = res.total
        this.tableData = res.list
      })
    },
    /**
     * @name: addSwagger
     * @description: 同步swagger
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addSwagger () {
      this.dialogVisibles = true
    },
    /**
     * @name: addSwaggerAction
     * @description: 同步swagger接口
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addSwaggerAction () {
      importSwagger(this.swaggerform).then((res) => {
        this.$message({
          message: '恭喜你,同步成功',
          type: 'success'
        })
        this.dialogVisible = false
      })
    },
    handleCurrentChange (val) {
      this.pageNum = val
    },

    /**
     * @name: HeaderRowAction
     * @description: 添加请求头
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    HeaderRowAction (index, type) {
      // type: 0 减；1 加
      if (type === 0) {
        this.form.headers.splice(index, 1)
      } else if (type === 1) {
        let item = {
          headerId: '',
          interfaceId: '',
          headerName: '',
          headerValue: ''
        }
        this.form.headers.push(item)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .apiList {
    height: 100%;
  .panle {
    background-color: #fff;
    border: none;
    /*border-radius: 5px;*/
  // position: relative;
    margin-bottom: 20px;
    box-shadow: 0 1px 1px 0 rgba(0, 0, 0, 0.15);
    box-sizing: border-box;
  }
  .panle {
  .fitterRow {
    padding: 20px 20px 10px 20px;
  }
  }
  .list {
    padding: 20px;
  }
  /deep/.el-form-item__content {
    display: flex;
  }
  }
  .add {
  // display: inline-block;
    display: block;
    height: 28px;
    width: 28px;
    margin-left: 5px;
    background: url("../../../../assets/img/line_add.png") center center;
    background-size: contain;
    float: right;
    margin-right: -5px;
  }
  .reduce {
  // display: inline-block;
    display: block;
    height: 28px;
    width: 28px;
    background: url("../../../../assets/img/line_reduce.png") center center;
    background-size: contain;
    float: right;
    margin-right: -5px;
  }
</style>
