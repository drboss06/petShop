package main

import (
	"log"
	"petShop"
	"petShop/pkg/handler"
	"petShop/pkg/repository"
	"petShop/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(petShop.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server is failed to start: %v", err.Error())
	}
}
