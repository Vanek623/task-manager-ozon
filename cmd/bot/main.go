package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")

	fmt.Println("Run bot")
}
