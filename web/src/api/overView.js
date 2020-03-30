/*
 * @Description:总览页面接口
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 吴文周
 * @LastEditTime: 2019-11-17 15:40:38
 */
import request from '../utils/request'

export function addProject (params) {
  return request({
    url: '/api/overview/addProject',
    method: 'POST',
    params: params
  })
}

export function getProjectList (params) {
  return request({
    url: '/api/overview/getProjectList',
    method: 'POST',
    params: params
  })
}

export function delProject (params) {
  return request({
    url: '/api/overview/delProject',
    method: 'POST',
    params: params
  })
}

// export function SOLIDER_61(params) {
//   return request({
//     url: '/compcag/SOLIDER_61',
//     method: 'POST',
//     params: params
//   })
// }
