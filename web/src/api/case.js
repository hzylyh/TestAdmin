/*
 * @Description:总览页面接口
 * @Author: 侯哲宇
 * @Github: https://github.com/hzylyh
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 侯哲宇
 * @LastEditTime: 2020-03-21 13:29:58
 */
import request from '../utils/request'

export function addModule (params) {
  return request({
    url: '/api/itfPart/module/add',
    method: 'POST',
    params: params
  })
}

export function getCaseList (params) {
  return request({
    url: '/api/itfPart/case/getList',
    method: 'POST',
    params: params
  })
}

export function getCaseTree (params) {
  return request({
    url: '/api/itfPart/case/getTree',
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
export function addCase (params) {
  return request({
    url: '/api/itfPart/case/add',
    method: 'POST',
    params: params
  })
}

export function delCase (params) {
  return request({
    url: '/api/itfPart/case/delete',
    method: 'POST',
    params: params
  })
}

export function getCaseStepList (params) {
  return request({
    url: '/api/itfPart/case/step/getList',
    method: 'POST',
    params: params
  })
}

export function addCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/add',
    method: 'POST',
    params: params
  })
}

export function getCaseStepDetail (params) {
  return request({
    url: '/api/itfPart/case/step/get',
    method: 'POST',
    params: params
  })
}

export function editCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/edit',
    method: 'POST',
    params: params
  })
}

export function delCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/delete',
    method: 'POST',
    params: params
  })
}

export function runCase (params) {
  return request({
    url: '/api/itfPart/case/run',
    method: 'POST',
    params: params
  })
}

export function getItfSelectOptions (params) {
  return request({
    url: '/api/itfPart/interface/getSelectOptions',
    method: 'POST',
    params: params
  })
}
