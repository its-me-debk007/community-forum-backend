package controllers

import (
	"fmt"
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
	likes := models.Likes{}

	for i := range data {
		database.DB.First(&author, "id = ?", data[i].AuthorID)
		data[i].Author = author

		database.DB.First(&likes, "id = ?", data[i].PostID)

		var i int
		for i = range likes.Users {
			if likes.Users[i] == 44 {
				break
			}
		}

		if i != len(likes.Users) {
			data[i].IsLiked = true
		}

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
		PostID uint `json:"post_id"`
	})

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: err.Error(),
		})
	}

	post := models.Post{}

	database.DB.First(&post, "post_id = ?", body.PostID)
	if post.PostID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: "post doesen't exist",
		})
	}

	likes := models.Likes{}
	database.DB.First(&likes, "id = ?", body.PostID)

	var i int
	for i = range likes.Users {
		if likes.Users[i] == 32 {
			break
		}
	} // implement binary search for optimisation

	if i != len(likes.Users) {
		post.IsLiked = true
	}

	if post.IsLiked {
		database.DB.Model(&post).Update("likes_count", post.LikesCount-1)
		// database.DB.Model(&post).Update("is_liked", false)
		likes.Users = append(likes.Users[:i], likes.Users[i+1:]...)
		database.DB.Model(&likes).Update("users", likes.Users)

	} else {
		database.DB.Model(&post).Update("likes_count", post.LikesCount+1)
		// database.DB.Model(&post).Update("is_liked", true)

		likes.Users = append(likes.Users, 44) // put likes.Users in sorted order

		if likes.ID != 0 {
			database.DB.Model(&likes).Update("users", likes.Users)

		} else {
			likes.ID = body.PostID
			database.DB.Save(&likes)
		}
	}

	return c.JSON(models.Message{
		Message: "request successful",
	})
}

func CommentPost(c *fiber.Ctx) error {
	body := new(
		struct {
			PostId     uint   `json:"post_id"`
			CommentMsg string `json:"comment_msg"`
		},
	)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: err.Error(),
		})
	}

	post := models.Post{}

	database.DB.First(&post, "id = ?", body.PostId)

	if post.PostID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Message{
			Message: "post doesen't exist",
		})
	}

	// post.Comments = append(post.Comments, body.CommentMsg)

	return c.JSON(models.Message{
		Message: "comment created successfully",
	})
}
