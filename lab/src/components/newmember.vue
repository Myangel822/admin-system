<template>
	  <el-dialog title="新增成员" :model-value="shownewmemberdlg" :visible.sync="shownewmemberdlg" @close="closenewdlg()" align-center="center" >
			<el-form :model="memberdata" label-width="100px"  ref="ruleFormRef">
				<el-form-item label="用户名" >
					<el-input v-model="memberdata.name" prefix-icon="User" placeholder="请输入姓名"></el-input>
				</el-form-item>
				<el-divider></el-divider>
				<el-form-item label="个人照片" style="height: auto;">
					<el-upload
								action=""
					           list-type="picture"
					           :auto-upload="false"
					           :multiple="false"
							   :on-change="uploadpc"
					           :limit="1"
							   
					   class="avatar-uploader"
					 >
					   <img v-if="imageUrl" :src="require('@/assets/'+imageUrl+'.jpg')" class="avatar" />
					   <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
					 </el-upload>
					 <!-- {{this.imageUrl}} -->
				</el-form-item>
				<el-divider></el-divider>
				<el-form-item label="类型">
					 <el-radio-group v-model="memberdata.type" class="ml-4">
					      <el-radio label="0" size="large">负责人</el-radio>
					      <el-radio label="1" size="large">教师</el-radio>
						  <el-radio label="2" size="large">在校学生</el-radio>
						  <el-radio label="3" size="large">已毕业学生</el-radio>
					    </el-radio-group>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="职位">
					<el-input v-model="memberdata.identity" prefix-icon="User" placeholder="请输入职位"></el-input>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="手机">
					<el-input v-model="memberdata.phone" prefix-icon="User" placeholder="请输入手机"></el-input>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="邮箱">
					<el-input v-model="memberdata.mail" prefix-icon="User" placeholder="请输入邮箱"></el-input>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="个人自述">
					<el-input v-model="memberdata.introduction" prefix-icon="User" placeholder="个人自述"></el-input>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="研究方向">
					<el-input v-model="memberdata.research" prefix-icon="User" placeholder="研究方向"></el-input>
				</el-form-item>
						<el-divider></el-divider>
				<el-form-item label="个人成果">
					<el-input v-model="memberdata.achievement" prefix-icon="User" placeholder="个人成果"></el-input>
				</el-form-item>
						<el-divider></el-divider>
			</el-form>
		 <template #footer>
		   <span class="dialog-footer">
		     <el-button @click="closenewdlg()">取消</el-button>
		     <el-button type="primary" @click="savenewmember()">
		       确认
		     </el-button>
		   </span>
		 </template>
	  </el-dialog>
</template>

<script>
	export default {
		name:'Newumemberdlg',
		props:{
			shownewmemberdlg:{
				type:Boolean,
				default:false,
				required:true
			}
		},
		data() {
			return {
				imageUrl:'',
				memberdata:{
					id:'',
					type:'',
					picture:'',
					identity:'',
					name:'',
					phone:'',
					mail:'',
					research:'',
					achievement:'',
					introduction:'',
					state:'可用',
				},
			}
		},
		methods: {
			closenewdlg() {
				this.$emit('closenewdlg')
			},
			savenewmember(){
				console.log(this.imageUrl)
				this.memberdata.picture=this.imageUrl
				this.$emit('closenewdlg')
			},
			uploadpc(file,filelists){
				this.imageUrl=filelists[0].name.split('.')[0];
			}
		},
	}
</script>

<style scoped="scoped">
	.avatar-uploader .avatar {
	  width: 178px;
	  height: 178px;
	  border: 2px dashed seashell;
	  display: block;
	}
	.avatar-uploader .el-upload {
	  border: 1px dashed seashell;
	  border-radius: 6px;
	  cursor: pointer;
	  position: relative;
	  overflow: hidden;
	  /* transition: var(--el-transition-duration-fast); */
	}
	.avatar-uploader .el-upload:hover {
	  border-color: skyblue;
	}
	
	.el-icon.avatar-uploader-icon {
	  font-size: 28px;
	  color: #8c939d;
	  width: 178px;
	  height: 178px;
	  border: 2px dashed seashell;
	  text-align: center;
	}
</style>
