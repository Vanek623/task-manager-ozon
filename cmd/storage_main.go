package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage.RunStorage(ctx, os.Getenv("DB_PASSWORD"))
}
