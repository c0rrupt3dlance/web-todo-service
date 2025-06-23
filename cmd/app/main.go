package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"web-todo-service/internal/app"
	"web-todo-service/internal/handlers"
	"web-todo-service/internal/repository"
	"web-todo-service/internal/services"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env")
		os.Exit(1)
	}
	pool, err := repository.NewPgPool(
		repository.PgConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		},
	)
	if err != nil {
		log.Println("failed to create postgres pool")
		os.Exit(1)
	}
	repo := repository.NewRepository(pool)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := new(app.Server)

	err = server.Run(os.Getenv("SERVER_PORT"), handler.InitRoutes())
	if err != nil {
		log.Println("failed to start server")
		os.Exit(1)
	}
}
