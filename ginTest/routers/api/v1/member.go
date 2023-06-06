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
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/enforcer"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type Member struct {
	Identity       string   `json:"identity"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	Achievement    string   `json:"achievement"`
	Mail           string   `json:"mail"`
	Research       string   `json:"research"`
	Introduction   string   `json:"introduction"`
	State          string   `json:"state"`
	Graduation     string   `json:"graduation"`
	Graduationtime string   `json:"graduationtime"`
	Image          []string `json:"image`
}

type Member1 struct {
	ID             string   `json:"id"`
	Identity       string   `json:"identity"`
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	Achievement    string   `json:"achievement"`
	Mail           string   `json:"mail"`
	Research       string   `json:"research"`
	Introduction   string   `json:"introduction"`
	State          string   `json:"state"`
	Graduation     string   `json:"graduation"`
	Graduationtime string   `json:"graduationtime"`
	Image          []string `json:"image`
}

func GetMemberByName(c *gin.Context) {
	name := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetMemberByNamemohu(name)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetMemberByNamemohuAndId(member.ID, name)

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

func GetMember(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistMemberByID(id) {
			data = models.GetMember(id)
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

func CountMembers(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["total"] = models.GetMemberTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetMembers(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		data["lists"] = models.GetMembers(util.GetPage(c), models.GetMemberTotal(maps), maps)
		data["total"] = models.GetMemberTotal(maps)
	} else {
		data["lists"] = models.GetMember(member.ID)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetMembersForword(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetMembers(util.GetPage(c), models.GetMemberTotal(maps), maps)
	data["total"] = models.GetMemberTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddMember(c *gin.Context) {
	json := Member{}
	c.BindJSON(&json)
	username := json.Username
	password := json.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePWD)

	name := json.Name
	identity := json.Identity
	phone := json.Phone
	mail := json.Mail
	graduation := json.Graduation
	graduationtime := json.Graduationtime
	graduationtime = string([]byte(graduationtime)[:10])
	achievement := json.Achievement
	research := json.Research
	introduction := json.Introduction
	state := json.State

	image := json.Image
	imageUrl1 := ""
	if len(image) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range image {
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
	valid.Required(name, "name").Message("姓名不能为空")
	valid.MaxSize(name, 100, "name").Message("姓名最长为100字符")
	valid.Required(identity, "identity").Message("身份不能为空")
	valid.MaxSize(identity, 100, "identity").Message("身份最长为100字符")
	valid.Required(phone, "phone").Message("电话年份不能为空")
	valid.MaxSize(phone, 100, "phone").Message("电话年份最长为10字符")
	valid.Required(research, "research").Message("研究方向不能为空")
	valid.MaxSize(research, 100, "research").Message("研究方向最长为100字符")
	valid.Required(mail, "mail").Message("邮箱不能为空")
	valid.MaxSize(mail, 100, "mail").Message("邮箱最长为100字符")
	valid.Required(achievement, "achievement").Message("成果不能为空")
	valid.MaxSize(achievement, 65535, "achievement").Message("成果最长为65535字符")
	valid.Required(introduction, "introduction").Message("个人介绍不能为空")
	valid.MaxSize(introduction, 65535, "introduction").Message("个人介绍最长为65535字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistMemberByUsername(id, username) {
			code = e.SUCCESS

			if identity == "系统管理员" {
				enforcer.EnforcerTool().AddGroupingPolicy(name, "admin")
				hasPolicy := enforcer.EnforcerTool().HasGroupingPolicy(name, "admin")
				fmt.Println(hasPolicy)
				enforcer.EnforcerTool().LoadPolicy()
			} else if identity == "毕业学生" {
				enforcer.EnforcerTool().AddGroupingPolicy(name, "graduate")
				enforcer.EnforcerTool().LoadPolicy()
			} else {
				enforcer.EnforcerTool().AddGroupingPolicy(name, "user")
				enforcer.EnforcerTool().LoadPolicy()
			}
			models.AddMember(username, encodePWD, name, phone, achievement, mail, research, identity, introduction, state, graduation, graduationtime, imageUrl1)

		} else {
			code = e.ERROR_EXIST_Member
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditMember(c *gin.Context) {

	json := Member1{}
	c.BindJSON(&json)

	id1 := json.ID
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	name := json.Name
	identity := json.Identity
	phone := json.Phone
	mail := json.Mail
	graduation := json.Graduation
	graduationtime := json.Graduationtime
	graduationtime = string([]byte(graduationtime)[:10])
	achievement := json.Achievement
	research := json.Research
	introduction := json.Introduction
	// state := json.State

	image := json.Image
	imageUrl1 := ""
	if len(image) != 0 {
		var (
			path  string
			path1 string
			enc   = base64.StdEncoding
		)
		for index, img := range image {
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

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistMemberByID(id) {

			// data := make(map[string]interface{})
			// data["name"] = name
			// data["phone"] = phone
			// data["mail"] = mail
			// data["identity"] = identity
			// data["achievement"] = achievement
			// data["introduction"] = introduction
			// data["research"] = research
			// data["state"] = state
			// data["graduation"] = graduation
			// data["graduationtime"] = graduationtime
			// data["imageUrl"] = imageUrl1
			if identity == "毕业学生" {
				enforcer.EnforcerTool().DeleteUser(name)
				enforcer.EnforcerTool().AddGroupingPolicy(name, "graduate")
				hasPolicy := enforcer.EnforcerTool().HasGroupingPolicy(name, "graduate")
				enforcer.EnforcerTool().LoadPolicy()
				fmt.Println(hasPolicy)
			}

			models.EditMember(id, name, phone, mail, identity, achievement, introduction, research, graduation, graduationtime, imageUrl1)

		} else {
			code = e.ERROR_NOT_EXIST_Member
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditMemberByName(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	username := c.Query("username")
	password := c.Query("password")

	name := c.Query("name")
	identity := c.Query("identity")
	phone := c.Query("phone")
	mail := c.Query("mail")
	achievement := c.Query("achievement")
	research := c.Query("research")
	introduction := c.Query("introduction")
	state := c.Query("state")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("姓名不能为空")
	valid.MaxSize(name, 100, "name").Message("姓名最长为100字符")
	valid.Required(research, "research").Message("研究方向不能为空")
	valid.MaxSize(research, 100, "research").Message("研究方向最长为100字符")
	valid.Required(identity, "identity").Message("身份不能为空")
	valid.MaxSize(identity, 100, "identity").Message("身份最长为100字符")
	valid.Required(phone, "phone").Message("电话年份不能为空")
	valid.MaxSize(phone, 100, "phone").Message("电话年份最长为10字符")
	valid.Required(mail, "mail").Message("邮箱不能为空")
	valid.MaxSize(mail, 100, "mail").Message("邮箱最长为100字符")
	valid.Required(achievement, "achievement").Message("成果不能为空")
	valid.MaxSize(achievement, 65535, "achievement").Message("成果最长为65535字符")
	valid.Required(introduction, "introduction").Message("个人介绍不能为空")
	valid.MaxSize(introduction, 65535, "introduction").Message("个人介绍最长为65535字符")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistMemberByID(id) {
			if !models.ExistMemberByName(id, name) {
				data := make(map[string]interface{})
				data["username"] = username
				data["password"] = password
				data["name"] = name
				data["phone"] = phone
				data["mail"] = mail
				data["identity"] = identity
				data["achievement"] = achievement
				data["introduction"] = introduction
				data["research"] = research
				data["state"] = state
				models.EditMemberByName(name, data)
			} else {
				code = e.ERROR_EXIST_Member
			}
		} else {
			code = e.ERROR_NOT_EXIST_Member
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func ResetPassword(c *gin.Context) {
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100
	}
	password := "111111"
	valid := validation.Validation{}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistMemberByID(id) {

			models.ResetPassword(id, password)
		} else {
			code = e.ERROR_NOT_EXIST_Member
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteMember(c *gin.Context) {
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
		if models.ExistMemberByID(id) {
			models.DeleteMember(id)
		} else {
			code = e.ERROR_NOT_EXIST_Member
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func GetCurrentUser(c *gin.Context) (userInfo models.Member) {
	session := sessions.Default(c)
	userInfo = session.Get("currentUser").(models.Member) // 类型转换一下
	return
}

func SetCurrentUser(c *gin.Context, userInfo models.Member) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}
