<template>
  <div class="login-page">
    <el-card class="login-form"
             size="small">
      <div slot="header">
        <span>欢迎登录</span>
      </div>
      <el-form ref="loginForm"
               :model="loginForm"
               label-width="60px">
        <el-form-item label="用户名:">
          <el-input ref="username"
                    v-model="loginForm.username"
                    size="small"></el-input>
        </el-form-item>
        <el-form-item label="密码:">
          <el-input ref="password"
                    v-model="loginForm.password"
                    type="password"
                    size="small"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleLogin"
                     type="primary"
                     size="small">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
// import { validUsername } from '@/utils/validate'
export default {
  name: 'user',
  data () {
    // const validateUsername = (rule, value, callback) => {
    //   if (!validUsername(value)) {
    //     callback(new Error('Please enter the correct user name'))
    //   } else {
    //     callback()
    //   }
    // }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: '',
        password: ''
      }
    }
  },
  mounted () {
    if (this.loginForm.username === '') {
      this.$refs.username.focus()
    } else if (this.loginForm.password === '') {
      this.$refs.password.focus()
    }
  },
  methods: {
    // 对象
    // this.$router.push({path: '/user?url=' + this.$route.path});

    // 命名的路由
    // router.push({ name: 'user', params: { userId: 123 }})

    // 带查询参数，变成/backend/order?selected=2
    // //this.$router.push({path: '/backend/order', query: {selected: "2"}});
    handleLogin () {
      // this.$store.dispatch('Login', loginInfo).then(() => {
      //   // this.$router.push({name: "itfDashboard"})
      //   console.log('登录成功')
      // })
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          // this.loading = true
          this.$store.dispatch('user/login', this.loginForm)
            .then(() => {
              // this.$router.push({ path: this.redirect || '/', query: this.otherQuery })
              this.$router.push({ name: 'OverviewIndex' })
              // this.loading = false
            })
            .catch(() => {
              // this.loading = false
            })
        } else {
          return false
        }
      })
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
