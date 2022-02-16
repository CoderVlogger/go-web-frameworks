package main

import (
	"strconv"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/gofiber/fiber/v2"
)

func entitiesList(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	entities, err := entitiesRepo.List(page, pageSize)
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

func entitiesGet(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	entity, err := entitiesRepo.Get(entityID)
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

func entitiesAdd(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	err = entitiesRepo.Add(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(entity)
	return nil
}

func entitiesUpdate(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	err = entitiesRepo.Update(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(entity)
	return nil
}

func entitiesDelete(c *fiber.Ctx) error {
	entityID := c.Params("id", "")

	err := entitiesRepo.Delete(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	c.JSON(pkg.TextResponse{Message: "entity deleted"})
	return nil
}
