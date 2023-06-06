import request from '../util/request.js'
import { getToken } from '@/util/auth'
export default {
	// methods:{
		getarticles(){
					return request({
						method:'get',
						url:`/api/get/articleforword`,
						params:{
							token:getToken()
						}
					})
		},
		getprojects(){
					return request({
						method:'get',
						url:`/api/get/getProjectsForword`,
						params:{
							token:getToken()
						}
					})
		},
		getstandards(){
					return request({
						method:'get',
						url:`/api/get/getStandardsForword`,
						params:{
							token:getToken()
						}
					})
		},
		getachievements(){
					return request({
						method:'get',
						url:`/api/get/achievementsForword`,
						params:{
							token:getToken()
						}
					})
		},
		deleteone(whichid){
			return request({
				method:'delete',
				url:`/achievement/deleteone`,
				params:{
					id:whichid
				}
			})
		},
		updateone(obj){
			return request({
				method:'post',
				url:`/achievement/updateone`,
				params:{
					web_achievement:obj
				}
			})
		},
		newone(obj){
			return request({
				method:'put',
				url:`/achievement/putone`,
				params:{
					web_achievement:obj
				}
			})
		}
}