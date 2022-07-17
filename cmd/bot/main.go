package main

import (
	"TaskAlertBot/internal/commander"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if len(token) == 0 {
		fmt.Print("Enter token: ")
		fmt.Scan(&token)
	}

	cmd, err := commander.Init(token)
	if err != nil {
		log.Panic(err)
	}
	handler

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	cmdr := commands.NewCommander(bot, productService)

	for update := range updates {
		cmdr.HandleUpdate(update)
	}

	fmt.Println("Run bot")
}
