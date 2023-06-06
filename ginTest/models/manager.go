package models

type Manager struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
}

func CheckManager(username, password string) bool {
	var manager Manager
	db.Select("id").Where(Manager{Username: username, Password: password}).First(&manager)
	if manager.ID > 0 {
		return true
	}

	return false
}

func AddManager(username, password, identity string) bool {
	db.Create(&Manager{
		Username: username,
		Password: password,
		Identity: identity,
	})

	return true
}

func ExistManagerByUsername(id int, username string) bool {
	var manager Manager
	db.Select("username").Where("username = ?", username).First(&manager)

	if manager.ID > 0 && manager.ID != id {
		return true
	} else {
		return false
	}
}
