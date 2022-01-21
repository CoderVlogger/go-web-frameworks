package main

import (
	"net/http"

	"github.com/CoderVlogger/go-web-frameworks/pkg"

	"github.com/labstack/echo/v4"
)

var (
	entityStorage pkg.EntityRepository = pkg.NewEntityMemoryRepository()
)

func main() {
	app := echo.New()
	entityStorage.Init()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/entities", listEntities)
	app.GET("/entities/:id", getEntity)

	app.Logger.Fatal(app.Start(":8080"))
}

func listEntities(ctx echo.Context) error {
	entities, err := entityStorage.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, entities)
}

func getEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	entity, err := entityStorage.Get(entityID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: err.Error()})
	}

	if entity == nil {
		return ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "entity not found"})
	}

	return ctx.JSON(http.StatusOK, entity)
}
