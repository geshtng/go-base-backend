package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/config"
	"github.com/geshtng/go-base-backend/internal/models"
)

var db *gorm.DB

func Connect() (err error) {
	conf := config.InitConfigDsn()

	db, err = gorm.Open(mysql.Open(conf), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Article{},
		&models.User{},
	)

	if err != nil {
		return err
	}

	return nil
}

func Get() *gorm.DB {
	return db
}
