import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/zh-CN'
import './styles.scss'

import './icons' // icon

Vue.config.productionTip = false

Vue.use(ElementUI, {
  locale,
  size: 'medium'
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
