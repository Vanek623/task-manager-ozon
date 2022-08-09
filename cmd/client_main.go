package main

import (
	"sync"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/client"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	clientFunc := func(ID uint) {
		defer wg.Done()
		client.Run(ID)
	}

	clientFunc(1)
	clientFunc(2)

	wg.Wait()
}
