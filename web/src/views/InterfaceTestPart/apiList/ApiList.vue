<!--  -->
<template>
  <div class=''>
    <!-- 接口过滤条件 -->
    <div>
      <el-form>
        <el-row type="flex"
                justify="start"
                align="middle"
                :gutter="10">
          <label style="font-size:14px;width:80px;color:#606266">接口名称</label>
          <el-input style="width:200px;margin-right:10px"
                    placeholder="请输入"></el-input>
          <el-button type="primary"
                     icon="el-icon-search">搜索</el-button>
          <el-button type="primary"
                     icon="el-icon-search">重置</el-button>
        </el-row>
        <span>
        </span>

      </el-form>
    </div>
    <!-- 接口列表 -->
    <div>
      <el-table :data="tableData"
                border
                style="width: 100%">
        <el-table-column prop="date"
                         label="日期"
                         width="180">
        </el-table-column>
        <el-table-column prop="name"
                         label="姓名"
                         width="180">
        </el-table-column>
        <el-table-column prop="address"
                         label="操作">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row)"
                       type="text"
                       size="small">查看</el-button>
            <el-button type="text"
                       size="small">编辑</el-button>
            <el-button type="text"
                       size="small">新增接口</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background
                     layout="prev, pager, next"
                     :total="1000"
                     style="float:right;margin-top:10px">
      </el-pagination>
    </div>
  </div>
</template>

<script>
// 这里可以导入其他文件（比如：组件，工具js，第三方插件js，json文件，图片文件等等）
// 例如：import 《组件名称》 from '《组件路径》';

export default {
  // import引入的组件需要注入到对象中才能使用
  components: {},
  data () {
    // 这里存放数据
    return {
      tableData: [{
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-08',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-06',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-07',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }]
    }
  },
  // 监听属性 类似于data概念
  computed: {},
  // 监控data中的数据变化
  watch: {},
  // 方法集合
  methods: {

  },
  // 生命周期 - 创建完成（可以访问当前this实例）
  created () {

  },
  // 生命周期 - 挂载完成（可以访问DOM元素）
  mounted () {

  },
  beforeCreate () { }, // 生命周期 - 创建之前
  beforeMount () { }, // 生命周期 - 挂载之前
  beforeUpdate () { }, // 生命周期 - 更新之前
  updated () { }, // 生命周期 - 更新之后
  beforeDestroy () { }, // 生命周期 - 销毁之前
  destroyed () { }, // 生命周期 - 销毁完成
  activated () { } // 如果页面有keep-alive缓存功能，这个函数会触发
}
</script>
<style lang='scss' scoped>
//@import url(); 引入公共css类
</style>
