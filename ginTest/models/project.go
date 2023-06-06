package models

//"time"

//"github.com/jinzhu/gorm"
type Project struct {
	Model

	Title   string `json:"title"`
	Content string `json:"content"`
	State   string `json:"state"`
	Date    string `json:"date"`
	Uid     int
}

func GetProjects(pageNum int, pageSize int, maps interface{}) (project []Project) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&project)

	return
}
func GetProject(id int) (project Project) {
	db.Where("id = ?", id).First(&project)

	return
}
func GetProjectTotal(maps interface{}) (count int) {
	db.Model(&Project{}).Where(maps).Count(&count)

	return
}
func ExistProjectByName(id int, name string) bool {
	var project Project
	db.Select("id").Where("title = ?", name).First(&project)
	if project.ID > 0 && project.ID != id {
		return true
	}

	return false
}

func AddProject(title string, content string, state string, date string, uid int) bool {
	db.Create(&Project{
		Title:   title,
		Content: content,
		State:   state,
		Date:    date,
		Uid:     uid,
	})

	return true
}
func DeleteProject(id int) bool {
	db.Where("id = ?", id).Delete(&Project{})

	return true
}

func ExistProjectByID(id int) bool {
	var project Project
	db.Select("id").Where("id = ?", id).First(&project)
	if project.ID > 0 {
		return true
	}

	return false
}

func EditProject(id int, content string, title string, state string, theyear string) bool {
	var project Project
	project.Content = content
	project.Date = theyear
	project.State = state
	project.Title = title

	db.Model(&Project{}).Where("id = ?", id).Updates(project)

	return true
}

func GetProjectByName(name string) (Project []Project) {

	db.Where("title like ?", name+"%").Find(&Project)
	return
}

func EditProjectState(id int, state string) bool {
	db.Model(&Project{}).Where("id = ?", id).Update("state", state)

	return true
}
func GetProjectByUid(uid int) (Project []Project) {
	db.Where("uid = ?", uid).Find(&Project)
	return
}

func GetProjectByNameAndUid(uid int, name string) (Project []Project) {

	db.Where("uid = ? and title like ?", uid, name+"%").Find(&Project)
	return
}
