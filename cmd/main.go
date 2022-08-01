package main

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/client"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
	taskManagerPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/manager"
)

func main() {
	var tm = taskManagerPkg.NewLocal()

	go server.RunREST()
	go server.RunGRPC(tm)
	go client.Run()
	bot.Run(tm)

}
