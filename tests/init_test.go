//go:build integration
// +build integration

package tests

import (
	clientPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/client"
)

var client *clientPkg.Client

func init() {
	const host = "localhost:8081"

	var err error
	client, err = clientPkg.New(host, 1)
	if err != nil {
		panic(err)
	}
}
