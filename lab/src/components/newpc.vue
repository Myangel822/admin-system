<template>
	  <el-dialog title="新增相片" :model-value="shownewpcdlg" :visible.sync="shownewpcdlg" @close="closenewdlg()" align-center="center" >
		  <template #footer>
		    <span class="dialog-footer">
		      <el-button @click="closenewdlg()">取消</el-button>
		      <el-button type="primary" @click="closenewdlg()">
		        确认
		      </el-button>
		    </span>
		  </template>
	  </el-dialog>
</template>

<script>
	export default {
		name:'Newpcdlg',
		props:{
			shownewpcdlg:{
				type:Boolean,
				default:false,
				required:true
			}
		},
		data() {
			return {
			}
		},
		methods: {
			closenewdlg() {
				this.$emit('closenewdlg')
			}
		},
	}
</script>

<style>
</style>
