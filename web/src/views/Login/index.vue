<template>
  <div class="login-page">
    <el-card class="login-form"
             size="small">
      <div slot="header">
        <span>欢迎登录</span>
      </div>
      <el-form :model="loginForm"
               label-width="60px">
        <el-form-item label="用户名:">
          <el-input v-model="loginForm.username"
                    size="small"></el-input>
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="loginForm.password"
                    type="password"
                    size="small"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="login(loginForm)"
                     type="primary"
                     size="small">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'user',
  data () {
    return {
      loginForm: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    // 对象
    // this.$router.push({path: '/user?url=' + this.$route.path});

    // 命名的路由
    // router.push({ name: 'user', params: { userId: 123 }})

    // 带查询参数，变成/backend/order?selected=2
    // //this.$router.push({path: '/backend/order', query: {selected: "2"}});
    login (loginInfo) {
      // this.$store.dispatch('Login', loginInfo).then(() => {
      //   // this.$router.push({name: "itfDashboard"})
      //   console.log('登录成功')
      // })
      this.$router.push({ name: 'OverviewIndex' })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.login-page {
  width: 100%;
  height: 100%;
  background: url("../../assets/login-bg.jpg") no-repeat;
  background-size: cover;
  position: relative;
  background-position: 50%;
  background-position-x: 50%;
  background-position-y: center;
  .login-form {
    position: absolute;
    right: 160px;
    top: 30%;
    width: 300px;
    .el-form-item {
      margin-bottom: 10px;
      button {
        float: right;
      }
    }
  }
}
</style>
