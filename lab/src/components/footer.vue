<template>
	<div class="footer">
		<ul v-for="f in data" :key="f.title">
			<font style="font-weight: 700px;">{{f.title}}</font>
			<li v-for="ff in f.childlist" :key="ff">
				{{ff}}
			</li>
		</ul>
		<span class="footertips">@CopyRight wendy <br/>Edited on 4 24, 2023<br/>联系地址：北四环中路</span>
	</div>
</template>

<script>
	export default {
	  name: 'Footer',
	  props: {
	    msg: String
	  },
	  data() {
	  	return{
			data:[
				{
					title:'Lab概况',
					childlist:["实验室简介","已获成就","现有成员","最新咨询","研究方向"]
				},
				{
					title:'Lab服务',
					childlist:["日历","预约实验室","信息服务","后勤管理"]
				},
				{   title:'Lab培养',
					childlist:["导师选择","本科生教育","研究生教育","留学生教育","博士教育"],
				}
			]
		}
	  }
	}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/footer.css");
</style>
