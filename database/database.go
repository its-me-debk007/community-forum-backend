package database

import (
	"fmt"
	"log"
	"os"
	"time"

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
		new(models.User),
		new(models.Post),
	); err != nil {
		log.Fatalln("AUTO_MIGRATION_ERROR")
	}

	// populateDB()
}

func populateDB() {
	for i := 0; i < 100; i++ {
		user := models.User{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:       fmt.Sprintf("User %d", i),
			ProfilePic: "https://cdn.tutsplus.com/gamedev/uploads/legacy/043_freeShmupSprites/Free_Shmup_Sprites_Boss_Battle.jpg",
		}
		DB.Create(&user)
	}
}
