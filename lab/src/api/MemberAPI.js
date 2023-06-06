import request from '../util/request.js'
import { getToken } from '@/util/auth'
export default {
	// methods:{
		getmembers(){
					return request({
						method:'get',
						url:`/api/get/getMembersForword`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/member/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/member/updateone`,
				params:{
					web_member:obj
				}
			})
		},
		putone(obj){
			return request({
				method:'put',
				url:`/member/putone`,
				params:{
					web_member:obj
				}
			})
		}
}