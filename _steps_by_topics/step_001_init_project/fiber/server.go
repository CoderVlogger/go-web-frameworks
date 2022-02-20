package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type JSONTextResponse struct {
	Message string
}

func main() {
	fmt.Println("Hello, World!")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, Fiber!")

		return c.JSON(JSONTextResponse{Message: "Hello, Fiber!"})
	})

	app.Listen(":8080")
}
