<!--
 * @Description: 接口测试部分用例管理
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
  <el-row class="sl-case-main">
    <el-col :span="5"
            class="sl-suit-tree white-back">
      <el-scrollbar>
        <el-tree
                :props="props"
                :load="loadNode"
                lazy
                show-checkbox
                @check-change="handleCheckChange">
        </el-tree>
      </el-scrollbar>
    </el-col>
    <el-col :span="19"
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
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'ItfCaseManage',
  data () {
    return {
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }],
      props: {
        label: 'name',
        children: 'zones'
      },
      count: 1
    }
  },
  methods: {
    handleCheckChange (data, checked, indeterminate) {
      console.log(data, checked, indeterminate)
    },
    handleNodeClick (data) {
      console.log(data)
    },
    loadNode (node, resolve) {
      if (node.level === 0) {
        return resolve([{ name: 'region1' }, { name: 'region2' }])
      }
      if (node.level > 3) return resolve([])

      var hasChild
      if (node.data.name === 'region1') {
        hasChild = true
      } else if (node.data.name === 'region2') {
        hasChild = false
      } else {
        hasChild = Math.random() > 0.5
      }

      setTimeout(() => {
        var data
        if (hasChild) {
          data = [{
            name: 'zone' + this.count++
          }, {
            name: 'zone' + this.count++
          }]
        } else {
          data = []
        }
        resolve(data)
      }, 500)
    }
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
