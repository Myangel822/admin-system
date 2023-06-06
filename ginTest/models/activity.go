package models

//"time"

//"github.com/jinzhu/gorm"

type Activity struct {
	Model

	Title    string `json:"title"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Kind     string `json:"kind"`
	ImageUrl string `json:"imageUrl"`
	Uid      int
}

func GetActivitys(pageNum int, pageSize int, maps interface{}) (activity []Activity) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&activity)

	return
}
func GetActivity(id int) (activity Activity) {
	db.Where("id = ?", id).First(&activity)
	return
}
func GetActivityTotal(maps interface{}) (count int) {
	db.Model(&Activity{}).Where(maps).Count(&count)
	return
}
func ExistActivityByTitle(id int, title string) bool {
	var activity Activity
	db.Select("id").Where("title = ?", title).First(&activity)
	if activity.ID > 0 && activity.ID != id {
		return true
	}

	return false
}

func AddActivity(title string, date string, content string, kind string, imageUrl string, uid int) bool {
	db.Create(&Activity{
		Title:    title,
		Content:  content,
		Date:     date,
		Kind:     kind,
		Uid:      uid,
		ImageUrl: imageUrl,
	})

	return true
}
func DeleteActivity(id int) bool {
	db.Where("id = ?", id).Delete(&Activity{})

	return true
}

func ExistActivityByID(id int) bool {
	var activity Activity
	db.Select("id").Where("id = ?", id).First(&activity)
	if activity.ID > 0 {
		return true
	}

	return false
}

func EditActivity(id int, content string, title string, kind string, date string, imageUrl string) bool {
	var activity Activity
	activity.Content = content
	activity.Date = date
	activity.ImageUrl = imageUrl
	activity.Kind = kind
	activity.Title = title
	db.Model(&Activity{}).Where("id = ?", id).Updates(activity)

	return true
}

func GetActivityByName(title string) (Activity []Activity) {

	db.Where("title like ?", title+"%").Find(&Activity)
	return
}

func GetActivityByKind(kind string) (Activity []Activity) {
	db.Where("kind = ?", kind).Find((&Activity))
	return
}

func GetActivityByUid(uid int) (Activity []Activity) {
	db.Where("uid = ?", uid).Find(&Activity)
	return
}

func GetActivityByNameAndUid(uid int, title string) (Activity []Activity) {

	db.Where("uid = ? and title like ?", uid, title+"%").Find(&Activity)
	return
}

func GetActivityByKindAndUid(uid int, kind string) (Activity []Activity) {
	db.Where("uid = ? and kind = ?", uid, kind).Find((&Activity))
	return
}
