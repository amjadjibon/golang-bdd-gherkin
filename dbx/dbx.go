package dbx

import (
	"os"

	"github.com/amjadjibon/golang-bdd-gherkin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDB() {
	url := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})

	DB = db
}
