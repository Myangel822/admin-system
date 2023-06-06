package models

//"time"

//"github.com/jinzhu/gorm"

type Article struct {
	Model

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

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}
func ExistArticleByTitle(id int, title string) bool {
	var article Article
	db.Select("id").Where("title = ?", title).First(&article)

	if article.ID > 0 && article.ID != id {
		return true
	}

	return false
}
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)

	return
}

func EditArticle(id int, title string, journal string, author string, authors string, date string, link string, papercode string, abstract string, theyear string, content string) bool {
	var article Article
	article.Title = title
	article.Author = author
	article.Authors = authors
	article.Journal = journal
	article.Link = link
	article.Date = date
	article.Papercode = papercode
	article.Abstract = abstract
	article.Theyear = theyear
	article.Content = content
	db.Model(&Article{}).Where("id = ?", id).Updates(article)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		Title:     data["title"].(string),
		Journal:   data["journal"].(string),
		Author:    data["author"].(string),
		Date:      data["date"].(string),
		Authors:   data["authors"].(string),
		Papercode: data["papercode"].(string),
		Link:      data["link"].(string),
		Abstract:  data["abstract"].(string),
		Theyear:   data["theyear"].(string),
		State:     data["state"].(string),
		Hide:      data["hide"].(string),
		Uid:       data["uid"].(int),
		Content:   data["content"].(string),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func GetArticleByName(title string) (Article []Article) {

	db.Where("title like ?", title+"%").Find(&Article)
	return
}

func EditArticleState(id int, state string) bool {
	db.Model(&Article{}).Where("id = ?", id).Update("state", state)
	return true
}

func EditArticleHide(id int, hide string) bool {
	db.Model(&Article{}).Where("id = ?", id).Update("hide", hide)

	return true
}

func GetArticleByUid(uid int) (Article []Article) {
	db.Where("uid = ?", uid).Find(&Article)
	return
}

func GetArticleByNameAndUid(uid int, title string) (Article []Article) {

	db.Where("uid = ? and title like ?", uid, title+"%").Find(&Article)
	return
}

func GetArticleForword(state string, hide string) (Article []Article) {
	db.Where("state = ? and hide = ?", state, hide).Find(&Article)
	return
}
