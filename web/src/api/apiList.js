/*
 * @Description:总览页面接口
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 吴文周
 * @LastEditTime: 2019-11-17 20:04:47
 */
import request from '../utils/request'

export function addApi (params) {
  return request({
    url: '/api/itfPart/interface/add',
    method: 'post',
    data: params
  })
}

export function editApi (params) {
  return request({
    url: '/api/itfPart/interface/edit',
    method: 'post',
    data: params
  })
}

export function deleteApi (params) {
  return request({
    url: '/api/itfPart/interface/delete',
    method: 'post',
    data: params
  })
}

export function getList (params) {
  return request({
    url: '/api/itfPart/interface/getList',
    method: 'post',
    data: params
  })
}
export function importSwagger (params) {
  return request({
    url: '/api/itfPart/interface/importSwagger',
    method: 'post',
    data: params
  })
}

export function getSingleApi (params) {
  return request({
    url: '/api/itfPart/interface/get',
    method: 'post',
    data: params
  })
}

// export function SOLIDER_61(params) {
//   return request({
//     url: '/compcag/SOLIDER_61',
//     method: 'post',
//     data: params
//   })
// }
