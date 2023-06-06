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

func GetAchievementByName(c *gin.Context) {
	name := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetAchievementByName(name)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetAchievementByNameAndUid(member.ID, name)
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
func GetAchievement(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistAchievementByID(id) {
			data = models.GetAchievement(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
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

func GetAchievementsForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetAchievements(util.GetPage(c), models.GetAchievementTotal(maps), maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetAchievements(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetAchievements(util.GetPage(c), models.GetAchievementTotal(maps), maps)
		data["total"] = models.GetAchievementTotal(maps)
	} else {
		data["lists"] = models.GetAchievementByUid(member.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddAchievement(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	kind := c.Query("kind")
	date := c.Query("date")
	date = string([]byte(date)[:10])
	id := 0
	uid := GetCurrentUser(c).ID
	valid := validation.Validation{}
	valid.Required(title, "name").Message("名称不能为空")
	valid.MaxSize(title, 100, "name").Message("名称最长为100字符")
	valid.Required(kind, "category").Message("类别不能为空")
	valid.MaxSize(kind, 100, "category").Message("类别最长为100字符")
	valid.Required(content, "address").Message("地址不能为空")
	valid.MaxSize(content, 100, "address").Message("地址最长为100字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistAchievementByName(id, title) {
			code = e.SUCCESS
			models.AddAchievement(title, content, kind, date, uid)
		} else {
			code = e.ERROR_EXIST_Achievement
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditAchievement(c *gin.Context) {
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
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "name").Message("名称最长为100字符")
	valid.MaxSize(content, 100, "address").Message("地址最长为100字符")
	valid.MaxSize(kind, 100, "category").Message("类别最长为100字符")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistAchievementByID(id) {

			// data := make(map[string]interface{})
			// if content != "" {
			// 	data["content"] = content
			// }
			// if title != "" {
			// 	data["title"] = title
			// }
			// if kind != "" {
			// 	data["category"] = kind
			// }
			// data["date"] = date
			models.EditAchievement(id, content, title, kind, date)

		} else {
			code = e.ERROR_NOT_EXIST_Achievement
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteAchievement(c *gin.Context) {
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
		if models.ExistAchievementByID(id) {
			models.DeleteAchievement(id)
		} else {
			code = e.ERROR_NOT_EXIST_Achievement
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
