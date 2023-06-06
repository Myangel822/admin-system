<template>
	<Topnav></Topnav>
	<!-- 实验室相册页面 -->
	<div class="bc">
		<div class="jaycore">
			<div class="picturetop">
				<span class="labnamee"><el-icon style="position: relative;top: 5px;"><Cpu /></el-icon>BistuLab</span>
				<span class="picturetitle"><el-icon style="position: relative;top: 5px;"><Camera /></el-icon>Picture列表</span>
			</div>
<!-- 			<ul class="picturecon">
				<li class="pictureone" v-for="i in 15">
					
				</li>
			</ul> -->
			<!-- 相册列表 -->
			<pageitem></pageitem>
		</div>
	</div>
</template>

<script>
	import pageitem from '../components/pageitem.vue'
	import Topnav from '../components/topnav.vue'
	export default{
		name:'PictureView',
		data:function(){
			return{
				
			}
		},
		props:{
			
		},
		components:{
			pageitem,
			Topnav,
		}
		
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/picturehome.css");
	@import url("../css/nav.css");
</style>
