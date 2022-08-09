package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
	servicePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	service, err := servicePkg.NewService()
	if err != nil {
		log.Fatal(err)
	}

	go server.RunREST(ctx)
	go server.RunGRPC(service)
	bot.Run(service)
}
