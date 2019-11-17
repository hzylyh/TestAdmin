/*
 * @Description:总览页面接口
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 吴文周
 * @LastEditTime: 2019-11-17 17:47:47
 */
import request from '../utils/request'

export function addApi (params) {
  return request({
    url: '/api/itfPart/interface/add',
    method: 'POST',
    params: params
  })
}

export function getList (params) {
  return request({
    url: '/api/itfPart/interface/getList',
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
