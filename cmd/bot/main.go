package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"

	"github.com/joho/godotenv"
)

func main() {
	var token string
	if err := godotenv.Load(); err != nil {
		fmt.Print("Token missing in .env! Enter token: ")
		if _, err = fmt.Scan(&token); err != nil {
			log.Fatal(err)
		}
	} else {
		token = os.Getenv("BOT_TOKEN")
	}

	var manager task.IManager = task.NewLocalManager()
	cmdr, err := commander.New(token, manager)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
