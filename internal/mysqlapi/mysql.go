package mysqlapi

import (
	"log"

	"github.com/just-umyt/go-api/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
