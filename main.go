package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type InputContact struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var phoneBook []Contact = make([]Contact, 0)

func process(input InputContact) (string, string) {
	switch input.Action {
	case "add":
		phoneBook = append(phoneBook, Contact{Name: input.Name, Phone: input.Phone})
		return "message", "success added"
	case "clear":
		phoneBook = make([]Contact, 0)
		return "message", "success deleted"
	default:
		return "error", "unknown action"
	}
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		input := InputContact{}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		messageType, message := process(input)
		return c.JSON(fiber.Map{
			messageType: message,
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": phoneBook,
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	app.Listen(":" + port)
}
