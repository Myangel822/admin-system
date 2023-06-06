package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	member := models.GetPasswordByUsername(username)
	valid.Required(member, "member").Message("用户不存在")
	password1 := member.Password
	if password1 == "111111" {
		a := models.Member{Username: username, Password: password}
		ok, _ := valid.Valid(&a)
		data := make(map[string]interface{})
		code := e.INVALID_PARAMS
		if ok {
			member := models.CheckMember(username, password)
			valid.Required(member, "member").Message("用户不存在")
			if !valid.HasErrors() {
				token, err1 := util.GenerateToken(username, password)
				v1.SetCurrentUser(c, member)
				if err1 != nil {
					code = e.ERROR_MANAGER_TOKEN
				} else {
					data["token"] = token
					data["user"] = member
					code = e.SUCCESS
				}
			} else {
				code = e.ERROR_MANAGER
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password)) //验证（对比）
		if err != nil {
			fmt.Printf("psd wrong")
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("psd right")
			a := models.Member{Username: username, Password: password}
			ok, _ := valid.Valid(&a)
			data := make(map[string]interface{})
			code := e.INVALID_PARAMS
			if ok {
				member := models.CheckMember(username, password1)
				valid.Required(member, "member").Message("用户不存在")
				if !valid.HasErrors() {
					token, err1 := util.GenerateToken(username, password1)
					v1.SetCurrentUser(c, member)
					if err1 != nil {
						code = e.ERROR_MANAGER_TOKEN
					} else {
						data["token"] = token
						data["user"] = member
						code = e.SUCCESS
					}
				} else {
					code = e.ERROR_MANAGER
				}
			} else {
				for _, err := range valid.Errors {
					log.Println(err.Key, err.Message)
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
		}

	}

}
