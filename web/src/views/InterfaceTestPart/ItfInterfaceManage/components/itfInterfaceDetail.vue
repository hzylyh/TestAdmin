<template>
  <div class="sl-itf-table">
    <div>
      <h2 class="sl-common-title">基本信息</h2>
      <el-row class="sl-each-part">
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            状态: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            更新时间: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口路径: test
          </el-col>
        </el-row>
      </el-row>

    </div>
    <div>
      <h2 class="sl-common-title">请求参数</h2>
      <el-row class="sl-each-part">
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            状态: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            更新时间: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口路径: test
          </el-col>
        </el-row>
      </el-row>

    </div>
    <div>
      <h2 class="sl-common-title">返回数据</h2>
      <el-row class="sl-each-part">
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            状态: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            更新时间: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口名称: test
          </el-col>
          <el-col class="sl-each-col" :span="12">
            创建人: test
          </el-col>
        </el-row>
        <el-row type="flex"
                class="sl-each-row"
                align="center">
          <el-col class="sl-each-col" :span="12">
            接口路径: test
          </el-col>
        </el-row>
      </el-row>

    </div>
  </div>
</template>

<script>
export default {
  name: 'itfInterfaceDetail'
}
</script>

<style lang="scss" scoped>
.sl-itf-table {
  height: 100%;
  background: #fff;
  padding-top: 10px;
  .sl-common-title {
    text-align: left;
    clear: both;
    font-weight: 400;
    /*margin-top: .48rem;*/
    /*margin-bottom: .16rem;*/
    /*padding-top: .48rem;*/
    border-left: 3px solid #2395f1;
    padding-left: 8px;
  }
  .sl-each-part {
    padding: 15px 15px 0 15px;
    .sl-each-row {
      text-align: left;
      margin-bottom: 10px;
    }
  }
}
</style>
