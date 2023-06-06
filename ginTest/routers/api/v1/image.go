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
	"github.com/EDDYCJY/go-gin-example/pkg/setting"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type Image struct {
	ID      string   `json:"id"`
	Date    string   `json:"date"`
	Content string   `json:"content"`
	Name    []string `json:"name"`
	Cookie  string   `json:"cookie"`
}

func GetImageByDate(c *gin.Context) {
	date := c.Query("date")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(date, "date").Message("date不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetImageByDate(date)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}

	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetImageByDateAndUid(member.ID, date)

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
func GetImage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistImageByID(id) {
			data = models.GetImage(id)
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

func GetImagesForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetImages(util.GetPage(c), models.GetImageTotal(maps), maps)
	data["total"] = models.GetImageTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetImages(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetImages(util.GetPage(c), setting.PageSize, maps)

		data["total"] = models.GetImageTotal(maps)

	} else {
		data["lists"] = models.GetImageByUid(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddImage(c *gin.Context) {
	json := Image{}
	c.BindJSON(&json)
	// log.Printf("%v", &json)

	name := json.Name
	content := json.Content
	// c.Request.Header.Add("Cookie", json.Cookie)

	print(content)
	date := json.Date

	imageUrl1 := ""
	uid := GetCurrentUser(c).ID
	if len(name) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range name {
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

	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		code = e.SUCCESS
		models.AddImage(imageUrl1, date, content, uid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditImage(c *gin.Context) {

	json := Image{}
	c.BindJSON(&json)
	// log.Printf("%v", &json)
	name := json.Name
	content := json.Content
	// c.Request.Header.Add("Cookie", json.Cookie)
	date := json.Date

	if date != "" {
		date = string([]byte(date)[:10])
	}
	id1 := json.ID
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	print(content)

	imageUrl1 := ""
	if len(name) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range name {
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

	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistImageByID(id) {

			// data := make(map[string]interface{})
			// if content != "" {
			// 	data["content"] = content
			// }
			// if date != "" {
			// 	data["date"] = date
			// }
			// data["name"] = imageUrl1

			models.EditImage(id, content, date, imageUrl1)

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

func DeleteImage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistImageByID(id) {
			models.DeleteImage(id)
		} else {
			code = e.ERROR_NOT_EXIST_Image
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
