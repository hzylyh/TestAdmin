"use strict";

// import Vue from 'vue';
import axios from "axios";
import router from '../router'
import store from '@/store'
import { getToken } from '@/utils/auth'
import { Message, MessageBox } from 'element-ui'

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

let config = {
  // baseURL: process.env.baseURL || process.env.apiUrl || ""
  // timeout: 60 * 1000, // Timeout
  // withCredentials: true, // Check cross-site Access-Control
  // baseURL: "localhost:36001",
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 60 * 1000,
};

const _axios = axios.create(config);
// console.log(process.env.VUE_APP_SECRET)
// console.log(process.env.VUE_APP_BASE_API)
_axios.interceptors.request.use(
  function (config) {
    // Do something before request is sent
    // if (config.method === 'post') {
    //   const formData = new FormData();
    //   Object.keys(config.data).forEach(key => formData.append(key, config.data[key]));
    //   config.data = formData;
    // }

    // if (config.url !== "/v4/user/login" && getToken() !== undefined ){
    //   config.headers.Authorization = getToken();
    // }
    if (localStorage.getItem('Authorization')) {
      console.log(getToken())
      config.headers.Authorization = localStorage.getItem('Authorization');
    }
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
_axios.interceptors.response.use(
  function (response) {
    // Do something with response data
    // return response;
    const res = response.data
    if (res.code == "21001") {
      router.replace({ name: "user" })
      return response
    } else if (res.code == "20000") {
      return response
    } else {
      Message({
        message: res.msg,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject('error')
    }
  },
  function (error) {
    // Do something with response error
    return Promise.reject(error);
  }
);

// Plugin.install = function(Vue, options) {
//   Vue.axios = _axios;
//   window.axios = _axios;
//   Object.defineProperties(Vue.prototype, {
//     axios: {
//       get() {
//         return _axios;
//       }
//     },
//     $axios: {
//       get() {
//         return _axios;
//       }
//     },
//   });
// };
//
// Vue.use(Plugin)

export default _axios;
