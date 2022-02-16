package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
)

var (
	entitiesRepo pkg.EntityRepository = pkg.NewEntityMemoryRepository()
	pageSize                          = 4
)

func main() {
	app := fiber.New()
	entitiesRepo.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", entitiesList)
	entitiesAPI.Get("/:id", entitiesGet)
	entitiesAPI.Post("/", entitiesAdd)
	entitiesAPI.Put("/", entitiesUpdate)
	entitiesAPI.Delete("/:id", entitiesDelete)

	app.Listen(":8080")
}
