package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func FileUpload(c *fiber.Ctx) error {

	if c.FormValue("key") == os.Getenv("API_KEY"){

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		// Save file to file directory
		err = c.SaveFile(file, fmt.Sprintf("./static/public/%s", file.Filename))
		if err != nil {
			return err
		}
		return c.Status(201).JSON(fiber.Map{"url": "http://localhost:3000/" + file.Filename})
	} else {
		return c.Status(403).SendString("API Key Invalid")
	}

}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}