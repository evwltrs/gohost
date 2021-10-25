package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func FileUpload(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := c.Request().Header.Peek("api_key")
	if string(key) == os.Getenv("API_KEY") {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		// Save file to file directory
		err = c.SaveFile(file, fmt.Sprintf("./static/public/%s", file.Filename))
		if err != nil {
			return err
		}
		return c.Status(201).JSON(fiber.Map{"url": os.Getenv("HOSTNAME") + "/" + file.Filename})
	} else {
		err := c.SendStatus(403)
		if err != nil {
			return err
		}
	}
	return nil
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
