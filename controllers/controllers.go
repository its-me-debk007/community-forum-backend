package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	return c.SendStatus(202)
}
