package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/its-me-debk007/community-forum-backend/database"
	"github.com/its-me-debk007/community-forum-backend/helpers"
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
	postImages := []string{}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Message{
				Message: fmt.Sprintf("Error in opening file:- %s", err.Error()),
			})
		}

		fileUrl, err := helpers.UploadImage(file, time.Now())
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Message{
				Message: err.Error(),
			})
		}

		postImages = append(postImages, fileUrl)
	}

	log.Println(postImages)

	data := models.Post{
		PostTitle:       postTitle,
		PostDescription: postDescription,
		PostImages:      postImages,
		AuthorID:        44,
	}

	if err := database.DB.Create(&data); err.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: err.Error.Error(),
		})
	}

	return c.JSON(models.Message{
		Message: "post created successfully",
	})
}

func LikePost(c *fiber.Ctx) error {
	body := new(struct {
		PostId uint `json:"post_id"`
	})

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: err.Error(),
		})
	}

	post := models.Post{}

	database.DB.First(&post, "post_id = ?", body.PostId)
	if post.PostID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: "post doesen't exist",
		})
	}

	database.DB.Model(&post).Update("likes_count", post.LikesCount+1)

	return c.JSON(models.Message{
		Message: "request successful",
	})
}