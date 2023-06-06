<template>
	<ul class="picturecon">
		<li class="pictopfunction">
			<div class="yearsearch">
				<ul>
					<li>年份：</li>
					<li v-for="i in years" @click="changeyear(i)">
						{{i}}
					</li>
				</ul>
			</div>
			<div class="keysearch">
				<el-input clearable="true" v-model="searchvalue" placeholder="请输入关键词" prefix-icon="search"></el-input>
			</div>
		</li>
		<li class="pictureone" v-for="(i,index) in filterpc()" :key="i.content">
			<div class="picpic">
				 <el-image style="width: 100%; height: 100%" :src="require('@/assets/'+i.name)"  :fit="fit" lazy="true" :preview-src-list="srcList"  :initial-index="index"/>
				
			</div>
			<div class="picird">
				<div class="picirdcon">{{i.content}}</div>
				<div class="picirdyear">{{i.date}}</div>
			</div>
		</li>
<!-- 		<li class="bottompage">
		<el-pagination small layout="prev, pager, next" :total="50" background="false" />	
		</li> -->
	</ul>
</template>
<script>
	import API from '../api/index.js'
	export default {
		name:'pageitem',
		data:function(){
			return{
				nowyear:'',
				searchvalue:'',
				years:['全部','2023','2022','2021','2020','2019'],
				srcList:[],
				picitems:[

				]
			}
		},
		created(){
			this.getallimages()
		},
		mounted:function(){
			this.preseepic()
		},
		updated:function(){
			this.preseepic()
		},
		methods:{
			getallimages(){
			API.PictureAPI.getpictures().then(res => {
				this.picitems=res.data.data.lists;
			})
			},
			preseepic(){
				for (var i = 0; i < this.filterpc().length; i++) {
						var temp=require('@/assets/'+this.filterpc()[i].name)
						this.srcList.push(temp);
					}
			 console.log(this.srcList)
			},
			changeyear(y){
				if(y=='全部')
				{
					this.nowyear=''
				}else{
					this.nowyear=y
				}
			},
			filterpc() {
				if (this.searchvalue != '' || this.nowyear !='') {
					var reg = new RegExp(this.searchvalue);
					var arr = [];
					if(this.nowyear !=''){
						for (var i = 0; i < this.picitems.length; i++) {
							
							if (reg.test(this.picitems[i].content)&&this.picitems[i].date==this.nowyear) {
								arr.push(this.picitems[i]);
							}
						}
					}
					else{
						for (var i = 0; i < this.picitems.length; i++) {
							
							if (reg.test(this.picitems[i].content)) {
								arr.push(this.picitems[i]);
							}
						}
					}
					// alert(arr)
					return arr;
				}
				else{
					// alert(this.team[0].members)
					return this.picitems;
				}
			},
		},
	}
</script>

<style scoped="scoped">
	@import url("../css/picturehome.css");
</style>
