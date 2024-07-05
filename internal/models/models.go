package models

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"unique;type:varchar(100)"`
}
