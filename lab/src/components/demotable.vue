<template>
	<el-table :data="filterdata()" style="width: 100%" class="usertablecss">
	    <el-table-column label="名称" prop="title" />
		<el-table-column label="日期" prop="created_on" />
	    <el-table-column align="right">
	      <template #header>
			  <el-button class="newonebtn" @click="shownew()">新增demo</el-button>
	        <el-input v-model="searchvalue" size="small" placeholder="搜索demo..." />
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
		name:'DemoTable',
		data:function(){
			return{
				searchvalue:'',
				demodata:[
				]
			}
		},
		created(){
			this.getalldemo()
		},
		methods:{
			getalldemo(){
				API.DemoAPI.getdemos().then(res=>{
					this.demodata=res.data.data
				})
			},
			shownew(){
				this.$emit('shownewdlg')
			},
			edit(i,j){
				
			},
			deleteone(i,j){
				console.log(6)
				API.DemoAPI.deleteone(this.filterdata()[i].id)
			},
			filterdata(){
			if(this.searchvalue !=''){
				var reg = new RegExp(this.searchvalue);
				var arr=[];
				for(var i=0;i<this.demodata.length;i++){
					if(reg.test(this.demodata[i].name)){
						arr.push(this.demodata[i])
					}
				}
				return arr;
			}else{
				return this.demodata;
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
