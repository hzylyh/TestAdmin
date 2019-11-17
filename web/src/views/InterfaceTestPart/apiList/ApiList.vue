<!--  -->
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
          </div>

        </el-row>
        <span>
        </span>

      </el-form>
    </div>
    <!-- 接口列表 -->
    <div class="panle list">
      <el-table :data="tableData"
                border
                style="width: 100%">
        <el-table-column prop="name"
                         label="接口名称"
                         width="100">
        </el-table-column>
        <el-table-column prop="url"
                         label="接口地址"
                         width="300">
        </el-table-column>
        <el-table-column prop="type"
                         label="请求方式"
                         width="100">
        </el-table-column>
        <el-table-column prop="dec"
                         label="接口描述">
        </el-table-column>
        <el-table-column prop="address"
                         label="操作"
                         width="100">
          <template slot-scope="scope">
            <el-button @click="handleClick(scope.row)"
                       type="text"
                       size="small">查看</el-button>
            <el-button type="text"
                       size="small">编辑</el-button>

          </template>
        </el-table-column>
      </el-table>
      <el-pagination background
                     layout="prev, pager, next"
                     :total="1000"
                     style="margin-top:20px">
      </el-pagination>
    </div>
    <el-dialog title="新建项目"
               :visible.sync="
               dialogVisible"
               width="60%">
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
        <el-form-item label="接口描述">
          <el-input v-model="form.desc"
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
// 这里可以导入其他文件（比如：组件，工具js，第三方插件js，json文件，图片文件等等）
// 例如：import 《组件名称》 from '《组件路径》';
import { addApi, getList } from '@/api/apiList.js'
export default {
  // import引入的组件需要注入到对象中才能使用
  components: {},
  data () {
    // 这里存放数据
    return {
      dialogVisible: false,
      options: [{
        value: 'GET',
        label: 'GET'
      },
      {
        value: 'POST',
        label: 'POST'
      }],
      tableData: [],
      form: {
        name: '',
        ulr: '',
        desc: '',
        type: 'POST'
      },
      queryForm: {
        pageNmber: 1,
        pageSize: 10
      }
    }
  },
  // 监听属性 类似于data概念
  computed: {},
  // 监控data中的数据变化
  watch: {},
  // 方法集合
  methods: {
    addApi () {
      this.dialogVisible = true
    },
    /**
     * @name: addPojectAction
     * @description: 新增项目
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    addPojectAction () {
      addApi(this.form).then((res) => {
        this.$message({
          message: '恭喜你,新增成功',
          type: 'success'
        })
        this.dialogVisible = false
      })
    },
    queryApiList () {
      getList(this.queryForm).then((res) => {
        this.tableData = res.list
      })
    }
  },
  // 生命周期 - 创建完成（可以访问当前this实例）
  created () {
    this.queryApiList()
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
.apiList {
  .panle {
    background-color: #fff;
    border: none;
    border-radius: 5px;
    // position: relative;
    margin-bottom: 20px;
    box-shadow: 0 5px 5px 0 rgba(0, 0, 0, 0.25);
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
</style>
