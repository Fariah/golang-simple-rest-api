package main

import (
	restapi "golang-simple-rest-api"
	"golang-simple-rest-api/pkg/handler"
	"golang-simple-rest-api/pkg/repository"
	"golang-simple-rest-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(restapi.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:  %s", err.Error())
	}
}
