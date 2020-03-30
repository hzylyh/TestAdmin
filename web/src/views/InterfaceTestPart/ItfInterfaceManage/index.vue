<!--
 * @Description: 接口测试部分接口管理
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->

<template>
  <div style="height: 100%">
    <el-row :gutter="5" style="height: 100%">
      <el-col :span="5" style="height: 100%">
        <el-scrollbar style="height: 100%; background: #fff">
          <el-tree ref="itfTree"
                   :data="treeData"
                   node-key="id"
                   @node-click="handleNodeClick">

          </el-tree>
        </el-scrollbar>
      </el-col>
      <el-col :span="19" style="height: 100%">
<!--        <itf-interface-table></itf-interface-table>-->
        <itf-interface-detail></itf-interface-detail>
      </el-col>
    </el-row>
  </div>
</template>

<script>
// 这里可以导入其他文件（比如：组件，工具js，第三方插件js，json文件，图片文件等等）
// 例如：import 《组件名称》 from '《组件路径》';
import ItfInterfaceTable from './components/itfInterfaceTable'
import ItfInterfaceDetail from './components/itfInterfaceDetail'
export default {
  name: 'ItfInterfaceManage',
  // import引入的组件需要注入到对象中才能使用
  components: {
    ItfInterfaceTable,
    ItfInterfaceDetail
  },
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
  // 监听属性 类似于data概念
  computed: {},
  // 监控data中的数据变化
  watch: {},
  // 方法集合
  methods: {

    handleNodeClick (data, node, element) {
      console.log(data)
    }
  },
  // 生命周期 - 创建完成（可以访问当前this实例）
  created () {
    // this.queryApiList()
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

</style>
