package models

type User struct {
	ID       int64  `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string
}
