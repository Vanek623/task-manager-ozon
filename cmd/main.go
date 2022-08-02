package main

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/client"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
	taskStoragePkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/storage"
)

func main() {
	tm := taskStoragePkg.NewLocal()

	go server.RunREST()
	go server.RunGRPC(tm)
	go client.Run(1)
	go client.Run(2)

	bot.Run(tm)

}
