package v1

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type Activity struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Kind     string   `json:"kind"`
	Date     string   `json:"date"`
	ImageUrl []string `json:"imageUrl"`
}

func GetActivityByKind(c *gin.Context) {
	kind := c.Query("kind")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(kind, "kind").Message("kind不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetActivityByKind(kind)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetActivityByKindAndUid(member.ID, kind)

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

func GetActivityByName(c *gin.Context) {
	title := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(title, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetActivityByName(title)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetActivityByNameAndUid(member.ID, title)

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
func GetActivity(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistActivityByID(id) {
			data = models.GetActivity(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_Activity
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

func GetActivitysForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetActivitys(util.GetPage(c), models.GetActivityTotal(maps), maps)
	data["total"] = models.GetActivityTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetActivitys(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetActivitys(util.GetPage(c), models.GetActivityTotal(maps), maps)
		data["total"] = models.GetActivityTotal(maps)
	} else {
		data["lists"] = models.GetActivityByUid(member.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddActivity(c *gin.Context) {
	json := Activity{}
	c.BindJSON(&json)
	title := json.Title
	content := json.Content
	kind := json.Kind
	date := json.Date
	date = string([]byte(date)[:10])
	imageUrl := json.ImageUrl
	imageUrl1 := ""
	uid := GetCurrentUser(c).ID
	if len(imageUrl) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range imageUrl {
			if img[11] == 'j' {
				img = img[23:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.jpg", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.jpg", time.Now().Format("2006-01-02_150405"), index)
			} else if img[11] == 'p' {
				img = img[22:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.png", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.png", time.Now().Format("2006-01-02_150405"), index)
			} else if img[11] == 'g' {
				img = img[22:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.gif", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.gif", time.Now().Format("2006-01-02_150405"), index)
			} else {
				fmt.Printf("不支持该文件类型")
			}

			data, err := enc.DecodeString(img)
			if err != nil {
				fmt.Printf(err.Error())
			}

			//图片保存到文件服务器
			f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(data)

			f1, err2 := os.OpenFile(path1, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err2 != nil {
				fmt.Println(err2.Error())
			}
			defer f1.Close()
			f1.Write(data)
			//保存图片地址到数据库
			path = path[22:]
			imageUrl1 = imageUrl1 + path + ","

		}
	}
	imageUrl1 = strings.TrimRight(imageUrl1, ",")
	id := 0

	valid := validation.Validation{}
	valid.Required(title, "title").Message("活动名称不能为空")
	valid.MaxSize(title, 100, "title").Message("活动名称最长为100字符")
	valid.Required(content, "content").Message("活动链接不能为空")
	valid.MaxSize(content, 100, "content").Message("活动内容最长为100字符")
	valid.Required(date, "date").Message("活动时间不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistActivityByTitle(id, title) {
			code = e.SUCCESS
			models.AddActivity(title, date, content, kind, imageUrl1, uid)
		} else {
			code = e.ERROR_EXIST_Activity
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditActivity(c *gin.Context) {

	json := Activity{}
	c.BindJSON(&json)
	id1 := json.ID
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	title := json.Title
	content := json.Content
	kind := json.Kind
	date := json.Date

	if date != "" {
		date = string([]byte(date)[:10])
	}
	imageUrl := json.ImageUrl
	imageUrl1 := ""
	if len(imageUrl) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range imageUrl {
			if img[11] == 'j' {
				img = img[23:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.jpg", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.jpg", time.Now().Format("2006-01-02_150405"), index)
			} else if img[11] == 'p' {
				img = img[22:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.png", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.png", time.Now().Format("2006-01-02_150405"), index)
			} else if img[11] == 'g' {
				img = img[22:]
				path = fmt.Sprintf("E:/bishe/ginTest/dist/%s_%d.gif", time.Now().Format("2006-01-02_150405"), index)
				path1 = fmt.Sprintf("E:/bishe/lab/src/assets/%s_%d.gif", time.Now().Format("2006-01-02_150405"), index)
			} else {
				fmt.Printf("不支持该文件类型")
			}

			data, err := enc.DecodeString(img)
			if err != nil {
				fmt.Printf(err.Error())
			}

			//图片保存到文件服务器
			f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(data)

			f1, err2 := os.OpenFile(path1, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err2 != nil {
				fmt.Println(err2.Error())
			}
			defer f1.Close()
			f1.Write(data)
			//保存图片地址到数据库
			path = path[22:]
			imageUrl1 = imageUrl1 + path + ","

		}
	}
	imageUrl1 = strings.TrimRight(imageUrl1, ",")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistActivityByID(id) {

			// data := make(map[string]interface{})
			// if content != "" {
			// 	data["content"] = content
			// }
			// if title != "" {
			// 	data["title"] = title
			// }
			// data["kind"] = kind
			// data["date"] = date
			// data["imageUrl"] = imageUrl1
			models.EditActivity(id, content, title, kind, date, imageUrl1)

		} else {
			code = e.ERROR_NOT_EXIST_Activity
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteActivity(c *gin.Context) {
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
		if models.ExistActivityByID(id) {
			models.DeleteActivity(id)
		} else {
			code = e.ERROR_NOT_EXIST_Activity
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
