package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/its-me-debk007/community-forum-backend/database"
	"github.com/its-me-debk007/community-forum-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("ENV LOADING ERROR", err.Error())
	}

	database.ConnectDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("App listen error:-\n" + err.Error())
	}
}
