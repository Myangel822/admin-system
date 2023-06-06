<template>
	<Topnav></Topnav>
	<!-- demo页面 -->
	<div class="bc">
		<div class="jaycore">
			<div class="demotopp">
				<span class="labnamee"><el-icon style="position: relative;top: 5px;"><Cpu /></el-icon>BistuLab</span>
				<span class="demotitle"><el-icon style="position: relative;top: 5px;"><View /></el-icon>Demo展示</span>
			</div>
			<ul class="demolist">
				<li v-for="j in jay" :key="j.title">
					<div class="demotop">
						<div class="demoname">{{j.title}}</div>
						<div class="demointroduce"><p>{{j.date}}</p></div>
						
					</div>
					<div class="demovideo">
						<video class="vd" :src="require('@/assets/'+j.Video)" controls="controls">//注意，demo必须放在assets下，并且数据库只存储了图片名称 不包含文件类型后缀和路径
						</video>
					</div>
				</li>
			</ul>
		</div>
	</div>
</template>

<script>
	import Topnav from '../components/topnav.vue'
	import API from '../api/index.js'
	export default{
		name:'DemoView',
		data:function(){
			return{
				jay:[
				]
			}
		},
		created(){
			this.getalldemo()
		},
		methods: {
			getalldemo() {
				API.DemoAPI.getdemos().then( res => {
					this.jay = res.data.data.lists
					console.log("------------------------")
					console.log(this.jay)
				})
			}
		},
		components:{
			Topnav,
		}
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/demohome.css");
	@import url("../css/nav.css");
</style>
