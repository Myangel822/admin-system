<template>
	<el-table :data="filterdata()" style="width: 100%" class="usertablecss">
	    <el-table-column label="id" prop="id" />
	    <el-table-column label="用户名" prop="username" />
	    <el-table-column align="right">
	      <template #header>
			  <el-button class="newonebtn" @click="shownew()">新增用户</el-button>
	        <el-input v-model="searchvalue" size="small" placeholder="搜索用户..." />
	      </template>
	      <template #default="scope">
	        <el-button size="small" @click="edit(scope.$index, scope.row)"
	          >Edit</el-button
	        >
			 <el-popconfirm
			    confirm-button-text="Yes"
			    cancel-button-text="No"
			    :icon="InfoFilled"
			    icon-color="#626AEF"
			    title="Are you sure to delete this?"
			    @confirm="deleteone(scope.$index, scope.row)"
			  >
			    <template #reference>
			     <el-button
			       size="small"
			       type="danger">Delete</el-button>
			    </template>
			  </el-popconfirm>
	      </template>
	    </el-table-column>
	  </el-table>
</template>

<script>
	import API from '../api/index.js'
	export default {
		name:'UserTable',
		data:function(){
			return{
				searchvalue:'',
				userdata:[
					
				]
			}
		},
		created(){
			this.getalluser()
		},
		methods:{
			getalluser(){
				API.ManagerAPI.getmanager().then( res => {
					this.userdata=res.data.data
				})
			},
			shownew(){
				this.$emit('shownewdlg')
			},
			edit(i,j){
				console.log(6)
			},
			deleteone(i,j){
				console.log(6)
				API.ManagerAPI.deleteone(this.filterdata()[i].id)
			},
			filterdata(){
			if(this.searchvalue !=''){
				var reg = new RegExp(this.searchvalue);
				var arr=[];
				for(var i=0;i<this.userdata.length;i++){
					if(reg.test(this.userdata[i].name)){
						arr.push(this.userdata[i])
					}
				}
				return arr;
			}else{
				return this.userdata;
			}
			}
		}
	}
</script>

<style scoped="scoped">
.usertablecss .el-input{
    border: silver  solid 1px !important;
    /* box-shadow: !important; */
    padding: 0px;
}
</style>
