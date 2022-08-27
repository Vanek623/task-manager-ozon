package main

import (
	"context"
	"os"

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

	ctx := context.Background()

	newCtx, span := otel.Tracer(tracer.Name).Start(ctx, tracer.MakeSpanName("main"))
	defer span.End()

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(newCtx); err != nil {
		log.Fatal(err)
	}
}
