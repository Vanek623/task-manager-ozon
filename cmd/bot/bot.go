package bot

import (
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/commander"

	"github.com/joho/godotenv"
)

func readToken() string {
	var token string
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	} else {
		token = os.Getenv("BOT_TOKEN")
	}

	if token == "" {
		fmt.Print("Token not found! Enter token: ")
		if _, err := fmt.Scan(&token); err != nil {
			log.Fatal(err)
		}
	}

	return token
}

type iTaskManager interface {
	Add(t models.Task) error
	Delete(ID uint) error
	List() ([]models.Task, error)
	Update(t models.Task) error
	Get(ID uint) (models.Task, error)
}

// Run запускает тг бота
func Run(tm iTaskManager) {
	token := readToken()
	if token == "" {
		log.Fatal("Empty token!")
	}

	cmdr, err := commander.New(token, tm)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmdr.Run(); err != nil {
		log.Fatal(err)
	}
}
