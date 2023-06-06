<template>
	<div class="register">
		<div class="registertop"><span class="registerlogo"><el-icon style="position: relative;top: 5px;"><Cpu /></el-icon>BistuLab注册</span></div>
	</div>
	<div class="registermain">
		<div class="registerform">
			<div class="registertip">
				<h3>账号登录</h3>
			</div>
				<el-form :model="formdata" label-width="100px" :rules="registerrule"  ref="ruleFormRef">
					<el-form-item label="用户名" prop="name">
						<el-input v-model="formdata.name" prefix-icon="User" placeholder="请输入用户名"></el-input>
					</el-form-item>
					
					<el-form-item label="密码" prop="password">
						<el-input v-model="formdata.password" prefix-icon="Lock" type="password" show-password="true" placeholder="请输入密码"></el-input>
					</el-form-item>
					<el-form-item label="确认密码" prop="checkpwd">
						<el-input v-model="formdata.checkpwd" prefix-icon="Lock" type="password" show-password="true" placeholder="请确认密码"></el-input>
					</el-form-item>
					<el-form-item label="验证码" prop="verifycode">
						<el-input class="verifyinput" v-model="formdata.verifycode" placeholder="请输入验证码"></el-input>
						<div class="verifypicture" @click="randomcode()">
						<span v-for="(s,index) in nowstrcode.length" :key="s" :ref="'code'+index">
							{{nowstrcode[s-1]}}
						</span>
						</div>
					</el-form-item>
					<el-form-item>
						<el-button @click="register()" >注册</el-button>
					</el-form-item>
					<div class="registertologin">
						<span  @click="tologin()"><a>已有账号？立即登录</a></span>
					</div>
				</el-form>
		</div>
	</div>
</template>

<script>
			// randomcode()
	export default {
		name:'RegisterView',
		data:function(){
			return {
				nowstrcode:'wendy',
				color:['brown','aqua','bisque','darkblue','skyblue','black','red','darkgoldenrod','darkviolet','deeppink'],
				decoration:["none", "underline", "line-through", "overline"],
				confirmcodestr:"0123456789zxcvbnmasdfghjklqwertyuiop",
				registerrule:{
					name:[{
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
						  } else if (value !== this.formdata.password) {
							callback(new Error("密码不一致！！！"))
						  } else {
							callback()
						  }
						},
						trigger:'blur',
					}],
					verifycode:[{
						validator:(rule, value, callback) => {
						  if (value === '') {
							callback(new Error('请输入验证码！！！'))
						  } else if (value !== this.nowstrcode) {
							callback(new Error("验证码错误！！！"))
						  } else {
							callback()
						  }
						},
						trigger:'blur',
					}]
				},
				formdata:{
					name:'',
					password:'',
					checkpwd:'',
					verifycode:'',
				},
			}
		},
		methods:{
			randomcode(){
				this.nowstrcode='';
				for(var i =0;i<4;i++){
					var randomindex = parseInt(Math.random()*36);
					this.nowstrcode+=this.confirmcodestr.substring(randomindex,randomindex+1);
					// this.$refs[`code${i}`][0].style.fontSize=parseInt(Math.random() *10)+20+"px";
					// this.$refs[`code${i}`][0].style.verticalAlign=parseInt(Math.random() *50)+"%";
					// this.$refs[`code${i}`][0].style.textDecoration=this.decoration[parseInt(Math.random() * 4)];
					// this.$refs[`code${i}`][0].style.color=this.color[parseInt(Math.random() * 10)];;
					// this.$refs[`code${i}`][0].style.fontWeight=parseInt(Math.random() *8)*100;
					// this.$refs[`code${i}`][0].style.transform="rotate("+(parseInt(Math.random() *20)-10)+"deg)";
				}
				for(var j=0;j<4;j++){
					this.$refs[`code${j}`][0].style.fontSize=parseInt(Math.random() *10)+20+"px";
					this.$refs[`code${j}`][0].style.verticalAlign=parseInt(Math.random() *50)+"%";
					this.$refs[`code${j}`][0].style.textDecoration=this.decoration[parseInt(Math.random() * 4)];
					this.$refs[`code${j}`][0].style.color=this.color[parseInt(Math.random() * 10)];;
					this.$refs[`code${j}`][0].style.fontWeight=parseInt(Math.random() *8)*100;
					this.$refs[`code${j}`][0].style.transform="rotate("+(parseInt(Math.random() *20)-10)+"deg)";
				}
				// this.change()
			},
			change(){
				for(var j=0;j<4;j++){
					this.$refs[`code${j}`][0].style.fontSize=parseInt(Math.random() *10)+20+"px";
					this.$refs[`code${j}`][0].style.verticalAlign=parseInt(Math.random() *50)+"%";
					this.$refs[`code${j}`][0].style.textDecoration=this.decoration[parseInt(Math.random() * 4)];
					this.$refs[`code${j}`][0].style.color=this.color[parseInt(Math.random() * 10)];;
					this.$refs[`code${j}`][0].style.fontWeight=parseInt(Math.random() *8)*100;
					this.$refs[`code${j}`][0].style.transform="rotate("+(parseInt(Math.random() *20)-10)+"deg)";
				}
			},
			register(){
				
			},
			tologin(){
				this.$router.push('./login')
				// this.randomcode();
				// alert(this.nowstrcode)
			},
		},
		mounted() {
			this.$nextTick(()=>{
				this.randomcode();
				this.change()
				// this.randomcode();
			});
		},
	}
			// randomcode()
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/register.css");
</style>
