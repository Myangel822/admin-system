package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"fmt"
	"strconv"

	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

func GetStandardsByName(c *gin.Context) {
	name := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetStandardsByName(name)
			data["total"] = models.CountStandardsByName(name)
			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.CountStandardsByNameAndUid(member.ID, name)

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

func GetStandards(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistStandardsByID(id) {
			data = models.GetStandards(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_Standards
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

func GetStandardss(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetStandardss(util.GetPage(c), models.GetStandardsTotal(maps), maps)
		data["total"] = models.GetStandardsTotal(maps)
	} else {
		data["lists"] = models.GetStandardsByUid(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetStandardssForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetStandardss(util.GetPage(c), models.GetStandardsTotal(maps), maps)
	data["total"] = models.GetStandardsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddStandards(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	kind := c.Query("kind")
	date := c.Query("date")
	date = string([]byte(date)[:10])
	state := c.Query("state")
	id := 0
	uid := GetCurrentUser(c).ID
	valid := validation.Validation{}
	valid.Required(title, "title").Message("技术标准名称不能为空")
	valid.MaxSize(title, 100, "title").Message("技术标准最长为100字符")
	valid.Required(content, "content").Message("技术标准内容不能为空")
	valid.MaxSize(content, 100, "content").Message("技术标准内容最长为100字符")
	valid.Required(date, "theyear").Message("技术标准时间不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistStandardsByTitle(id, title) {

			data := make(map[string]interface{})
			data["title"] = title
			data["content"] = content
			data["date"] = date
			data["kind"] = kind
			data["state"] = state
			data["uid"] = uid
			models.AddStandards(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_Standards
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditStandards(c *gin.Context) {
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
	kind := c.Query("kind")
	date := c.Query("date")
	if date != "" {
		date = string([]byte(date)[:10])
	}

	state := c.Query("state")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistStandardsByID(id) {

			data := make(map[string]interface{})
			if content != "" {
				data["content"] = content
			}
			if title != "" {
				data["title"] = title
			}
			data["kind"] = kind
			data["date"] = date
			data["state"] = state
			models.EditStandards(id, content, title, kind, date, state)

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

func EditStandardState(c *gin.Context) {
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
		if models.ExistStandardsByID(id) {

			models.EditStandardsState(id, state)

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

func DeleteStandards(c *gin.Context) {
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
		if models.ExistStandardsByID(id) {
			models.DeleteStandards(id)
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
