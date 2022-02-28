package main

import (
	"context"
	"demoapp/internal"
	"log"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var cfg internal.Config

	err := envconfig.Process("app", &cfg)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	internal.Run(ctx, &cfg)
}
