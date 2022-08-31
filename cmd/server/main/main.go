package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/tracer"
	"go.opentelemetry.io/otel"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
)

const envPath = "/home/ivan/GolandProjects/TaskBot/bin/.env"

func main() {
	log.SetLevel(log.InfoLevel)
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

	newCtx, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("main"))
	defer span.End()

	if err := server.Run(newCtx); err != nil {
		cancel()
		log.Error(err)
	}
}
