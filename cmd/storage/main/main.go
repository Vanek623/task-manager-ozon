package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	storage.Run(os.Getenv("DB_PASSWORD"))
}