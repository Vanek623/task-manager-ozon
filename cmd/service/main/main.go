package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/service"
)

const envPath = "/home/ivan/GolandProjects/TaskBot/bin/.env"

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		log.Info("Shooting down server...")
		cancel()
	}()

	srv, err := service.NewServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	srv.Run()
}
