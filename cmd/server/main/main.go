package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
