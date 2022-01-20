package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/entities", listEntities)
	app.GET("/entities/:id", getEntity)

	app.Logger.Fatal(app.Start(":8080"))
}

func listEntities(ctx echo.Context) error {
	entities := []Entity{
		{"1", PersonEntityType, "John Doe", "John Doe is a person"},
		{"2", CompanyEntityType, "Google", "Google is a company"},
		{"3", PlaceEntityType, "New York", "New York is a place"},
		{"4", BookEntityType, "The Hitchhiker's Guide to the Galaxy", "The Hitchhiker's Guide to the Galaxy is a book"},
		{"5", MovieEntityType, "Star Wars", "Star Wars is a movie"},
		{"6", TvSeriesEntityType, "Game of Thrones", "Game of Thrones is a tv series"},
		{"7", GameEntityType, "Minecraft", "Minecraft is a game"},
		{"8", AlbumEntityType, "The Beatles", "The Beatles is an album"},
		{"9", SongEntityType, "Yesterday", "Yesterday is a song"},
	}

	return ctx.JSON(http.StatusOK, entities)
}

func getEntity(ctx echo.Context) error {
	entityID := ctx.Param("id")

	entity := Entity{
		ID:          entityID,
		Type:        UknownEntityType,
		Name:        "Unknown",
		Description: "Unknown",
	}

	return ctx.JSON(http.StatusOK, entity)
}
