<template>
	
	<el-dialog title="" :model-value="memberitemdialogflag" :visible.sync="memberitemdialogflag" @close="closememberitemdialog()" align-center="center" top="5vh" fullscreen="true">
		<div class="bc">
			<div class="core">
			<div class="memberinfo">
				<div class="top">
					<div class="membertitle">
						<h5>{{this.nowmember.name}}</h5>
					</div>
					<div class="membermenu">
						<a href="#mt">个人自述</a>
						<a href="#ms">研究方向</a>
						<a href="#mf">成果展示</a>
						<span class="clsbtn" @click="closememberitemdialog()"><el-icon><Close /></el-icon></span>
					</div>
				</div>
				<div class="content">
					<div class="memberpicture">
						<img :src="require('@/assets/'+this.nowmember.ImageUrl)"  />
						
					</div>
					<div class="memberdetailinfo">
						<h2 style="margin-bottom: 20px;">bistuLab</h2>
						<ul>
							<!-- {{this.nowmember.picture}} -->
							<li>
								职位：{{this.nowmember.identity}}
							</li>
							<li>
								邮箱：{{this.nowmember.mail}}
							</li>
							<li>
								联系方式：{{this.nowmember.phone}}
							</li>
						</ul>
					</div>
				</div>
			</div>
			<div class="memberintroduce" id="mt">
				<div class="top">
					<div class="membertitle">
						<h5>个人自述</h5>
					</div>
				</div>
				
				<div class="ccontent">
					<p v-for="i in this.getirdarr()" :key="i">{{i}}</p>
				</div>
			</div>
			<div class="memberstudy" id="ms">
				<div class="top">
					<div class="membertitle">
						<h5>研究方向</h5>
					</div>
				</div>
				<div class="ccontent">
					<p>{{this.nowmember.research}}</p>
				</div>
			</div>
		
			<div class="memberfruit" id="mf">
				<div class="top">
					<div class="membertitle">
						<h5>成果展示</h5>
					</div>
				</div>
				<div class="fruit">
					<ul>
						<li v-for="i in getacarr()" :key="i">{{i}}</li>
					</ul>
				</div>
			</div>
			</div>
		</div>
	</el-dialog>
</template>

<script>
	export default {
	  name: 'MemberItem',
	  props: {
	    memberitemdialogflag:{
			type:Boolean,
			default:false,
			required:true
		},
		nowmember:{
			type:Object,
			default:false,
			required:true
		},
	  },
	  data() {
	  	return{
			// info:["职位","邮箱","联系方式"]
			ird:[],
			ac:[]
		}
	  },
	  methods:{
		  getirdarr(){
			  console.log("个人简介")
			console.log(this.nowmember.introduction)
			this.ird=this.nowmember.introduction.split("\\n")
			 console.log(this.ird) 
			 return this.ird;
		  },
		  getacarr(){
			  console.log("个人成就")
			  console.log(this.nowmember.achievement)
		  			this.ac=this.nowmember.achievement.split("\\n")
		  			console.log(this.ac) 
		  			 return this.ac;
		  },
		  closememberitemdialog(){
			  this.$emit('closememberitemdialog')
		  }
	  }
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/membersdetail.css");
</style>
