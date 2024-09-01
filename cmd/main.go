package main

import (
	"log"
	todo "service_todo"
	"service_todo/pkg/hanlders"
)

func main() {
	handlers := new(hanlders.Handler)
	server := new(todo.Server)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
