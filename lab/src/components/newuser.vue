<template>
	  <el-dialog title="新增用户" width="30%" :model-value="shownewuserdlg" :visible.sync="shownewuserdlg" @close="closenewdlg()" align-center="center">
		   <el-form :model="userdata" label-width="100px" :rules="registerrule"  ref="ruleFormRef">
		   	<el-form-item label="用户名" prop="username">
		   		<el-input v-model="userdata.username" prefix-icon="User" placeholder="请输入用户名"></el-input>
		   	</el-form-item>
		   	<el-divider></el-divider>
		   	<el-form-item label="密码" prop="password">
		   		<el-input v-model="userdata.password" prefix-icon="Lock" type="password" show-password="true" placeholder="请输入密码"></el-input>
		   	</el-form-item>
			<el-divider></el-divider>
		   	<el-form-item label="确认密码" prop="checkpwd">
		   		<el-input v-model="this.checkpwd" prefix-icon="Lock" type="password" show-password="true" placeholder="请确认密码"></el-input>
		   	</el-form-item>
			<el-divider></el-divider>
		   </el-form>
		    <template #footer>
		      <span class="dialog-footer">
		        <el-button @click="closenewdlg()">取消</el-button>
		        <el-button type="primary" @click="savenewuser()">
		          确认
		        </el-button>
		      </span>
		    </template>
	  </el-dialog>
</template>

<script>
	export default {
		name:'Newuserdlg',
		props:{
			shownewuserdlg:{
				type:Boolean,
				default:false,
				required:true
			}
		},
		data() {
			return {
				userdata:{
					id:'',
					username:'',
					password:'',
				},
				checkpwd:'',
				registerrule:{
					username:[{
						required:true,
						trigger:'blur',
						message:'用户名不能为空！！！',// 密码强度正则，最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符var passwordReg =/^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$/;// 匹配邮箱地址var mailReg =/\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/;// 中国六位数的邮政编码var postalCode =/^\d{6}$/;// 匹配15~18位身份证var IDCard =/(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)/;// 匹配18位的新版身份证var IDCard_18 =/^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/;
					},
					{
						min:4,
						max:16,
						message:'用户名必须4-16位！！！',
						trigger:'blur'
					}],
					password:[{
						required:true,
						trigger:'blur',
						message:'密码不能为空！！！',
					},{
						min:6,
						max:16,
						trigger:'blur',
						message:'密码必须6-16位！！！',
					},{
						trigger:'blur',
						message:'密码包括大小写字母，数字和特殊字符！！！',
						pattern:/^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#-*/+$%^&*? ]).*$/,
					}
					],
					checkpwd:[{
						validator:(rule, value, callback) => {
						  if (value === '') {
							callback(new Error('请确认密码！！！'))
						  } else if (value !== this.userdata.password) {
							callback(new Error("密码不一致！！！"))
						  } else {
							callback()
						  }
						},
						trigger:'blur',
					}],
					}
			}
		},
		methods: {
			closenewdlg() {
				this.$emit('closenewdlg')
			},
			savenewuser(){
				console.log(this.userdata)
				this.$emit('closenewdlg')
			}
		},
	}
</script>

<style scoped="scoped">
.el-form-item__content {
	border: #DCDCDC solid 1px;
}
</style>
