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
	AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Run запускает тг бота
func Run(ctx context.Context, s iService) {
	token := readToken()
	if token == "" {
		log.Println("Empty token!")
		return
	}

	cmdr, err := commander.New(token, s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("bot run")
	if err = cmdr.Run(ctx); err != nil {
		log.Println(err)
	}
}
