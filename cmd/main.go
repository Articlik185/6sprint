package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Сервер запущен на порту :8080")

	err := srv.Start()
	if err != nil {
		logger.Fatal(err)
	}
}
