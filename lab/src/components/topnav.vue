<template>
	<nav>
		<div class="nav">
			<!-- pc导航logo区 -->
			<div class="lablogo">
				<a href="https://www.bistu.edu.cn/"><img src='../assets/logo.png'></a>
			</div>
			<span class="labname"><el-icon style="position: relative;top: 5px;"><Cpu /></el-icon>BistuLab</span>
			<div class="navmenu">
				<ul class="pcul">
					<li v-for="menuitem in navmenu" :key="menuitem.path" @click="clickmenu(menuitem.path)">
						{{menuitem.name}}
					</li>
				</ul>
				 <el-dropdown class="ydmu">
				     <el-icon style="position: relative;top: 8px;" ><Menu /></el-icon>
				      <template #dropdown>
				        <el-dropdown-menu>
				          <el-dropdown-item v-for="menuitem in navmenu" :key="menuitem.path" @click="clickmenu(menuitem.path)">{{menuitem.name}}</el-dropdown-item>
				        </el-dropdown-menu>
				      </template>
				</el-dropdown>
			</div>
		</div>
		
	</nav>
</template>

<script>
	export default {
		name:'Topnav',
		data:function(){
			return {
				navmenu:[
					{
						path:'/home',
						name:'主页模块'
					},
					{
						path:'/team',
						name:'成员模块'
					},
					{
						path:'/achievement',
						name:'成果模块'
					},
					{
						path:'/demo',
						name:'demo模块'
					},
					{
						path:'/picture',
						name:'相册模块'
					},
				]
			}
		},
		methods: {
			clickmenu(path) {
				this.$router.push(path);
			},
		
		},
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/nav.css");
</style>
