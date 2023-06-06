<template>
	<Topnav></Topnav>
	<!-- 主页面 -->
  <div class="home app">
	  <!-- 功能点1：实验室新闻 -->
		<div class="newscrl">
		  <el-carousel indicator-position="none">
		    <el-carousel-item v-for="item in basenews" :key="item.title">
				<img :src="require('@/assets/'+item.pitcture+'.jpg')" />
				<p>{{item.title}}</p>
		    </el-carousel-item>
		  </el-carousel>
		</div>
		  <!-- 功能点2：实验室介绍与研究方向 -->
		  <div class="mainfunction1">
			  <!-- <div style="display: block;border-bottom: #666666 solid 0.1px ;"></div> -->
			  <!-- 实验室介绍 -->
			<div class="labintroduce">
				<div class="left">
					<h3>平台介绍</h3>
					<el-divider></el-divider>
					<p>{{this.labintroduce1}}</p>
					<p>{{this.labintroduce2}}</p>
				</div>
				<div class="right">
					<h3>近期荣誉</h3>
					 <el-carousel :interval="4000" type="card" height="400px" indicator-position="none">
					    <el-carousel-item v-for="item in achievements" :key="item">
					      <img :src="require('@/assets/'+item.pitcture+'.jpg')" />
					      <p style="font-size: 20px; color: black;">{{item.introduce}}</p>
					    </el-carousel-item>
					  </el-carousel>
				</div>
			</div>
			<!-- 实验室研究方向 -->
			<div class="labtarget">
				<h3 style="margin-left: 100px; font-size: 25px;">研究方向</h3>
				<el-divider></el-divider>
				<div class="studypicture" v-for="item in studies" :key="item.name">
					<a :href="item.path"><img :src="require('@/assets/'+item.picture+'.png')"/></a>
					<h5>{{item.name}}</h5>
					<p>{{item.introduce}}</p>
				</div>
			</div>
		  </div>
		  <!-- 功能点1：实验室新闻（续） -->
		  <div class="labnews">
			  <div class="top"><h3 style="color: rgb(6, 63, 152); font-size: 30px;">Lab要闻</h3> <el-button>新闻中心</el-button></div>
			  <!-- 新闻左侧流动图 -->
			  <div class="labnewsleft">
				  <el-carousel indicator-position="none" height="450px">
				    <el-carousel-item v-for="item in basenews.slice(0,3)" :key="item.title">
				  		<div style="height: 450px; width: 500px; ">
							
							<div style="height: 20%; width: 100%;">
							<span class="newsdate">
								<font size="6px" style="position: absolute; top: 6px; left: 16px; font-weight: 700;">08</font>
								<font size="2px" style="position: absolute; bottom: 6px; right: 16px;">2023.03</font>
							</span>
							<span class="newstitle">{{item.title}}</span>
							</div>
								<img :src="require('@/assets/'+item.pitcture+'.jpg')" />>
							
						</div>		
				    </el-carousel-item>
				  </el-carousel>
			  </div>
			  
			  <div class="labnewsmid">
				  <ul>
				  	<li v-for="newsitem in news.slice(0,4)" :key="newsitem.title" @click="shownewsitemdialog(newsitem.id)">
						<div class="listtitle">{{newsitem.title}}</div>
						<div class="listcontent">{{newsitem.content}}</div>
					</li>
				  </ul>
			  </div>
			  <!-- 新闻栏右侧列表 -->
			  <div class="labnewsright">
				  <ul>
				  	<li v-for="newsitem in news.slice(4,8)" :key="newsitem.title" @click="shownewsitemdialog(newsitem.id)">
				  		<div class="listtitle">{{newsitem.title}}</div>
				  		<div class="listcontent">{{newsitem.content}}</div>
				  	</li>
				  </ul>
			  </div>
			  <!-- <div></div> -->
		  </div>
		  <!-- 新闻详情组件 -->
		  <NewsItem  :newsitemdialogflag="newsitemdialogflag" @closenewsitemdialog="closenewsitemdialog" :nownews="nownews"></NewsItem>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import NewsItem from '@/components/newsitem.vue'
