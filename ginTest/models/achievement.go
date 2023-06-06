package models

//"time"

//"github.com/jinzhu/gorm"
type Achievement struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
	Kind    string `json:"kind"`
	Date    string `json:"date"`
	Uid     int
}

func GetAchievements(pageNum int, pageSize int, maps interface{}) (achievement []Achievement) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&achievement)

	return
}
func GetAchievement(id int) (achievement Achievement) {
	db.Where("id = ?", id).First(&achievement)

	return
}
func GetAchievementTotal(maps interface{}) (count int) {
	db.Model(&Achievement{}).Where(maps).Count(&count)

	return
}
func ExistAchievementByName(id int, name string) bool {
	var achievement Achievement
	db.Select("id").Where("title = ?", name).First(&achievement)
	if achievement.ID > 0 && achievement.ID != id {
		return true
	}

	return false
}

func AddAchievement(title string, content string, kind string, date string, uid int) bool {
	db.Create(&Achievement{
		Title:   title,
		Content: content,
		Kind:    kind,
		Date:    date,
		Uid:     uid,
	})

	return true
}
func DeleteAchievement(id int) bool {
	db.Where("id = ?", id).Delete(&Achievement{})

	return true
}

func ExistAchievementByID(id int) bool {
	var achievement Achievement
	db.Select("id").Where("id = ?", id).First(&achievement)
	if achievement.ID > 0 {
		return true
	}

	return false
}

func EditAchievement(id int, content string, title string, kind string, date string) bool {
	var achievement Achievement
	achievement.Content = content
	achievement.Date = date
	achievement.Kind = kind
	achievement.Title = title
	db.Model(&Achievement{}).Where("id = ?", id).Updates(achievement)

	return true
}

func GetAchievementByName(name string) (Achievement []Achievement) {

	db.Where("title like ?", name+"%").Find(&Achievement)
	return
}

func GetAchievementByUid(uid int) (Achievement []Achievement) {
	db.Where("uid = ?", uid).Find(&Achievement)
	return
}

func GetAchievementByNameAndUid(uid int, name string) (Achievement []Achievement) {

	db.Where("uid = ? and title like ?", uid, name+"%").Find(&Achievement)
	return
}
