package database

import "github.com/anuragrao04/maaya-concert-qr/models"

func GetUser(prn string) (user models.User, err error) {
	err = DB.Where("prn = ?", prn).First(&user).Error
	return
}

func SetPresent(user *models.User) (err error) {
	err = DB.Model(&user).Update("is_present", true).Error
	return
}

func CreateUser(user *models.User) (err error) {
	err = DB.Create(user).Error
	return
}