import Topnav from '../components/topnav.vue'
import API from '../api/index.js'
export default {
  name: 'HomeView',
  created(){
	  //vue创建时获取数据内容
	  this.getallnews()
  },
  methods:{
	  //获取所有新闻数据
	  getallnews(){
		API.NewsAPI.getnews().then(res => {
			this.news=res.data.data.lists
			console.log(res.data.data)
			console.log(this.news)
		})
	  },
	  //展示具体的新闻详情，通过id查找然后传入
  	  shownewsitemdialog(id){
		  for(var i=0;i<this.news.length;i++)
  		 {
			 if(id==this.news[i].id)
			 {
				 this.nownews=this.news[i];
			 }
		 }
  		  this.newsitemdialogflag=true;
  		  // alert(this.newsitemdialogflag);
  	  },
  	  closenewsitemdialog(){
  		  this.newsitemdialogflag=false;
  	  }
  },
  data:function(){
  	return{
		newsitemdialogflag:false,
		//当前详情新闻内容
		nownews:{
			id:"",
			imageUrl:'',
			title:'',
			created_on:"",
			content:"",
		},
		basenews:[
			{	
				id:"",
				pitcture:'news1',
				title:'实验室成员们积极探讨中！！！',
				created_on:"2000-10-20",
				content:"",
				
			},
			{
				id:"",
				pitcture:'news2',
				title:'ChatGPT爆火启示：软硬科技协同创新正当时',
				created_on:"2000-02-14",
				content:"",
			},
			{
				id:"",
				pitcture:'news3',
				title:'习近平！两会上的暖心发言，值得实验室成员们学习！',
				created_on:"2018-04-01",
				content:"",
			}
		],
		news:[

		],
		achievements:[
			{
				pitcture:'ac1',
				introduce:'顶刊论文',
			},
			{
				pitcture:'ac2',
				introduce:'受邀SCIP做报告',
			},
			{
				pitcture:'ac3',
				introduce:'成果获北京市三等奖',
			},
			{
				pitcture:'ac4',
				introduce:'授权发明性专利',
			},
			{
				pitcture:'ac5',
				introduce:'国际专利',
			}
		],
		studies:[
			{
				name:'知识图谱',
				introduce:'研究知识图谱领域知识本体构建、知识实例抽取、知识关系抽取、知识属性值抽取、知识融合等知识图谱核心关键技术，深度处理各类多源异构海量数据并构建专业领域知识图谱，为各类上层应用（例如智能问答、语义检索、规律挖掘等）提供底层知识图谱数据支撑。',
				picture:'sdy1',
				path:'https://baike.baidu.com/link?url=GjYNTbukLkKG9ZOxyADPXyUHHtXyl4uHZeshcmp3TCzWwHvkggCVqXQs-P2yb-E8XKCRsJLt4fQFox9U6MlHLQlHgVRt337-3qPAh88qe5x-Yqf8xaiET5czJZOHF9oq'
			},
			{
				name:'时序数据分析',
				introduce:'采集并利用海量开源的轨迹数据，通过时序深度学习技术，理解多个目标之间的运动关系，预测感兴趣的目标轨迹，并对目标行为进行多角度的分析和预测。',
				picture:'sdy2',
				path:'https://baike.baidu.com/item/%E6%97%B6%E9%97%B4%E5%BA%8F%E5%88%97%E5%88%86%E6%9E%90/8724605?fr=aladdin'
			},
			{
				name:'计算机视觉',
				introduce:'针对空中拍摄图像，采用图像增强、图像配准等图像处理技术，对于所拍摄到的目标物体进行检测、识别、跟踪等操作，并在此基础上实现对于场景内容的髙语义理解。',
				picture:'sdy3',
				path:'https://baike.baidu.com/item/%E8%AE%A1%E7%AE%97%E6%9C%BA%E8%A7%86%E8%A7%89/2803351?fr=aladdin'
			},
			{
				name:'边缘计算',
				introduce:'针对端边云协同快速发展趋势，探索面向性能优化的多级算力自适应协同机制，实现复杂计算任务所需算力实时动态调配，使端、边、云节点可以进行数据和资源的分布式统一调度',
				picture:'sdy4',
				path:'https://baike.baidu.com/item/%E8%BE%B9%E7%BC%98%E8%AE%A1%E7%AE%97/9044985?fr=aladdin'
			},
		],
		labintroduce1:" 数据科学与情报（信息）分析实验室（Data Science & Intelligence Analysis，DSCI）,"+
		"由高校及企业的专家学者共同组建，包括清华、中科院、北理、北邮在内的多位老师联合指导，其中60%以上的老师具有海外交流/留学经历"+
		"（匹兹堡大学、英国萨里大学、CSIRO）；实验室承担多项国家重点研发计划、国家自然科学基金及国防科研/工程项目，"+
		"在TON、TSC等顶级刊物发表论文多篇。同时实验室与中科院煤转化国家重点实验室、北邮网络技术研究院（国家重点实验室）和匹兹堡大学生物信息学系等国内外知名实验室有着长期学术交流合作。",
		labintroduce2:"实验室致力于情报理论与方法、文本处理、机器视觉处理、机器学习、开源情报、大数据分析等前沿交叉学科技术在政府、产业、企业、军工等领域情报分析决策中的产学研用。实验室紧密跟踪研究国际在计算机科学与情报分析应用领域的前沿技术、热点课题及学科研究发展方向，围绕大数据、人工智能与情报采集、分析、决策等的融合应用等开展项目实践、方法研究、学术交流与人才培养等具体工作，努力将该机构建设成为在国内外情报分析决策应用领域具有重要影响力的智库型研究机构。"
	}
  },
  components: {
	// Footer,
	NewsItem,
    HelloWorld,
	Topnav,
  },
}
</script>

<style scoped="scoped">
	@import url("../css/common.css");
	@import url("../css/mainhome.css");
	@import url("../css/nav.css");
</style>
