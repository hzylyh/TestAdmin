import myrequest from '@/utils/axios'

export function demoGet (pageNum, pageSize, proId) {
  return myrequest({
    url: '/demoGet',
    method: 'get',
    params: { pageNum, pageSize, proId }
  })
}

export function demoPost (params) {
  return myrequest({
    url: '/demoPost',
    method: 'post',
    data: params
  })
}

