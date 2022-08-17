package commander

import (
	"context"
	"fmt"
	"log"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander/command"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (uint64, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// Commander структура бота
type Commander struct {
	bot     *tgbotapi.BotAPI
	manager command.Manager
}

// New инициализация бота
func New(token string, s iService) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Printf("Authorized on acconut %s", bot.Self.UserName)

	cmdr := &Commander{bot, command.NewManager(s)}

	return cmdr, nil
}

const timeOutValue = 60

// Run запуск бота
func (cmdr *Commander) Run(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeOutValue
	updates := cmdr.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, cmdr.handleMessage(ctx, update.Message))

		_, err := cmdr.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}

	return nil
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
