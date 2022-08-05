package main

import (
	"context"
	servicePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"
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

	pool := server.ConnectToDB(ctx, os.Getenv("DB_PASSWORD"))
	service := servicePkg.NewService(pool)

	go server.RunREST(ctx)
	go server.RunGRPC(service)

	go client.Run(1)
	go client.Run(2)

	bot.Run(service)
}
