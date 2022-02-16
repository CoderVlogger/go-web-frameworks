package main

import (
	"net/http"
	"strconv"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/labstack/echo/v4"
)

func (wa *WebApplication) initAPIRoutes() {
	wa.echoApp.POST(apiPrefix+"entities", wa.addEntity)
	wa.echoApp.PUT(apiPrefix+"entities", wa.updateEntity)
	wa.echoApp.GET(apiPrefix+"entities", wa.listEntities)
	wa.echoApp.GET(apiPrefix+"entities/:id", wa.getEntity)
	wa.echoApp.DELETE(apiPrefix+"entities/:id", wa.deleteEntity)
}

func (wa *WebApplication) addEntity(ctx echo.Context) error {
	entity := pkg.Entity{}

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	if err := wa.entityRepository.Add(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, entity)
}

func (wa *WebApplication) updateEntity(ctx echo.Context) error {
	entity := pkg.Entity{}

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	if err := wa.entityRepository.Update(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (wa *WebApplication) getEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	entity, err := wa.entityRepository.Get(entityID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.TextResponse{Message: err.Error()})
	}

	if entity == nil {
		return ctx.JSON(http.StatusNotFound, pkg.TextResponse{Message: "entity not found"})
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (wa *WebApplication) listEntities(ctx echo.Context) error {
	var err error

	page := 1

	pageStr := ctx.QueryParam("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
		}
	}

	entities, err := wa.entityRepository.List(page, pageSize)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, entities)
}

func (wa *WebApplication) deleteEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	if err := wa.entityRepository.Delete(entityID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, pkg.TextResponse{Message: "entity deleted"})
}
