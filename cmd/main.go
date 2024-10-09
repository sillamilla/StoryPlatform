package main

import (
	Controllers "StoryPlatforn_GIN/internal/app/controller"
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/config"
	"StoryPlatforn_GIN/internal/infrastructure/db"
	"StoryPlatforn_GIN/internal/infrastructure/gin"
	"context"
	"log"
)

func main() {
	cfg := config.GetConfig()

	conn, err := db.NewPostgresDB(cfg.Postgres.URI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	repository := repository.New(conn)
	service := service.New(repository)
	controller := Controllers.New(service)

	router := gin.SetupRouter(*controller)

	router.Run(":" + cfg.HTTP.Port)
}
