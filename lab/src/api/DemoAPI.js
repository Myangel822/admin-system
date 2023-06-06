import { getToken } from '@/util/auth.js'
import request from '../util/request.js'
export default {
	// methods:{
		getdemos(){
					return request({
						method:'get',
						url:`/api/get/getDemoForword`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/demo/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/demo/updateone`,
				params:{
					web_demo:obj
				}
			})
		}
		,
		newone(obj){
			return request({
				method:'put',
				url:`/demo/putone`,
				params:{
					web_demo:obj
				}
			})
		}
}