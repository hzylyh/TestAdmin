// vue.config.js
module.exports = {
  publicPath: './', //vueConf.baseUrl, // 根域上下文目录

  assetsDir: 'static',
  // 选项...
  devServer: {
    host: 'localhost',
    port: 8081,
    proxy: { // 配置跨域
      '/*': {
        //要访问的跨域的api的域名
        target: 'http://localhost:8089/salotto',
        ws: false,
        changeOrigin: true,
      }
    },
  }
}
