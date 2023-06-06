package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type News struct {
	Model

	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedOn string `json:"created_on"`
	ImageUrl  string `json:"imageUrl"`
	Uid       int
	img       []string
}

func GetNews(pageNum int, pageSize int, maps interface{}) (news []News) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&news)

	return
}

func GetNewsTotal(maps interface{}) (count int) {
	db.Model(&News{}).Where(maps).Count(&count)

	return
}
func ExistNewsByTitle(id int, title string) bool {
	var news News
	db.Select("id").Where("title = ?", title).First(&news)
	if news.ID > 0 && news.ID != id {
		return true
	}

	return false
}

func AddNews(title string, content string, createOn string, imageUrl string, uid int) bool {
	db.Create(&News{
		Title:     title,
		Content:   content,
		CreatedOn: createOn,
		ImageUrl:  imageUrl,
		Uid:       uid,
	})

	return true
}
func DeleteNews(id int) bool {
	db.Where("id = ?", id).Delete(&News{})

	return true
}
func (news *News) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Format("2006-01-02"))
	return nil
}
func ExistNewsByID(id int) bool {
	var news News
	db.Select("id").Where("id = ?", id).First(&news)
	if news.ID > 0 {
		return true
	}

	return false
}

func EditNews(id int, content string, title string, imageUrl string) bool {
	var news News
	news.Content = content
	news.Title = title
	news.ImageUrl = imageUrl

	db.Model(&News{}).Where("id = ?", id).Updates(news)

	return true
}

func GetNewsByName(title string) (News []News) {

	db.Where("title like ?", title+"%").Find(&News)
	return
}

func GetNewsByUid(uid int) (News []News) {
	db.Where("uid = ?", uid).Find(&News)
	return
}

func GetNewsByNameAndUid(uid int, title string) (News []News) {

	db.Where("uid = ? and title like ?", uid, title+"%").Find(&News)
	return
}
