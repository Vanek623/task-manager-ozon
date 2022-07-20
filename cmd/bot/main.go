package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/commander"
	"log"
	"os"
)

func main() {
	var token string
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Token missing in .env! Enter token: ")
		if _, err := fmt.Scan(&token); err != nil {
			log.Fatal(err)
		}
	} else {
		token = os.Getenv("BOT_TOKEN")
	}

	cmdr, err := commander.Init(token)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
