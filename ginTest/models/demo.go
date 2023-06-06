package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Demo struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Video   string
	Uid     int
}

func GetDemos(pageNum int, pageSize int, maps interface{}) (Demo []Demo) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Demo)

	return
}

func GetDemosTotal(maps interface{}) (count int) {
	db.Model(&Demo{}).Where(maps).Count(&count)

	return
}
func ExistDemoByTitle(id int, title string) bool {
	var Demo Demo
	db.Select("id").Where("title = ?", title).First(&Demo)
	if Demo.ID > 0 && Demo.ID != id {
		return true
	}

	return false
}

func AddDemo(title string, content string, date string, video string, uid int) bool {
	db.Create(&Demo{
		Title:   title,
		Content: content,
		Date:    date,
		Video:   video,
		Uid:     uid,
	})

	return true
}
func DeleteDemo(id int) bool {
	db.Where("id = ?", id).Delete(&Demo{})

	return true
}
func (Demo *Demo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Format("2006-01-02"))
	return nil
}
func ExistDemoByID(id int) bool {
	var Demo Demo
	db.Select("id").Where("id = ?", id).First(&Demo)
	if Demo.ID > 0 {
		return true
	}

	return false
}

func EditDemo(id int, content string, title string, date string, video string) bool {
	var demo Demo
	demo.Content = content
	demo.Date = date
	demo.Title = title
	demo.Video = video
	db.Model(&Demo{}).Where("id = ?", id).Updates(demo)

	return true
}

func GetDemoByName(name string) (Demo []Demo) {

	db.Where("title like ?", name+"%").Find(&Demo)
	return
}

func GetDemoByUid(uid int) (Demo []Demo) {
	db.Where("uid = ?", uid).Find(&Demo)
	return
}

func GetDemoByNameAndUid(uid int, name string) (Demo []Demo) {

	db.Where("uid = ? and title like ?", uid, name+"%").Find(&Demo)
	return
}
