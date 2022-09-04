package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/storage"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		log.Info("Shooting down storage server...")
		cancel()
	}()

	srv, err := storage.NewServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	srv.Run()

	log.Info("Storage server shutdown")
}
