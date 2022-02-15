package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
)

var (
	pageSize = 4
)

func main() {
	app := fiber.New()

	entitiesRepo := pkg.NewEntityMemoryRepository()
	entitiesRepo.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", func(c *fiber.Ctx) error {
		pageStr := c.Query("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		entities, err := entitiesRepo.List(page, pageSize)
		if err != nil {
			// c.JSON(pkg.TextResponse{Message: err.Error()})
			// return fiber.NewError(fiber.StatusInternalServerError, err.Error())

			errMsg := pkg.TextResponse{Message: err.Error()}
			return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
		}

		c.JSON(entities)
		return nil
	})
	entitiesAPI.Get("/:id", func(c *fiber.Ctx) error {
		entityID := c.Params("id", "")

		entity, err := entitiesRepo.Get(entityID)
		if err != nil {
			c.JSON(pkg.TextResponse{Message: err.Error()})
			return fiber.ErrInternalServerError
		}

		if entity == nil {
			return fiber.ErrNotFound
		}

		c.JSON(entity)
		return nil
	})

	app.Listen(":8080")
}
