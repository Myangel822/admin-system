package models

//"time"

// "github.com/jinzhu/gorm"
type Member struct {
	Model

	Identity       string `json:"identity"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Achievement    string `json:"achievement"`
	Mail           string `json:"mail"`
	Research       string `json:"research"`
	Introduction   string `json:"introduction"`
	State          string `json:"state"`
	Graduation     string
	Graduationtime string
	ImageUrl       string
}

func GetMembers(pageNum int, pageSize int, maps interface{}) (member []Member) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&member)

	return
}
func GetMember(id int) (member []Member) {
	db.Where("id = ?", id).First(&member)

	return
}

func GetMemberByName(name string) (member []Member) {
	db.Where("name = ?", name).Find(&member)
	return
}

func GetMemberTotal(maps interface{}) (count int) {
	db.Model(&Member{}).Where(maps).Count(&count)
	return
}
func ExistMemberByName(id int, name string) bool {
	var member Member
	db.Select("id").Where("name = ?", name).First(&member)
	if member.ID > 0 && member.ID != id {
		return true
	}

	return false
}

func ExistMemberByUsername(id int, username string) bool {
	var member Member
	db.Select("id").Where("name = ?", username).First(&member)
	if member.ID > 0 && member.ID != id {
		return true
	}

	return false
}

func AddMember(username string, password string, name string, phone string, achievement string, mail string, research string, identity string, introduction string, state string, graduation string, graduationtime string, imageUrl string) bool {
	db.Create(&Member{
		Username:       username,
		Password:       password,
		Name:           name,
		Identity:       identity,
		Phone:          phone,
		Mail:           mail,
		Achievement:    achievement,
		Research:       research,
		Introduction:   introduction,
		State:          state,
		Graduation:     graduation,
		Graduationtime: graduationtime,
		ImageUrl:       imageUrl,
	})

	return true
}
func DeleteMember(id int) bool {
	db.Where("id = ?", id).Delete(&Member{})

	return true
}

func ExistMemberByID(id int) bool {
	var member Member
	db.Select("id").Where("id = ?", id).First(&member)
	if member.ID > 0 {
		return true
	}

	return false
}

func EditMember(id int, name string, phone string, mail string, identity string, achievement string, introduction string, research string, graduation string, graduationtime string, imageUrl1 string) bool {
	var member Member
	member.Achievement = achievement
	member.Graduation = graduation
	member.Graduationtime = graduationtime
	member.Identity = identity
	member.ImageUrl = imageUrl1
	member.Introduction = introduction
	member.Mail = mail
	member.Name = name
	member.Phone = phone
	member.Research = research
	db.Model(&Member{}).Where("id = ?", id).Updates(member)

	return true
}

func EditMemberByName(name string, data interface{}) bool {
	db.Model(&Member{}).Where("name = ?", name).Updates(data)

	return true
}

func ResetPassword(id int, password string) bool {
	db.Model(&Member{}).Where("id = ?", id).Update("password", password)
	return true
}

func GetMemberByNamemohu(name string) (Member []Member) {

	db.Where("name like ?", name+"%").Find(&Member)
	return
}

func CheckMember(username, password string) (Member Member) {
	db.Model(&Member).Where("username = ? and password = ?", username, password).First(&Member)

	return Member
}

func GetPasswordByUsername(username string) (member Member) {
	db.Model(&member).Where("username = ?", username).First(&member)
	return member
}

func GetMemberByNamemohuAndId(id int, name string) (Member []Member) {

	db.Where("id = ? and name like ?", id, name+"%").Find(&Member)
	return
}
