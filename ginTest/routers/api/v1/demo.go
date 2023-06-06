package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

func GetDemoByName(c *gin.Context) {
	title := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(title, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetDemoByName(title)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetDemoByNameAndUid(member.ID, title)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
func GetDemo(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetDemos(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetDemosTotal(maps)
	} else {
		data["lists"] = models.GetDemoByUid(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetDemoForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetDemos(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetDemosTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddDemo(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	date := c.PostForm("date")
	date = date[:10]
	video, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}

	video.Filename = "/bishe/ginTest/dist/" + time.Now().Format("2006-01-02_150405") + ".mp4"

	c.SaveUploadedFile(video, video.Filename)

	video.Filename = "/bishe/lab/src/assets/" + time.Now().Format("2006-01-02_150405") + ".mp4"
	c.SaveUploadedFile(video, video.Filename)

	videoUrl := video.Filename[22:]
	valid := validation.Validation{}
	id := 0
	uid := GetCurrentUser(c).ID
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistDemoByTitle(id, title) {
			code = e.SUCCESS
			models.AddDemo(title, content, date, videoUrl, uid)
		} else {
			code = e.ERROR_EXIST_Demo
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditDemo(c *gin.Context) {
	id1 := c.PostForm("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	title := c.PostForm("title")
	content := c.PostForm("content")
	date := c.PostForm("date")

	if date != "" {
		date = string([]byte(date)[:10])
	}
	video, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}

	video.Filename = "/bishe/ginTest/dist/" + time.Now().Format("2006-01-02_150405") + ".mp4"

	c.SaveUploadedFile(video, video.Filename)

	video.Filename = "/bishe/lab/src/assets/" + time.Now().Format("2006-01-02_150405") + ".mp4"
	c.SaveUploadedFile(video, video.Filename)

	videoUrl := video.Filename[22:]
	valid := validation.Validation{}

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistDemoByID(id) {

			data := make(map[string]interface{})
			if content != "" {
				data["content"] = content
			}
			if title != "" {
				data["title"] = title
			}
			data["video"] = videoUrl
			data["date"] = date
			models.EditDemo(id, content, title, date, videoUrl)

		} else {
			code = e.ERROR_NOT_EXIST_Demo
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteDemo(c *gin.Context) {
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistDemoByID(id) {
			models.DeleteDemo(id)
		} else {
			code = e.ERROR_NOT_EXIST_Demo
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
