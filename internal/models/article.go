package models

type Article struct {
	ID          int64 `gorm:"primary_key"`
	Title       string
	Description string
}
