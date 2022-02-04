package main

import (
	"net/http"
	"strconv"

	"github.com/CoderVlogger/go-web-frameworks/pkg"

	"github.com/labstack/echo/v4"
)

var (
	pageSize                           = 4
	entityStorage pkg.EntityRepository = pkg.NewEntityMemoryRepository()
)

func main() {
	app := echo.New()
	entityStorage.Init()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	app.POST("/entities", addEntity)
	app.PUT("/entities", updateEntity)
	app.GET("/entities", listEntities)
	app.GET("/entities/:id", getEntity)
	app.DELETE("/entities/:id", deleteEntity)

	app.Logger.Fatal(app.Start(":8080"))
}

func addEntity(ctx echo.Context) error {
	entity := pkg.Entity{}

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	if err := entityStorage.Add(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, entity)
}

func updateEntity(ctx echo.Context) error {
	entity := pkg.Entity{}

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	if err := entityStorage.Update(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, entity)
}

func getEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	entity, err := entityStorage.Get(entityID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.TextResponse{Message: err.Error()})
	}

	if entity == nil {
		return ctx.JSON(http.StatusNotFound, pkg.TextResponse{Message: "entity not found"})
	}

	return ctx.JSON(http.StatusOK, entity)
}

func listEntities(ctx echo.Context) error {
	var err error

	page := 1

	pageStr := ctx.QueryParam("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
		}
	}

	entities, err := entityStorage.List(page, pageSize)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, entities)
}

func deleteEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	if err := entityStorage.Delete(entityID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, pkg.TextResponse{Message: "entity deleted"})
}
