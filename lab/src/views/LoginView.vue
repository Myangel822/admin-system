<template>
	<div class="login">
		<div class="logintop"><span class="loginlogo"><el-icon style="position: relative;top: 5px;"><Cpu /></el-icon>BistuLab登录</span></div>
		<div class="loginmain">
			<div class="loginform">
				<div class="logintip">
					<h3>账号登录</h3>
				</div>
				<div class="logininput">
					<span class="inputicon"><el-icon><User /></el-icon></span>
					<span class="inputinput"><el-input placeholder="请输入账号" v-model="loginid"></el-input></span>
				</div>
				<div class="logininput">
					<span class="inputicon"><el-icon><Lock /></el-icon></span>
					<span class="inputinput"><el-input placeholder="请输入密码" v-model="loginpwd" type="password" show-password="true"></el-input></span>
				</div>
				<div class="loginbutton">
					<el-button @click="login()"> 登录 </el-button>
				</div>
				<div class="logintoregister">
					<span  @click="toregister()"><a>还没账号？马上注册</a></span>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	// import Footer from '../components/footer.vue'
	import { login, logout, getInfo } from '@/api/user'
	export default {
		name:'LoginView',
		data:function(){
			return{
				loginid:"",
				loginpwd:"",
			}
		},
		methods:{
			toregister(){
				this.$router.push('./register')
			},
			login(){
				var loginForm = {
					username:this.loginid,
					password:this.loginpwd
				}
				this.loading = true
				login(loginForm).then(res => {
					this.$router.push( './home' )
					const {data} = res
					// commit('SET_TOKEN', data.token)
					})
				    
				}
				
		},
		components:{
			// Footer,
		}
	}
	
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/login.css");
</style>
