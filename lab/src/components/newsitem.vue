<template>
	<el-dialog title="" :model-value="newsitemdialogflag" :visible.sync="newsitemdialogflag" @close="closenewsitemdialog()" align-center="center" top="5vh" fullscreen="true">
		
		<div class="bc">
			<div class="core">
			<div class="detailtitle"> {{this.nownews.title}} <span class="clsbtn" @click="closenewsitemdialog()"><el-icon><Close /></el-icon></span></div>
			<div class="detaildate">{{this.nownews.created_on}}</div>
			<div class="detailcontent">{{this.nownews.content}}</div>
			<!-- <el-button class="closebt" @click="closenewsitemdialog()">关闭</el-button> -->
			</div>
		</div>
	</el-dialog>
</template>

<script>
	export default {
	  name: 'NewsItem',
	  nownews:{
	  	id:"",
	  	pitcture:'',
	  	title:'',
	  	created_on:"",
	  	content:"",
	  },
	  props: {
	    newsitemdialogflag:{
			type:Boolean,
			default:false,
			required:true
		},
		nownews:{
			type:Object,
			default:false,
			required:true
		},
	  },
	  data() {
	  	return{
			
		}
	  },
	  methods:{
		  closenewsitemdialog(){
			  this.$emit('closenewsitemdialog')
		  }
	  }
	}
</script>

<style scoped="scoped">
	@import url("../css/newsdetail.css");
	@import url("../css/common.css");
</style>
