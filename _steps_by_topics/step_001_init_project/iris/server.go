package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	fmt.Println("Hello, Iris!")

	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	})

	app.Logger().Fatal(app.Listen(":8080"))
}
