package main

import (
	"context"
	"log"
	"os"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/client"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage := server.ConnectToDB(ctx, os.Getenv("DB_PASSWORD"))

	go server.RunREST(ctx)
	go server.RunGRPC(storage)

	go client.Run(1)
	go client.Run(2)

	bot.Run(storage)
}
