package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"

	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type Article struct {
	Title     string `json:"title"`
	Journal   string `json:"journal"`
	Author    string `json:"author"`
	Authors   string `json:"authors"`
	Date      string `json:"date"`
	Link      string `json:"link"`
	Papercode string `json:"papercode"`
	Abstract  string `json:"abstract"`
	Theyear   string `json:"theyear"`
	State     string `json:"state"`
	Hide      string `json:"hide"`
	Uid       int
	Content   string `json:"content"`
}

func GetArticleForword(c *gin.Context) {
	state := "已发表"
	hide := "显示"
	data := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data["lists"] = models.GetArticleForword(state, hide)

		code = e.SUCCESS
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
func CountArticle(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetArticleByName(c *gin.Context) {
	title := c.Query("name")

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(title, "name").Message("name不能为空")
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {
		if !valid.HasErrors() {
			data["lists"] = models.GetArticleByName(title)

			code = e.SUCCESS
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			data["lists"] = models.GetArticleByNameAndUid(member.ID, title)

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

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
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

// 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	code := e.INVALID_PARAMS
	member := GetCurrentUser(c)
	if member.Identity == "系统管理员" {

		if !valid.HasErrors() {
			code = e.SUCCESS

			data["lists"] = models.GetArticles(util.GetPage(c), models.GetArticleTotal(maps), maps)
			data["total"] = models.GetArticleTotal(maps)

		} else {
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	} else {
		if !valid.HasErrors() {
			code = e.SUCCESS

			data["lists"] = models.GetArticleByUid(member.ID)

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

// 新增文章
func AddArticle(c *gin.Context) {
	title := c.Query("title")
	journal := c.Query("journal")
	author := c.Query("author")
	authors := c.Query("authors")
	date := c.Query("date")
	date = string([]byte(date)[:10])
	link := c.Query("link")
	papercode := c.Query("papercode")
	abstract := c.Query("abstract")
	theyear := c.Query("theyear")
	state := c.Query("state")
	hide := c.Query("hide")
	content := author + "." + authors + "." + title + "." + journal + "." + date
	id := 0
	uid := GetCurrentUser(c).ID
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(journal, "journal").Message("期刊不能为空")
	valid.MaxSize(journal, 100, "journal").Message("期刊最长为100字符")
	valid.Required(author, "author").Message("第一作者不能为空")
	valid.MaxSize(author, 100, "author").Message("第一作者最长为100字符")
	valid.Required(authors, "authors").Message("其他作者不能为空")
	valid.MaxSize(authors, 100, "authors").Message("其他作者最长为100字符")
	valid.Required(date, "date").Message("日期不能为空")
	valid.MaxSize(date, 100, "date").Message("日期最长为100字符")
	valid.Required(link, "link").Message("详情页链接不能为空")
	valid.MaxSize(link, 100, "link").Message("详情页链接最长为100字符")
	valid.Required(papercode, "papercode").Message("代码链接不能为空")
	valid.MaxSize(papercode, 100, "papercode").Message("代码链接最长为100字符")
	valid.Required(abstract, "abstract").Message("摘要不能为空")
	valid.MaxSize(abstract, 65535, "abstract").Message("摘要最长为65535字符")
	valid.Required(theyear, "theyear").Message("论文年份不能为空")
	valid.MaxSize(theyear, 10, "theyear").Message("论文年份最长为10字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistArticleByTitle(id, title) {
			data := make(map[string]interface{})
			data["title"] = title
			data["journal"] = journal
			data["author"] = author
			data["authors"] = authors
			data["link"] = link
			data["papercode"] = papercode
			data["abstract"] = abstract
			data["date"] = date
			data["theyear"] = theyear
			data["state"] = state
			data["hide"] = hide
			data["uid"] = uid
			data["content"] = content
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {
	var article Article
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)
	title := c.Query("title")
	article.Title = title
	journal := c.Query("journal")
	article.Journal = journal
	author := c.Query("author")
	article.Author = author
	authors := c.Query("authors")
	article.Authors = authors
	date := c.Query("date")
	if date != "" {
		date = string([]byte(date)[:10])
	}

	article.Date = date
	link := c.Query("link")
	article.Link = link
	papercode := c.Query("papercode")
	article.Papercode = papercode
	abstract := c.Query("abstract")
	article.Abstract = abstract
	theyear := c.Query("theyear")
	article.Theyear = theyear
	content := author + "." + authors + "." + title + "." + journal + "." + date
	article.Content = content
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {

			models.EditArticle(id, title, journal, author, authors, date, link, papercode, abstract, theyear, content)
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
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
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
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
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
		"data": make(map[string]string),
	})
}

func EditArticleState(c *gin.Context) {
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
		if models.ExistArticleByID(id) {

			models.EditArticleState(id, state)

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

func EditArticleHide(c *gin.Context) {
	id1 := c.Query("id")
	id, err := strconv.Atoi(id1)
	if err != nil {

		fmt.Println("can't convert to int")

	} else {

		fmt.Printf("type:%T value:%#v\n", id, id) //type:int value:100

	}
	println(id1)

	hide := c.Query("hide")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {

			models.EditArticleHide(id, hide)

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
