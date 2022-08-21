//go:build integration
// +build integration

package tests

import (
	"os"

	"github.com/joho/godotenv"
	clientPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/client"
)

var client *clientPkg.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client, err = clientPkg.New(os.Getenv("SERVICE_HOST"), 1)
	if err != nil {
		panic(err)
	}
}
