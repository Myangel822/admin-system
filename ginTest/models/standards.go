package models

//"time"

//"github.com/jinzhu/gorm"

type Standards struct {
	Model

	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Kind    string `json:"kind"`
	State   string `json:"state"`
	Uid     int
}

func GetStandardss(pageNum int, pageSize int, maps interface{}) (Standards []Standards) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Standards)

	return
}
func GetStandards(id int) (Standards Standards) {
	db.Where("id = ?", id).First(&Standards)
	return
}
func GetStandardsTotal(maps interface{}) (count int) {
	db.Model(&Standards{}).Where(maps).Count(&count)
	return
}

func ExistStandardsByTitle(id int, title string) bool {
	var Standards Standards
	db.Select("id").Where("title = ?", title).First(&Standards)
	if Standards.ID > 0 && Standards.ID != id {
		return true
	}

	return false
}

func AddStandards(data map[string]interface{}) bool {
	db.Create(&Standards{
		Title:   data["title"].(string),
		Content: data["content"].(string),
		Date:    data["date"].(string),
		Kind:    data["kind"].(string),
		State:   data["state"].(string),
		Uid:     data["uid"].(int),
	})

	return true
}
func DeleteStandards(id int) bool {
	db.Where("id = ?", id).Delete(&Standards{})

	return true
}

func ExistStandardsByID(id int) bool {
	var standards Standards
	db.Select("id").Where("id = ?", id).First(&standards)
	if standards.ID > 0 {
		return true
	}

	return false
}

func GetStandardsByName(name string) (Standards []Standards) {

	db.Where("title like ?", name+"%").Find(&Standards)
	return
}

func CountStandardsByName(name string) (count int) {

	db.Where("title like ?", name+"%").Count(&count)
	return
}

func EditStandards(id int, content string, title string, kind string, date string, state string) bool {
	var standards Standards
	standards.Content = content
	standards.Title = title
	standards.Date = date
	standards.Kind = kind
	standards.State = state
	db.Model(&Standards{}).Where("id = ?", id).Updates(standards)

	return true
}

func EditStandardsState(id int, state string) bool {
	db.Model(&Standards{}).Where("id = ?", id).Update("state", state)

	return true
}
func GetStandardsByUid(uid int) (Standards []Standards) {
	db.Where("uid = ?", uid).Find(&Standards)
	return
}

func CountStandardsByNameAndUid(uid int, name string) (count int) {

	db.Where("uid = ? and title like ?", uid, name+"%").Count(&count)
	return
}
