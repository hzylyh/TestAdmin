/*
 * @Description:总览页面接口
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 吴文周
 * @LastEditTime: 2019-11-24 15:29:58
 */
import request from '../utils/request'

export function add (params) {
  return request({
    url: '/api/itfPart/case/add',
    method: 'POST',
    params: params
  })
}

export function getList (params) {
  return request({
    url: '/api/itfPart/module/getList',
    method: 'POST',
    params: params
  })
}
export function addModule (params) {
  return request({
    url: '/api/itfPart/case/add',
    method: 'POST',
    params: params
  })
}
