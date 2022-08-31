package main

import (
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	conf := storage.GetConfig()
	storage.Run(&conf)
}
