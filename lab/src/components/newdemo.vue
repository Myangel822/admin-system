<template>
	  <el-dialog title="新增视频" :model-value="shownewdemodlg" :visible.sync="shownewdemodlg" @close="closenewdlg()" align-center="center" >
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
		name:'Newdemodlg',
		props:{
			shownewdemodlg:{
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
