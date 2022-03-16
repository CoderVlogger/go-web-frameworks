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
	storage := pkg.NewEntityMemoryRepository()
	storage.Init()

	eh := &entitiesHTTP{
		config:     cfg,
		app:        app,
		repository: storage,
	}

	entitiesAPI := eh.app.Group("/entities")
	entitiesAPI.Get("/", eh.list)
	entitiesAPI.Get("/:id", eh.get)
	entitiesAPI.Post("/", eh.add)
	entitiesAPI.Put("/:id", eh.update)
	entitiesAPI.Delete("/:id", eh.delete)

	return eh
}

func (eh *entitiesHTTP) list(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	entities, err := eh.repository.List(page, eh.config.GetPageSize())
	if err != nil {
		// c.JSON(pkg.TextResponse{Message: err.Error()})
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())

		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entities)
}

func (eh *entitiesHTTP) get(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	entity, err := eh.repository.Get(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}

		if err == pkg.ErrEntityNotFound {
			return c.Status(fiber.StatusNotFound).JSON(errMsg)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	if entity == nil {
		return fiber.ErrNotFound
	}

	return c.JSON(entity)
}

func (eh *entitiesHTTP) add(c *fiber.Ctx) error {
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

	return c.JSON(entity)
}

func (eh *entitiesHTTP) update(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	if entity.ID != entityID {
		errMsg := pkg.TextResponse{Message: "Entity ID mismatch"}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = eh.repository.Update(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}

		if err == pkg.ErrEntityNotFound {
			return c.Status(fiber.StatusNotFound).JSON(errMsg)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	return c.JSON(entity)
}

func (eh *entitiesHTTP) delete(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	err := eh.repository.Delete(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}

		if err == pkg.ErrEntityNotFound {
			return c.Status(fiber.StatusNotFound).JSON(errMsg)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	return c.JSON(pkg.TextResponse{Message: "entity deleted"})
}
