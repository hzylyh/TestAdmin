// vue.config.js
module.exports = {
  lintOnSave: false,
  publicPath: './', // vueConf.baseUrl, // 根域上下文目录

  assetsDir: 'static',
  // 选项...
  devServer: {
    host: 'localhost',
    port: 8081,
    proxy: {
      // 配置跨域
      '/api': {
        // 要访问的跨域的api的域名
        target: 'http://123.207.189.27:9000',
        ws: false
      }
    }
  }
}
