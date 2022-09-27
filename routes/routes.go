package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/its-me-debk007/community-forum-backend/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/forum/post", controllers.GetAllPosts)
}
