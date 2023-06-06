<template>
	<div class="acdetail">
		<div class="topnav">
					<span class="navtitle">成果详情</span>
				    <el-breadcrumb separator=">">
				      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
				      <el-breadcrumb-item>成果展示</el-breadcrumb-item>
				      <el-breadcrumb-item>成果详情</el-breadcrumb-item>
				     </el-breadcrumb>
		</div>
		<div class="accontent">
			<div class="detailtitle">
				{{achievement.title}}
			</div>
			<div class="detaildate">
				<font>{{achievement.date}}</font>
			</div>
			<div class="detaicontent">
				{{achievement.content}}
			</div>
		</div>
	</div>
</template>

<script>
	import bus from "../eventBus.js"
	export default {
		name:'acitem',
		data:function(){
			return{
				msg:'',
				achievement:
					{
						id:"",
						user_id:"",
						title:"",
						date:"",
						content:"",
						kind:"",
					}
			}
		},
		beforeCreate() {
			bus.on("senddetail",(res) => {
				this.achievement=res
			})
		}
	}
</script>

<style scoped="scoped">
	@import url("../css/acdetail.css");
</style>
