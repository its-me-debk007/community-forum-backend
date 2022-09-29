package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/its-me-debk007/community-forum-backend/database"
	"github.com/its-me-debk007/community-forum-backend/models"
)

func GetAllPosts(c *fiber.Ctx) error {
	data := []models.Post{}
	database.DB.Find(&data)

	author := models.User{}

	for i := range data {
		database.DB.First(&author, "id = ?", data[i].AuthorID)
		data[i].Author = author
		author = models.User{}
	}

	return c.JSON(data)
}

func CreatePost(c *fiber.Ctx) error {
	// postTitle := c.FormValue("post_title")
	// postDescription := c.FormValue("post_description")
	form, err := c.MultipartForm()
	postTitle := form.Value["post_title"][0]
	postDescription := form.Value["post_description"][0]

	if err != nil {
		return c.JSON(models.Message{
			Message: fmt.Sprintf("MULTIPART FORM ERROR:- %s", err.Error()),
		})
	}

	files := form.File["post_images"]
	// log.Println("\n" + files[0].Filename + "\n")

	data := models.Post{
		PostTitle:       postTitle,
		PostDescription: postDescription,
		PostImages:      files[0].Filename,
		AuthorID:        54,
	}

	if err := database.DB.Create(&data); err.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: fmt.Sprintf("DATABASE DATA CREATE ERROR :-  %s", err.Error.Error()),
		})
	}

	return c.JSON(models.Message{
		Message: "post created successfully",
	})
}
