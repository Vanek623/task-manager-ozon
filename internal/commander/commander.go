package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log"
)

type CmdHandler func(string) string

type Commander struct {
	bot    *tgbotapi.BotAPI
	router map[string]CmdHandler
}

func Init(token string) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on acconut %s", bot.Self.UserName)

	return &Commander{
		bot:    bot,
		router: make(map[string]CmdHandler),
	}, nil
}

func (c *Commander) AddHandler(cmd string, ch CmdHandler) {
	c.router[cmd] = ch
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if h, ok := c.router[cmd]; ok {
				msg.Text = h(update.Message.CommandArguments())
			} else {
				msg.Text = "Unknown command"
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName,
				update.Message.Text)
			msg.Text = fmt.Sprintf("You send: %v", update.Message.Text)
		}

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}

	return nil
}
