/*
 * @Description:总览页面接口
 * @Author: 侯哲宇
 * @Github: https://github.com/hzylyh
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 侯哲宇
 * @LastEditTime: 2020-03-21 13:29:58
 */
import request from '../utils/request'

export function getCaseRunInfo (params) {
  return request({
    url: '/api/itfPart/dashboard/caseRunInfo',
    // method: 'POST',
    method: 'post',
    data: params
  })
}
