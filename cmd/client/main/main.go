package main

import (
	"context"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/client"
)

func main() {
	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	client.RunClients(ctx, 2)
}
