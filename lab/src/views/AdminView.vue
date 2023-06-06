<template>
	<Topnav></Topnav>
	<div class="bc">
		<div class="core">
			<div class="leftnav">
				<ul>
					<li class="navtitle"><el-icon style="position: relative; top: 3px;right: 5px;"><Setting /></el-icon>系统管理中心</li>
					<li v-for="(item,index) in adminmenu" :key="index" @click="showwhich(index)">
						{{item}}
						<el-icon style="position: relative;left:50px;top:3px;"><ArrowRight /></el-icon>
					</li>
				</ul>
			</div>
			<div class="rightcontent">
				<Usertabla v-show="tablename=='user'" @shownewdlg="shownewdlg"></Usertabla>
				<Membertable v-show="tablename=='member'" @shownewdlg="shownewdlg"></Membertable>
				<Actable v-show="tablename=='ac'" @shownewdlg="shownewdlg"></Actable>
				<Demotable v-show="tablename=='demo'" @shownewdlg="shownewdlg"></Demotable>
				<Elsetable v-show="tablename=='else'" @shownewdlg="shownewdlg"></Elsetable>
				<Newstable v-show="tablename=='news'" @shownewdlg="shownewdlg"></Newstable>
			</div>
			<Newuserdlg :shownewuserdlg="shownewuserdlg" @closenewdlg="closenewdlg"></Newuserdlg>
			<Newmemberdlg :shownewmemberdlg="shownewmemberdlg" @closenewdlg="closenewdlg"></Newmemberdlg>
			<Newacdlg :shownewacdlg="shownewacdlg" @closenewdlg="closenewdlg"></Newacdlg>
			<Newdemodlg :shownewdemodlg="shownewdemodlg" @closenewdlg="closenewdlg"></Newdemodlg>
			<Newpcdlg  :shownewpcdlg="shownewpcdlg" @closenewdlg="closenewdlg"></Newpcdlg>
			<Newnewsdlg  :shownewnewsdlg="shownewnewsdlg" @closenewdlg="closenewdlg"></Newnewsdlg>
			<!-- <el-button @click="shownewdlg()">新增</el-button> -->
		</div>
	</div>
</template>

<script>
	import Topnav from '../components/topnav.vue'
	import Usertabla from '../components/userstable.vue'
	import Membertable from '../components/membertable.vue'
	import Actable from '../components/actable.vue'
	import Demotable from '../components/demotable.vue'
	import Elsetable from '../components/elsetable.vue'
	import Newstable from '../components/newstable.vue'
	import Newuserdlg from '../components/newuser.vue'
	import Newmemberdlg from '../components/newmember.vue'
	import Newacdlg from '../components/newac.vue'
	import Newdemodlg from '../components/newdemo.vue'
	import Newpcdlg from '../components/newpc.vue'
	import Newnewsdlg from '../components/newnews.vue'
	export default {
		name:'AdminView',
		data:function(){
			return{
				adminmenu:['用户管理','成员管理','成果管理','视频管理','相册管理',"新闻管理"],
				tablename:'user',
				shownewuserdlg:false,
				shownewmemberdlg:false,
				shownewacdlg:false,
				shownewdemodlg:false,
				shownewpcdlg:false,
				shownewnewsdlg:false
			}
		},
		methods:{
			shownewdlg(){
				if(this.tablename=='user'){
					this.shownewuserdlg=true
				}else if(this.tablename=='member'){
					this.shownewmemberdlg=true
				}else if(this.tablename=='ac'){
					this.shownewacdlg=true
				}else if(this.tablename=='demo'){
					this.shownewdemodlg=true
				}else if(this.tablename=='else'){
					this.shownewpcdlg=true
				}else if(this.tablename=='news'){
					this.shownewnewsdlg=true
				}
			},
			closenewdlg(){
				if(this.tablename=='user'){
					this.shownewuserdlg=false
				}else if(this.tablename=='member'){
					this.shownewmemberdlg=false
				}else if(this.tablename=='ac'){
					this.shownewacdlg=false
				}else if(this.tablename=='demo'){
					this.shownewdemodlg=false
				}else if(this.tablename=='else'){
					this.shownewpcdlg=false
				}else if(this.tablename=='news'){
					this.shownewnewsdlg=false
				}
			},
			showwhich(index){
				if(index==0){
					this.tablename='user'
				}
				else if(index==1){
					this.tablename='member'
				}else if(index==2){
					this.tablename='ac'
				}
				else if(index==3){
					this.tablename='demo'
				}else if(index==4){
					this.tablename='else'
				}
				else if(index==5){
					this.tablename='news'
				}
			}
		},
		components:{
			Topnav,
			Usertabla,
			Membertable,
			Actable,
			Demotable,
			Elsetable,
			Newstable,
			Newuserdlg,
			Newmemberdlg,
			Newacdlg,
			Newdemodlg,
			Newpcdlg,
			Newnewsdlg
		}
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/adminhome.css");
</style>
