package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage.Run(ctx)
}
