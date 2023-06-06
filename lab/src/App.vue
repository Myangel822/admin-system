<template>
<div class="app">
	<!-- <Topnav></Topnav> -->
  <el-backtop :right="100" :bottom="100" />
    <router-view/>
    <Footer></Footer>
</div>
</template>
<script>
	import Footer from '@/components/footer.vue'
	// import Topnav from './components/topnav.vue'
	export default {
		name:'LabApp',
		data:function(){
			return {
			}
		},
		components:{
			Footer,
			// Topnav
		}
	}
</script>
<style scoped="scoped">
@import url("css/common.css");
/* @import url("css/nav.css"); */
</style>
