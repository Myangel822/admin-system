import request from '../util/request.js'
import { getToken } from '@/util/auth'
export function login(data) { 
    var username = data.username
    var password = data.password
  return request({
    url: '/manager',
    method: 'get',
    params: { username , password }
  })
}

export function getInfo(token) {
  return request({
    url: '/api/get/articles',
    method: 'get',
    params: { token }
  })
}

export function logout(token) {
  return request({
    url: '/api/get/articles',
    method: 'get',
    params: { token }
  })
}

