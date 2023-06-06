import request from '../util/request.js'
import { getToken } from '@/util/auth'
export default {
	// methods:{
		getmanager(){
					return request({
						method:'get',
						url:`/api/get/manager/getall`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/manager/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/manager/updateone`,
				params:{
					web_manager:obj
				}
			})
		},
		newone(obj){
			return request({
				method:'put',
				url:`/manager/putone`,
				params:{
					web_manager:obj
				}
			})
		}
}