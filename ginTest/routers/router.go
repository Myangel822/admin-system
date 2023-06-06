package routers

//注册路由
import (
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/EDDYCJY/go-gin-example/enforcer"
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/middleware/kuayu"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Cookie struct {
	Cookie string `json:"cookie"`
}

func interceptor(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		if find := strings.Contains(obj, "?"); find {
			i := strings.Index(obj, "?")
			obj = obj[0:i]
		}
		cookie := c.Request.Header.Get("Cookies")
		c.Request.Header.Add("Cookie", cookie)
		sub := v1.GetCurrentUser(c).Name
		//
		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(kuayu.Cors())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	// 渲染前端渲染静态文件模板
	r.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	r.LoadHTMLFiles("dist/*/*")          // 定义资源路径
	r.Static("/static/dist", "./dist/")  // 添加资源路径
	r.StaticFile("/", "dist/index.html") // 添加前端接口

	store := cookie.NewStore([]byte("secret")) // 设置生成sessionId的密钥
	// mysession是返回給前端的sessionId名
	r.Use(sessions.Sessions("mysession", store))

	//获取token
	r.GET("/manager", api.Login)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	apiv1.Use(interceptor(enforcer.EnforcerTool()))
	gob.Register(models.Member{})
	{
		//新建新闻
		apiv1.POST("/news", v1.AddNews)
		//更新指定新闻
		apiv1.PUT("/editNews", v1.EditNews)
		//删除指定新闻
		apiv1.DELETE("/deleteNews", v1.DeleteNews)

		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/editArticle", v1.EditArticle)

		apiv1.PUT("/editArticleState", v1.EditArticleState)
		apiv1.PUT("/editArticleHide", v1.EditArticleHide)
		//删除指定文章
		apiv1.DELETE("/deleteArticle", v1.DeleteArticle)

		//新建项目
		apiv1.POST("/project", v1.AddProject)
		//更新指定项目
		apiv1.PUT("/editProject", v1.EditProject)
		apiv1.PUT("/editProjectState", v1.EditProjectState)
		//删除指定项目
		apiv1.DELETE("/deleteProject", v1.DeleteProject)

		//新建demo
		apiv1.POST("/demo", v1.AddDemo)
		//更新指定demo
		apiv1.PUT("/editDemo", v1.EditDemo)
		//删除指定demo
		apiv1.DELETE("/deleteDemo", v1.DeleteDemo)

		//新建技术标准
		apiv1.POST("/standards", v1.AddStandards)
		//更新指定技术标准
		apiv1.PUT("/editStandard", v1.EditStandards)
		//改变技术标准状态
		apiv1.PUT(("/editStandardState"), v1.EditStandardState)
		//删除指定技术标准
		apiv1.DELETE("/deleteStandard", v1.DeleteStandards)

		//新建活动
		apiv1.POST("/activity", v1.AddActivity)
		//更新指定活动
		apiv1.PUT("/editActivity", v1.EditActivity)
		//删除指定活动
		apiv1.DELETE("/deleteActivity", v1.DeleteActivity)

		//新建成员
		apiv1.POST("/member", v1.AddMember)
		//更新指定成员
		apiv1.PUT("/editMember", v1.EditMember)
		//重置成员密码
		apiv1.PUT("/resetPassword", v1.ResetPassword)
		//删除指定成员
		apiv1.DELETE("/deleteMember", v1.DeleteMember)

		//新建图片
		apiv1.POST("/image", v1.AddImage)
		//更新指定图片
		apiv1.PUT("/editImage", v1.EditImage)
		//删除指定图片
		apiv1.DELETE("/deleteImage", v1.DeleteImage)

		//新建成果
		apiv1.POST("/achievement", v1.AddAchievement)
		//更新指定成果
		apiv1.PUT("/editAchievement", v1.EditAchievement)
		//删除指定成果
		apiv1.DELETE("/deleteAchievement", v1.DeleteAchievement)
	}
	apiget := r.Group("/api/get")
	apiget.Use(interceptor(enforcer.EnforcerTool()))
	{
		//获取新闻列表
		apiget.GET("/news", v1.GetNews)
		apiget.GET("/mohuNews", v1.GetNewsByName)
		apiget.GET("/getnewsForword", v1.GetNewsForword)
		//获取项目列表
		apiget.GET("/project", v1.GetProjects)
		apiget.GET("/mohuProject", v1.GetProjectByName)
		apiget.GET("/countProject", v1.CountProject)
		apiget.GET("/getProjectsForword", v1.GetProjectsForword)
		//获取文章列表
		apiget.GET("/articles", v1.GetArticles)
		apiget.GET("/mohuArticle", v1.GetArticleByName)
		apiget.GET("/countArticle", v1.CountArticle)
		apiget.GET("/articleforword", v1.GetArticleForword)
		//获取成员列表
		apiget.GET("/members", v1.GetMembers)
		apiget.GET("/mohuMember", v1.GetMemberByName)
		apiget.GET("/countMember", v1.CountMembers)
		apiget.GET("/getMembersForword", v1.GetMembersForword)
		//获取图片列表
		apiget.GET("/images", v1.GetImages)
		apiget.GET("/mohuImages", v1.GetImageByDate)
		apiget.GET("/getImagesForword", v1.GetImagesForword)
		//获取成果列表
		apiget.GET("/achievements", v1.GetAchievements)
		apiget.GET("/mohuAchievement", v1.GetAchievementByName)
		apiget.GET("/achievementsForword", v1.GetAchievementsForword)
		//获取技术标准列表
		apiget.GET("/standards", v1.GetStandardss)
		//模糊查询技术标准
		apiget.GET("/mohuStandards", v1.GetStandardsByName)
		apiget.GET("/getStandardsForword", v1.GetStandardssForword)
		//获取活动列表
		apiget.GET(("/activity"), v1.GetActivitys)
		apiget.GET("/mohuActivity", v1.GetActivityByName)
		apiget.GET("/getActivityByKind", v1.GetActivityByKind)
		apiget.GET("/getActivitysForword", v1.GetActivitysForword)
		//获取Demo列表
		apiget.GET("/demo", v1.GetDemo)
		apiget.GET("/mohuDemo", v1.GetDemoByName)
		apiget.GET("/getDemoForword", v1.GetDemoForword)
		//获取指定文章
		apiget.GET("/article/:id", v1.GetArticle)
		//获取指定项目
		apiget.GET("/project/:id", v1.GetProject)
		//获取指定成员
		apiget.GET("/member/:id", v1.GetMember)
		//获取指定图片
		apiget.GET("/image/:id", v1.GetImage)
		//获取指定成果
		apiget.GET("/achievement/:id", v1.GetAchievement)
		//获取指定技术标准
		apiget.GET("/standards/:id", v1.GetStandards)
	}
	return r
}
