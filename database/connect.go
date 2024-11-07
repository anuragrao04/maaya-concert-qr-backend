package database

import (
	"sync"

	"github.com/anuragrao04/maaya-concert-qr/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // other packages can access this variable
var DBLock sync.Mutex

func Connect(dbFilePath string) error {
	var err error

	DBLock.Lock()
	defer DBLock.Unlock()
	DB, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&models.User{})
	return err
}
