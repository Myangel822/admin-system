import request from '../util/request.js'
import { getToken } from '@/util/auth'
export default {
	// methods:{
		getpictures(){
					return request({
						method:'get',
						url:`/api/get/getImagesForword`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/image/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/image/updateone`,
				params:{
					web_image:obj
				}
			})
		}
		,
		newone(obj){
			return request({
				method:'put',
				url:`/image/putone`,
				params:{
					web_image:obj
				}
			})
		}
}