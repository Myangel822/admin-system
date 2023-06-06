package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

func CountProject(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["total"] = models.GetProjectTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetProjectByName(c *gin.Context) {
	name := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetProjectByName(name)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetProjectByNameAndUid(member.ID, name)

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
func GetProject(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}

	if !valid.HasErrors() {
		if models.ExistProjectByID(id) {
			data = models.GetProject(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_Project
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetProjects(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetProjects(util.GetPage(c), models.GetProjectTotal(maps), maps)
		data["total"] = models.GetProjectTotal(maps)
	} else {
		data["lists"] = models.GetProjectByUid(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetProjectsForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetProjects(util.GetPage(c), models.GetProjectTotal(maps), maps)
	data["total"] = models.GetProjectTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddProject(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	state := c.Query("state")
	theyear := c.Query("theyear")
	id := 0
	uid := GetCurrentUser(c).ID
	valid := validation.Validation{}
	valid.Required(title, "name").Message("项目名称不能为空")
	valid.MaxSize(title, 100, "name").Message("项目名称最长为100字符")
	valid.Required(content, "link").Message("项目链接不能为空")
	valid.MaxSize(content, 100, "link").Message("项目链接最长为100字符")
	valid.Required(theyear, "theyear").Message("项目年份不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistProjectByName(id, title) {
			code = e.SUCCESS
			models.AddProject(title, content, state, theyear, uid)
		} else {
			code = e.ERROR_EXIST_Project
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditProject(c *gin.Context) {
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	title := c.Query("title")
	content := c.Query("content")
	state := c.Query("state")
	theyear := c.Query("theyear")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistProjectByID(id) {

			data := make(map[string]interface{})
			if content != "" {
				data["content"] = content
			}
			if title != "" {
				data["title"] = title
			}
			data["state"] = state
			data["theyear"] = theyear
			models.EditProject(id, content, title, state, theyear)

		} else {
			code = e.ERROR_NOT_EXIST_Project
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteProject(c *gin.Context) {
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
		if models.ExistProjectByID(id) {
			models.DeleteProject(id)
		} else {
			code = e.ERROR_NOT_EXIST_Project
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditProjectState(c *gin.Context) {
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)

	state := c.Query("state")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistProjectByID(id) {

			models.EditProjectState(id, state)

		} else {
			code = e.ERROR_NOT_EXIST_Standards
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
