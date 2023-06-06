<template>
	<Topnav></Topnav>
	<!-- 成果页面 -->
	<div class="bc">
		<div class="core">
			<!-- 左侧导航部分 -->
			<div class="leftnav">
				<ul>
					<li class="navtitle"><el-icon style="position: relative;top: 3px;right: 5px;"><Trophy /></el-icon>成果展示</li>
					<li v-for="(item,index) in leftnavs" :key="index" @click="whichac(index)">
						{{item}}
						<el-icon style="position: relative;left:50px;top:3px;"><ArrowRight /></el-icon>
					</li>
				</ul>
			</div>
			<!-- 右侧内容部分 -->
			<div class="rightcontent">
				
				<ac v-show="coname=='ac'" @showdetail="showdetail" :actitle="actitle" :achievement="achievement"></ac>
				
				<acitem v-show="coname=='acitem'"></acitem>
				<!-- 注意所有的成果均是展示年份和信息字段，只是根据不同的类别划分在不同的数组中 -->
			</div>
		</div>
	</div>
</template>

<script>
	import ac from "../components/aclist.vue"
	import acitem from "../components/acitem.vue"
	import Topnav from '../components/topnav.vue'
	import API from '../api/index.js'
	export default {
		name:'AchievementView',
		data:function(){
			return{
				coname:'ac',
				actitle:'论文成果',
				leftnavs:["论文","专利","著作","软著","项目","其他","标准","获奖"],
				allac:[],
				achievement:[
				],
				
			}
		},
		created(){
			this.getallac();
		},
		components:{
			ac,
			acitem,
			Topnav,
		},
		methods:{
			getallac(){
				API.AchievementAPI.getachievements().then( res => {
					this.allac=res.data.data.lists
					console.log(this.allac)
					
					})
				
				
				API.AchievementAPI.getstandards().then(res =>{
									for(let i=0;i<res.data.data.lists.length;i++){
										var temp={
											id:'',
											title:'',
											date:'',
											content:'',
											kind:''
										};
										temp.id=res.data.data.lists[i].id
										temp.title=res.data.data.lists[i].title
										temp.date=res.data.data.lists[i].date
										temp.content=res.data.data.lists[i].content
										temp.kind="标准"
										this.allac.push(temp)
									}
										
								})
								API.AchievementAPI.getprojects().then(res =>{
									for(let i=0;i<res.data.data.lists.length;i++){
										var temp={
											id:'',
											title:'',
											date:'',
											content:'',
											kind:''
										};
										temp.id=res.data.data.lists[i].id
										temp.title=res.data.data.lists[i].title
										temp.date=res.data.data.lists[i].date
										temp.content=res.data.data.lists[i].content
										temp.kind="项目"
										this.allac.push(temp)
									}
										
								})
								API.AchievementAPI.getarticles().then(res =>{
									for(let i=0;i<res.data.data.lists.length;i++){
										var temp={
											id:'',
											title:'',
											date:'',
											content:'',
											kind:''
										};
										temp.id=res.data.data.lists[i].id
										temp.title=res.data.data.lists[i].title
										temp.date=res.data.data.lists[i].date
										temp.content=res.data.data.lists[i].content
										temp.kind="论文"
										this.allac.push(temp)
									}
									console.log("论文：")
									console.log(this.allac)
									for(var i=0;i<this.allac.length;i++){
										if(this.allac[i].kind=='论文'){
											this.achievement.push(this.allac[i])
										}
										
								}
								})
				

			},
			whichac(index){
				this.coname='ac';
				var temp=[];
				this.actitle=this.leftnavs[index];
				for(var i=0;i<this.allac.length;i++){
					if(this.allac[i].kind==this.leftnavs[index]){
						temp.push(this.allac[i])
					}
				}
				this.achievement=temp;
			},
			showdetail(item){
				// alert(item);
				this.coname='acitem';
			}
		},
	}
</script>

<style scoped="scoped">
@import url("../css/common.css");
@import url("../css/achievementhome.css");
@import url("../css/nav.css");
</style>
