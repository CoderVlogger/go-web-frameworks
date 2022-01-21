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
		booksAPI.Get("/", listEntities)
		booksAPI.Get("/{id}", getEntity)
	}

	app.Logger().Fatal(app.Listen(":8080"))
}

func addEntity(ctx iris.Context) {
}

func updateEntity(ctx iris.Context) {
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
}
