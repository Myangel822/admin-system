<template>
	<Topnav></Topnav>
	<!-- 实验室成员页面 -->
	<div class="bc">
	<div class="core">
		<!-- 顶部logo与搜索框 -->
		<div class="topfunction">
			<div class="weblogo"><el-icon style="position: relative; top: 3px;"><Cpu /></el-icon>BistuLab人才中心</div>
			<div class="search">
				<el-input clearable="true" v-model="searchvalue" placeholder="请输入人名" prefix-icon="search"></el-input>
				<!-- <el-button>搜索</el-button> -->
			</div>
		</div>
		<!-- 人员列表 -->
		
		<div v-for="(item,index) in team" :key="item.type"  class="showmembers">
			<div class="typetitle">
				{{item.identity}}
			</div>
			<div class="typelist">
				<ul>
					<li v-for="i in filtermembers(index)" :key="i.name" @click="showmemberitemdialog(item.identity,i.id)">
						<el-icon style="position: relative; top:2px;"><UserFilled /></el-icon>
						{{i.name}}
					</li>
<!-- 					<li v-for="i in testmembers" :key="i.id_isp_test">
						<el-icon style="position: relative; top:2px;"><UserFilled /></el-icon>
						{{i.name_isp_test}}
					</li> -->
				</ul>
			</div>
		</div> 
	</div>
	<!-- 人物详情组件 -->
	<MemberItem  :memberitemdialogflag="memberitemdialogflag" @closememberitemdialog="closememberitemdialog" :nowmember="nowmember"></MemberItem>
	</div>
</template>

<script>
	import MemberItem from '../components/memberitem.vue'
	import Topnav from '../components/topnav.vue'
	import API from '../api/index.js'
	export default {
		name:'TeamView',
		created() {
			this.getallmembers()
		},
		methods:{
			//获取所有新闻，根据类型传入不同的member中
			getallmembers(){
				API.MemberAPI.getmembers().then( res => {
					console.log(res.data.data)
					for(var i = 0; i < res.data.data.lists.length;i++)
					{
						console.log(res.data.data.lists[i])
						if(res.data.data.lists[i].identity==='系统管理员'){
							this.team[0].members.push(res.data.data.lists[i])
						}
						if(res.data.data.lists[i].identity==='教师'){
							this.team[1].members.push(res.data.data.lists[i])
						}
						if(res.data.data.lists[i].identity==='在校学生'){
							this.team[2].members.push(res.data.data.lists[i])
						}
						if(res.data.data.lists[i].identity==='毕业学生'){
							this.team[3].members.push(res.data.data.lists[i])
						}
					}
					console.log(this.team)
				})
			},
			filtermembers(index) {
				if (this.searchvalue != '') {
					var reg = new RegExp(this.searchvalue);
					var arr = [];
					for (var i = 0; i < this.team[index].members.length; i++) {
						
						if (reg.test(this.team[index].members[i].name)) {
							arr.push(this.team[index].members[i]);
						}
					}
					// alert(arr)
					return arr;
				}
				else{
					// alert(this.team[0].members)
					return this.team[index].members;
				}
			},
			showmemberitemdialog(identity,id){
				//展示具体的人物详情，根据点击对象的id获取具体人物后传入组件中
							for(let i=0;i<this.team.length;i++)
							{
								if(this.team[i].identity==identity){
									for(let j=0;j<this.team[i].members.length;j++){
										if(id==this.team[i].members[j].id){
											this.nowmember=this.team[i].members[j];
										}
									}
								}
							}
							  this.memberitemdialogflag=true;
							  // alert(this.newsitemdialogflag);
			},
			closememberitemdialog(){
							  this.memberitemdialogflag=false;
			}
		},
		components: {
			MemberItem,
			Topnav,
		},

		data:function(){
			return{
				memberitemdialogflag:false,
				searchvalue:"",
				testmembers:[],
				nowmember:{
					id:"",
					type:"",
					picture:"",
					identity:"",
					name:"",
					phone:"",
					mail:"",
					research:"",
					achievement:"",
					introduction:"",
					state:"",
				},
				//注意！！！这里的成员是一个二级的累数组 嵌套，现根据成员类型分为四种大体，然后再将对应的人物添加到各个大体中
				team:[
					{
						identity:'系统管理员',
						members:[
							
						]
					},
					{
						identity:'教师',
						members:[
							
						]
					},
					{
						identity:'在校学生',
						members:[
						
						]
					},
					{
						identity:'毕业学生',
						members:[
							
						]
					}
				]
			}
		}
		}
		
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/teamhome.css");
	@import url("../css/nav.css");
</style>
