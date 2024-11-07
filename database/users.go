package database

import (
	"github.com/anuragrao04/maaya-concert-qr/backups"
	"github.com/anuragrao04/maaya-concert-qr/models"
)

func GetUserByID(id uint) (user models.User, err error) {
	err = DB.First(&user, id).Error
	return
}
func GetUser(prn string) (user models.User, err error) {
	err = DB.Where("prn = ?", prn).First(&user).Error
	return
}

func GetUserByEmail(email string) (user models.User, err error) {
	err = DB.Where("email = ?", email).First(&user).Error
	return
}

func SetPresent(user *models.User) (err error) {
	err = DB.Model(&user).Update("is_present", true).Error
	go backups.IncrementWriteCount()
	return
}

func CreateUser(user *models.User) (err error) {
	err = DB.Create(user).Error
	go backups.IncrementWriteCount()

	return
}

func GetAllUsers() (users []models.User, err error) {
	err = DB.Find(&users).Error
	return
}
