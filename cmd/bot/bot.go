package bot

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander"
)

func readToken() (token string, err error) {
	if token = os.Getenv("BOT_TOKEN"); token != "" {
		return
	}

	fmt.Print("Token not found! Enter token: ")

	_, err = fmt.Scan(&token)

	return
}

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Run запускает тг бота
func Run(ctx context.Context, s iService, cs *counters.Counters) {
	token, err := readToken()
	if err != nil {
		log.WithField("error", err).Error("Empty token!")
		return
	}

	cmdr, err := commander.New(token, s, cs)
	if err != nil {
		log.Println(err)
		return
	}

	log.Info("Bot run")
	cmdr.Run(ctx)
}
