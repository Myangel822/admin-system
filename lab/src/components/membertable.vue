<template>
	<el-table :data="filterdata()" style="width: 100%" class="usertablecss">
	    <el-table-column label="职位" prop="identity" />
	    <el-table-column label="姓名" prop="name" />
		<el-table-column label="联系方式" prop="phone" />
	    <el-table-column align="right">
	      <template #header>
			  <el-button class="newonebtn" @click="shownew()">新增成员</el-button>
	        <el-input v-model="searchvalue" size="small" placeholder="搜索成员..." />
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
		name:'MemberTable',
		data:function(){
			return{
				searchvalue:'',
				memberdata:[
		
				]
			}
		},
		created(){
			this.getallmember()
		},
		methods:{
			getallmember(){
				API.MemberAPI.getmembers().then( res=> {
					this.memberdata=res.data.data
				})
			},
			shownew(){
				this.$emit('shownewdlg')
			},
			edit(i,j){
				
			},
			deleteone(i,j){
				console.log(6)
				API.MemberAPI.deleteone(this.filterdata()[i].id)
			},
			filterdata(){
			if(this.searchvalue !=''){
				var reg = new RegExp(this.searchvalue);
				var arr=[];
				for(var i=0;i<this.memberdata.length;i++){
					if(reg.test(this.memberdata[i].name)){
						arr.push(this.memberdata[i])
					}
				}
				return arr;
			}else{
				return this.memberdata;
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

