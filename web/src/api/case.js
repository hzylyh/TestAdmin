/*
 * @Description:总览页面接口
 * @Author: 侯哲宇
 * @Github: https://github.com/hzylyh
 * @Date: 2019-06-05 18:57:53
 * @LastEditors: 侯哲宇
 * @LastEditTime: 2020-03-21 13:29:58
 */
import request from '../utils/request'

export function getSingleNodeInfo (params) {
  return request({
    url: '/api/itfPart/case/tree/get',
    // method: 'POST',
    method: 'post',
    data: params
  })
}

export function addNode (params) {
  return request({
    url: '/api/itfPart/case/tree/add',
    method: 'post',
    data: params
  })
}

export function editNode (params) {
  return request({
    url: '/api/itfPart/case/tree/edit',
    method: 'post',
    data: params
  })
}

export function delNode (params) {
  return request({
    url: '/api/itfPart/case/tree/delete',
    method: 'post',
    data: params
  })
}

export function addModule (params) {
  return request({
    url: '/api/itfPart/module/add',
    method: 'post',
    data: params
  })
}

export function delModule (params) {
  return request({
    url: '/api/itfPart/module/delete',
    method: 'post',
    data: params
  })
}

export function getCaseList (params) {
  return request({
    url: '/api/itfPart/case/getList',
    method: 'post',
    data: params
  })
}

export function getTree (params) {
  return request({
    url: '/api/itfPart/case/getTree',
    method: 'post',
    data: params
  })
}

export function getList (params) {
  return request({
    url: '/api/itfPart/module/getList',
    method: 'post',
    data: params
  })
}
export function addCase (params) {
  return request({
    url: '/api/itfPart/case/add',
    method: 'post',
    data: params
  })
}

export function addCaseTree (params) {
  return request({
    url: '/api/itfPart/case/tree/init',
    method: 'post',
    data: params
  })
}

export function editCase (params) {
  return request({
    url: '/api/itfPart/case/edit',
    method: 'post',
    data: params
  })
}

export function getCase (params) {
  return request({
    url: '/api/itfPart/case/get',
    method: 'post',
    data: params
  })
}

export function delCase (params) {
  return request({
    url: '/api/itfPart/case/delete',
    method: 'post',
    data: params
  })
}

export function getCaseStepList (params) {
  return request({
    url: '/api/itfPart/case/step/getList',
    method: 'post',
    data: params
  })
}

export function addCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/add',
    method: 'post',
    data: params
  })
}

export function getCaseStepDetail (params) {
  return request({
    url: '/api/itfPart/case/step/get',
    method: 'post',
    data: params
  })
}

export function editCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/edit',
    method: 'post',
    data: params
  })
}

export function delCaseStep (params) {
  return request({
    url: '/api/itfPart/case/step/delete',
    method: 'post',
    data: params
  })
}

export function runCase (params) {
  return request({
    url: '/api/itfPart/case/run',
    method: 'post',
    data: params
  })
}

export function getItfSelectOptions (params) {
  return request({
    url: '/api/itfPart/interface/getSelectOptions',
    method: 'post',
    data: params
  })
}
