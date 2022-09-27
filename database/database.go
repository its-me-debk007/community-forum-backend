package database

import (
	"log"
	"os"

	"github.com/its-me-debk007/community-forum-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	if err := db.AutoMigrate(
		new(models.Author),
		new(models.Post),
	); err != nil {
		log.Fatalln("AUTO_MIGRATION_ERROR")
	}

}
