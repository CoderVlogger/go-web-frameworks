package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
)

type JSONTextResponse struct {
	Message string
}

var (
	entitiesRepo pkg.EntityRepository = pkg.NewEntityMemoryRepository()
	pageSize     int                  = 4
)

func main() {
	entitiesRepo.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, Fiber!")

		return c.JSON(JSONTextResponse{Message: "Hello, Fiber!"})
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", entitiesList)
	entitiesAPI.Get("/:id", entitiesGet)
	entitiesAPI.Post("/", entitiesAdd)
	entitiesAPI.Put("/", entitiesUpdate)
	entitiesAPI.Delete("/:id", entitiesDelete)

	app.Listen(":8080")
}

func entitiesList(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	entities, err := entitiesRepo.List(page, pageSize)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusNotFound).JSON(errMsg)
	}

	return c.JSON(entities)
}

func entitiesGet(c *fiber.Ctx) error {
	entityID := c.Params("id")

	entity, err := entitiesRepo.Get(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	if entity == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(entity)
}

func entitiesAdd(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = entitiesRepo.Add(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entity)
}

func entitiesUpdate(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = entitiesRepo.Update(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entity)
}

func entitiesDelete(c *fiber.Ctx) error {
	entityID := c.Params("id")

	err := entitiesRepo.Delete(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(pkg.TextResponse{Message: "entity deleted"})
}
