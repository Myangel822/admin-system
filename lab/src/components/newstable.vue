<template>
	<el-table :data="filterdata()" style="width: 100%" class="usertablecss">
	    <el-table-column label="标题" prop="title" show-overflow-tooltip="true"/>
		<el-table-column label="内容" prop="content" show-overflow-tooltip="true"/>
	    <el-table-column align="right">
	      <template #header>
			  <el-button class="newonebtn" @click="shownew()">新增新闻</el-button>
	        <el-input v-model="searchvalue" size="small" placeholder="搜索新闻..." />
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
		name:'NewsTable',
		data:function(){
			return{
				searchvalue:'',
				news:[
			
				]
			}
		},
		created(){
			this.getallnews()
		},
		methods:{
			getallnews(){
				API.NewsAPI.getnews().then(res=>{
					this.news=res.data.data;
				})
			},
			shownew(){
				this.$emit('shownewdlg')
			},
			edit(i,j){
				
			},
			deleteone(i,j){
				console.log(6)
				API.NewsAPI.deleteone(this.filterdata()[i].id)
			},
			filterdata(){
			if(this.searchvalue !=''){
				var reg = new RegExp(this.searchvalue);
				var arr=[];
				for(var i=0;i<this.news.length;i++){
					if(reg.test(this.news[i].name)){
						arr.push(this.news[i])
					}
				}
				return arr;
			}else{
				return this.news;
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
