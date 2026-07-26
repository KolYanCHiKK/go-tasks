package main

import (
	"log"
	"net/http"
	"project/go-tasks/configs"
	"project/go-tasks/internal/verify"
)

func main() {
	log.Println("Server is listening on port: 8080")

	router := http.NewServeMux()
	conf, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	verify.NewHandler(router, &verify.HandlerDeps{Config: conf})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
