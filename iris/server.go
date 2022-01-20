package main

import "github.com/kataras/iris/v12"

type ServerError struct {
	Message string `json:"message"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
}

func main() {
	app := iris.New()

	booksAPI := app.Party("/users")
	{
		booksAPI.Use(iris.Compression)
		booksAPI.Get("/", list)
		booksAPI.Get("/{username}", get)
	}

	app.Logger().Fatal(app.Listen(":8080"))
}

func list(ctx iris.Context) {
	ctx.JSON([]User{
		{
			ID:       1,
			Username: "kataras",
			Bio:      "Go is a concurrent programming language created at Google.",
		},
		{
			ID:       2,
			Username: "iris",
			Bio:      "Iris is a powerful tool for rapid and painless web development.",
		},
	})
}

func get(ctx iris.Context) {
	username := ctx.Params().Get("username")
	user := User{
		ID:       1,
		Username: username,
		Bio:      "Go is a concurrent programming language created at Google.",
	}

	ctx.JSON(user)
}
