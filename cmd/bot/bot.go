package bot

import (
	"context"
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander"
)

func readToken() string {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		fmt.Print("Token not found! Enter token: ")
		if _, err := fmt.Scan(&token); err != nil {
			log.Fatal(err)
		}
	}

	return token
}

type iService interface {
	AddTask(ctx context.Context, data models.AddTaskData) (uint, error)
	DeleteTask(ctx context.Context, data models.DeleteTaskData) error
	TasksList(ctx context.Context, data models.ListTaskData) ([]models.TaskBrief, error)
	UpdateTask(ctx context.Context, data models.UpdateTaskData) error
	GetTask(ctx context.Context, data models.GetTaskData) (*models.DetailedTask, error)
}

// Run запускает тг бота
func Run(s iService) {
	token := readToken()
	if token == "" {
		log.Fatal("Empty token!")
	}

	cmdr, err := commander.New(token, s)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	log.Println("bot run")
	if err = cmdr.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
