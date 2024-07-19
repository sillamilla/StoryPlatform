package main

import (
	"StoryPlatforn_GIN/internal/app/controller"
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/infrastructure/db"
	"StoryPlatforn_GIN/internal/infrastructure/gin"
	"context"
	"log"
	"os"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := db.NewPostgresDB(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	repository := repository.New(conn)
	service := service.New(repository)
	controller := Controllers.New(service)

	router := gin.SetupRouter(controller)

	router.Run() //8080
}
