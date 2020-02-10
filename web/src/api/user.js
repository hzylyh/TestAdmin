/*
 * @Description:总览页面接口
 * @Author: 侯哲宇
 * @Github: https://github.com/hzylyh
 * @Date: 2020-02-10 9:57:53
 * @LastEditors: 侯哲宇
 * @LastEditTime: 2020-02-10 9:57:53
 */
import request from '../utils/request'

export function login (params) {
  return request({
    url: '/api/user/login',
    method: 'POST',
    params: params
  })
}
