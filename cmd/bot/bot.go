package bot

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

// Run запускает тг бота
func Run(tm task.IManager) {
	token := readToken()
	cmdr, err := commander.New(token, tm)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
