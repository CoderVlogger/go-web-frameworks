package internal

import (
	"context"
	entitiesHTTP "demoapp/internal/pkg/http/entities"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(ctx context.Context, cfg *Config) {
	app := fiber.New()

	_ = entitiesHTTP.New(cfg, app)

	go func(ctx context.Context) {
		<-ctx.Done()
		err := app.Shutdown()
		if err != nil {
			log.Printf("%v", err)
		}
	}(ctx)

	log.Println(app.Listen(fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)))
}
