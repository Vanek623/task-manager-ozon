package commander

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander/command"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Commander структура бота
type Commander struct {
	bot     *tgbotapi.BotAPI
	manager command.Manager
	cs      *counters.Counters
}

const commanderGroupName = "tg_bot"

// New инициализация бота
func New(token string, s iService) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	cmdr := &Commander{bot, command.NewManager(s), counters.New(commanderGroupName)}

	return cmdr, nil
}

const timeOutValue = 60

// Run запуск бота
func (cmdr *Commander) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeOutValue
	updates := cmdr.bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}
			cmdr.cs.Inc(counters.Incoming)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, cmdr.handleMessage(ctx, update.Message))
			log.WithField("bot incoming message", msg.Text)

			_, err := cmdr.bot.Send(msg)
			if err != nil {
				log.Error(err)
				cmdr.cs.Inc(counters.Fail)
			} else {
				cmdr.cs.Inc(counters.Success)
			}
		}
	}()

	<-ctx.Done()
	cmdr.bot.StopReceivingUpdates()
}

var startCommandName = "start"

func (cmdr *Commander) handleMessage(ctx context.Context, msg *tgbotapi.Message) string {
	c := cmdr.manager.GetCommand(msg.Command())
	if c == nil {
		return fmt.Sprintf("command /%s not found", msg.Command())
	}

	var args string
	if c.Name() == startCommandName {
		args = fmt.Sprintf("%s %s", msg.Chat.FirstName, msg.Chat.LastName)
	} else {
		args = msg.CommandArguments()
	}

	return c.Execute(ctx, args)
}
