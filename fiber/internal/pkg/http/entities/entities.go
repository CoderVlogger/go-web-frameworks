package entities

import (
	"strconv"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/gofiber/fiber/v2"
)

type Config interface {
	GetPageSize() int
}

type entitiesHTTP struct {
	config     Config
	app        *fiber.App
	repository pkg.EntityRepository // TODO: Change to app.service.
}

func New(cfg Config, app *fiber.App) *entitiesHTTP {
	eh := &entitiesHTTP{
		config:     cfg,
		app:        app,
		repository: pkg.NewEntityMemoryRepository(),
	}

	entitiesAPI := eh.app.Group("/entities")
	entitiesAPI.Get("/", eh.List)
	entitiesAPI.Get("/:id", eh.Get)
	entitiesAPI.Post("/", eh.Add)
	entitiesAPI.Put("/", eh.Update)
	entitiesAPI.Delete("/:id", eh.Delete)

	return eh
}

func (eh *entitiesHTTP) List(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	entities, err := eh.repository.List(page, eh.config.GetPageSize())
	if err != nil {
		// c.JSON(pkg.TextResponse{Message: err.Error()})
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())

		// TODO: Replace 500 error with 4xx for paging related errors.
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(entities)
	return nil
}

func (eh *entitiesHTTP) Get(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	entity, err := eh.repository.Get(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	if entity == nil {
		return fiber.ErrNotFound
	}

	c.JSON(entity)
	return nil
}

func (eh *entitiesHTTP) Add(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	err = eh.repository.Add(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(entity)
	return nil
}

func (eh *entitiesHTTP) Update(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = eh.repository.Update(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(entity)
	return nil
}

func (eh *entitiesHTTP) Delete(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	err := eh.repository.Delete(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(pkg.TextResponse{Message: "entity deleted"})
	return nil
}
