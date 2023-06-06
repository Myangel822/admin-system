import request from '../util/request.js'
import { getToken } from '@/util/auth'
export default {
	// methods:{
		getnews(){
					return request({
						method:'get',
						url:`/api/get/getnewsForword`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/news/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/news/updateone`,
				params:{
					web_news:obj
				}
			})
		},
		newone(obj){
			return request({
				method:'put',
				url:`/news/putone`,
				params:{
					web_news:obj
				}
			})
		}
}