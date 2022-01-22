package main

import (
	"github.com/CoderVlogger/go-web-frameworks/pkg"

	"github.com/kataras/iris/v12"
)

var (
	entityStorage pkg.EntityRepository = pkg.NewEntityMemoryRepository()
)

func main() {
	app := iris.New()
	entityStorage.Init()

	booksAPI := app.Party("/entities")
	{
		booksAPI.Use(iris.Compression)
		booksAPI.Post("/", addEntity)
		booksAPI.Put("/", updateEntity)
		booksAPI.Get("/", listEntities)
		booksAPI.Get("/{id}", getEntity)
		booksAPI.Delete("/{id}", deleteEntity)
	}

	app.Logger().Fatal(app.Listen(":8080"))
}

func addEntity(ctx iris.Context) {
	entity := pkg.Entity{}
	err := ctx.ReadJSON(&entity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}

	err = entityStorage.Add(&entity)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(entity)
}

func updateEntity(ctx iris.Context) {
	entity := pkg.Entity{}
	err := ctx.ReadJSON(&entity)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}

	err = entityStorage.Update(&entity)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}
}

func getEntity(ctx iris.Context) {
	entityID := ctx.Params().Get("id")

	entity, err := entityStorage.Get(entityID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if entity == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(pkg.ErrorResponse{Message: "entity not found"})
		return
	}

	ctx.JSON(entity)
}

func listEntities(ctx iris.Context) {
	entities, err := entityStorage.List()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(entities)
}

func deleteEntity(ctx iris.Context) {
	entityID := ctx.Params().Get("id")

	err := entityStorage.Delete(entityID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(pkg.ErrorResponse{Message: err.Error()})
		return
	}
}
