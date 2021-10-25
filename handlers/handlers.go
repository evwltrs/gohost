package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func FileUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Save file to file directory
	return c.SaveFile(file, fmt.Sprintf("./static/public/%s", file.Filename))
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}