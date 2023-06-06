<template>
	<ul>
		<li class="topnav">
					<span class="navtitle">{{this.actitle}}</span>
				    <el-breadcrumb separator=">">
				      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
				      <el-breadcrumb-item>成果展示</el-breadcrumb-item>
				      <el-breadcrumb-item>{{this.actitle}}</el-breadcrumb-item>
				     </el-breadcrumb>
		</li>
		<li v-for="item in achievement" :key="item" @click="showdetail(item)">
			<span class="acinfo">{{item.content}}</span>
			<span class="acdate">{{item.date}}</span>
		</li>
	</ul>
</template>

<script>
	import bus from '../eventBus.js'
	export default {
		name:'ac',
		data:function(){
			return{
				nowac:[
				]
			}
		},
		props:{
			actitle:{
				type:String,
				default:'论文成果',
				required:true
			},
			achievement:{
				type:Array,
				required:true
			}
		},
		data:function(){
			return{
				msg:"Wendy's nice try",
			}
		},
		methods:{
			showdetail(item){
				// alert(msg)
				// let info='from '+this.actitle+' '+item+' '+msg
     			this.$emit('showdetail',item)
				bus.emit("senddetail",item)
			}
		}
	}
</script>

<style scoped="scoped">
	@import url("../css/ac.css");
</style>
