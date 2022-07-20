package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log"
)

type Commander struct {
	bot      *tgbotapi.BotAPI
	commands map[string]Command
}

func Init(token string) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Printf("Authorized on acconut %s", bot.Self.UserName)

	cmdr := &Commander{bot, make(map[string]Command)}
	cmdr.registerCommand(NewStartCommand())
	cmdr.registerCommand(NewAddCommand())
	cmdr.registerCommand(NewListCommand())
	cmdr.registerCommand(NewGetCommand())
	cmdr.registerCommand(NewUpdateCommand())
	cmdr.registerCommand(NewDeleteCommand())

	cmdr.registerCommand(NewHelpCommand(cmdr.commands))

	return cmdr, nil
}

func (cmdr *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
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

func (cmdr *Commander) registerCommand(command Command) {
	cmdr.commands[command.Name] = command
}

var helpCommandName = NewHelpCommand(nil).Name

func (cmdr *Commander) handleMessage(msg *tgbotapi.Message) (string, error) {
	c, ok := cmdr.commands[msg.Command()]
	if !ok {
		return "", errors.Errorf("Command /%s not found", msg.Command())
	}

	var args string
	if c.Name == helpCommandName {
		args = fmt.Sprintf("%s %s", msg.Chat.FirstName, msg.Chat.LastName)
	} else {
		fmt.Println("This is not help")
		args = msg.CommandArguments()
	}

	res, err := c.Execute(args)
	if err != nil {
		return "", err
	}

	return res, nil
}
