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

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type News struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	ImageUrl []string `json:"imageUrl"`
	Cookie   string   `json:"cookie"`
}

func GetNewsByName(c *gin.Context) {
	title := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(title, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetNewsByName(title)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}

	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetNewsByNameAndUid(member.ID, title)

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

func GetNewsForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetNews(util.GetPage(c), setting.PageSize, maps)

	data["total"] = models.GetNewsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetNews(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetNews(util.GetPage(c), setting.PageSize, maps)

		data["total"] = models.GetNewsTotal(maps)

	} else {
		data["lists"] = models.GetNewsByUid(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddNews(c *gin.Context) {

	json := News{}
	c.BindJSON(&json)
	// log.Printf("%v", &json)

	title := json.Title
	content := json.Content
	// c.Request.Header.Add("Cookie", json.Cookie)

	print(content)
	created_on := time.Now().Format("2006-01-02_15:04:05")
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
			f, err1 := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err1 != nil {
				fmt.Println(err1.Error())
			}

			f1, err2 := os.OpenFile(path1, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err2 != nil {
				fmt.Println(err2.Error())
			}
			defer f.Close()
			f.Write(data)
			//保存图片地址到数据库
			defer f1.Close()
			f1.Write(data)
			path = path[22:]
			imageUrl1 = imageUrl1 + path + ","

		}
	}
	imageUrl1 = strings.TrimRight(imageUrl1, ",")
	valid := validation.Validation{}
	id := 0
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistNewsByTitle(id, title) {
			code = e.SUCCESS
			models.AddNews(title, content, created_on, imageUrl1, uid)
		} else {
			code = e.ERROR_EXIST_News
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func EditNews(c *gin.Context) {

	json := News{}
	c.BindJSON(&json)
	// log.Printf("%v", &json)
	title := json.Title
	content := json.Content
	// c.Request.Header.Add("Cookie", json.Cookie)

	id1 := json.ID
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	print(content)
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
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistNewsByID(id) {

			data := make(map[string]interface{})
			if content != "" {
				data["content"] = content
			}
			if title != "" {
				data["title"] = title
			}
			data["imageUrl"] = imageUrl1

			models.EditNews(id, content, title, imageUrl1)

		} else {
			code = e.ERROR_NOT_EXIST_News
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteNews(c *gin.Context) {
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
		if models.ExistNewsByID(id) {
			models.DeleteNews(id)
		} else {
			code = e.ERROR_NOT_EXIST_News
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
