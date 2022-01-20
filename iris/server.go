package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	booksAPI := app.Party("/entities")
	{
		booksAPI.Use(iris.Compression)
		booksAPI.Get("/", listEntities)
		booksAPI.Get("/{id}", getEntity)
	}

	app.Logger().Fatal(app.Listen(":8080"))
}

func listEntities(ctx iris.Context) {
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

	ctx.JSON(entities)
}

func getEntity(ctx iris.Context) {
	entityID := ctx.Params().Get("id")

	entity := Entity{
		ID:          entityID,
		Type:        UknownEntityType,
		Name:        "Unknown",
		Description: "Unknown",
	}

	ctx.JSON(entity)
}
