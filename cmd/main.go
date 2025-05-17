package main

import (
	"MalDB/internal/config"
	"MalDB/internal/core/app"
	"MalDB/internal/core/http/server"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()
	application := app.New(ctx, cfg)

	_, err := application.CreateRedis()
	if err != nil {
		log.Fatalf("Redis: %s", err)
	}
	log.Println("Redis successfully connect")

	_, err = application.CreatePgDB()
	if err != nil {
		log.Fatalf("Postgres: %s", err)
	}
	log.Println("Postgres successfully connect")

	defer application.Close()

	server.Run(application)
}
