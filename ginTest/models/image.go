package models

//"time"

//"github.com/jinzhu/gorm"
type Image struct {
	Model
	Name    string `json:"name"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Uid     int
}

func GetImages(pageNum int, pageSize int, maps interface{}) (image []Image) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&image)

	return
}
func GetImage(id int) (image Image) {
	db.Where("id = ?", id).First(&image)

	return
}
func GetImageTotal(maps interface{}) (count int) {
	db.Model(&Image{}).Where(maps).Count(&count)

	return
}
func ExistImageByName(id int, name string) bool {
	var image Image
	db.Select("id").Where("name = ?", name).First(&image)
	if image.ID > 0 && image.ID != id {
		return true
	}

	return false
}

func AddImage(name string, date string, content string, uid int) bool {
	db.Create(&Image{
		Name:    name,
		Date:    date,
		Content: content,
		Uid:     uid,
	})

	return true
}
func DeleteImage(id int) bool {
	db.Where("id = ?", id).Delete(&Image{})

	return true
}

func ExistImageByID(id int) bool {
	var image Image
	db.Select("id").Where("id = ?", id).First(&image)
	if image.ID > 0 {
		return true
	}

	return false
}

func EditImage(id int, content string, date string, imageUrl string) bool {
	var image Image
	image.Content = content
	image.Date = date
	image.Name = imageUrl

	db.Model(&Image{}).Where("id = ?", id).Updates(image)

	return true
}

func GetImageByDate(date string) (Image []Image) {

	db.Where("date like ?", date+"%").Find(&Image)
	return
}

func GetImageByUid(uid int) (Image []Image) {
	db.Where("uid = ?", uid).Find(&Image)
	return
}

func GetImageByDateAndUid(uid int, date string) (Image []Image) {

	db.Where("uid = ? and date like ?", uid, date+"%").Find(&Image)
	return
}
