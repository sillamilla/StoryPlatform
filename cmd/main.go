package main

import (
	Controllers "StoryPlatforn_GIN/internal/app/controller"
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/infrastructure/db"
	"StoryPlatforn_GIN/internal/infrastructure/gin"
	"context"
	"log"
	"os"
)

func main() {
	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		log.Fatal("DATABASE_URI is not set")
	}

	conn, err := db.NewPostgresDB(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	repository := repository.New(conn)
	service := service.New(repository)
	controller := Controllers.New(service)

	router := gin.SetupRouter(*controller)

	router.Run() //8080
}
