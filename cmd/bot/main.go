package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/commander"
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

	cmdr, err := commander.Init(token)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
