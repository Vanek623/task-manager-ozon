package commander

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// Commander структура бота
type Commander struct {
	bot      *tgbotapi.BotAPI
	commands map[string]command
}

// Init инициализация бота
func Init(token string) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Printf("Authorized on acconut %s", bot.Self.UserName)

	cmdr := &Commander{bot, make(map[string]command)}
	cmdr.registerCommand(newStartCommand())
	cmdr.registerCommand(newAddCommand())
	cmdr.registerCommand(newListCommand())
	cmdr.registerCommand(newGetCommand())
	cmdr.registerCommand(newUpdateCommand())
	cmdr.registerCommand(newDeleteCommand())

	cmdr.registerCommand(newHelpCommand(cmdr.commands))

	return cmdr, nil
}

const timeOutValue = 60

// Run запуск бота
func (cmdr *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeOutValue
	updates := cmdr.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if res, err := cmdr.handleMessage(update.Message); err != nil {
			msg.Text = err.Error()
		} else {
			msg.Text = res
		}

		_, err := cmdr.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}

	return nil
}

func (cmdr *Commander) registerCommand(c command) {
	cmdr.commands[c.name] = c
}

var startCommandName = newStartCommand().name

func (cmdr *Commander) handleMessage(msg *tgbotapi.Message) (string, error) {
	c, ok := cmdr.commands[msg.Command()]
	if !ok {
		return "", errors.Errorf("command /%s not found", msg.Command())
	}

	var args string
	if c.name == startCommandName {
		args = fmt.Sprintf("%s %s", msg.Chat.FirstName, msg.Chat.LastName)
	} else {
		args = msg.CommandArguments()
	}

	res, err := c.Execute(args)
	if err != nil {
		return "", err
	}

	return res, nil
}
