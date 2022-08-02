package bot

import (
	"context"
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

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) error
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, t models.Task) error
	Get(ctx context.Context, ID uint) (*models.Task, error)
}

// Run запускает тг бота
func Run(tm iTaskStorage) {
	token := readToken()
	if token == "" {
		log.Fatal("Empty token!")
	}

	cmdr, err := commander.New(token, tm)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cl := context.WithCancel(context.Background())
	defer cl()

	if err = cmdr.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
