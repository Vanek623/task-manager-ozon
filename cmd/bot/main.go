package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"

	"github.com/joho/godotenv"
)

func readToken() string {
	var token string
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file didn't open!")
	} else {
		token = os.Getenv("BOT_TOKEN")
	}

	if token == "" {
		fmt.Print("Token not found! Enter token: ")
		if _, err := fmt.Scan(&token); err != nil {
			log.Fatal(err)
		}
	}

	if token == "" {
		log.Fatal("Empty token!")
	}

	return token
}

func main() {
	token := readToken()
	var manager task.IManager = task.NewLocalManager()
	cmdr, err := commander.New(token, manager)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
