package main

import (
	"net/http"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/labstack/echo/v4"
)

func (wa *WebApplication) initWebsite() {
	wa.echoApp.Static("/static", "assets")
	wa.echoApp.File("/", "assets/pages/index.html") // Method: get
	wa.echoApp.POST("/", wa.formSubmitEntity)
}

func (wa *WebApplication) formSubmitEntity(ctx echo.Context) error {
	entity := pkg.Entity{}

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.TextResponse{Message: err.Error()})
	}

	if err := wa.entityRepository.Add(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, pkg.TextResponse{Message: err.Error()})
	}

	// TODO: What's the convention when we do form submit.
	// return ctx.Redirect(http.StatusTemporaryRedirect, "/website")
	return ctx.File("assets/pages/index.html")
}
