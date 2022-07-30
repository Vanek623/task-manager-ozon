package main

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
)

func main() {
	var tm task.IManager = task.NewLocalManager()

	//go bot.Run(tm)
	go server.RunREST()
	server.RunGRPC(tm)

	//client.Run()
}
