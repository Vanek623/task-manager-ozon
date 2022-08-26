package main

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	storage.Run(ctx)
}
