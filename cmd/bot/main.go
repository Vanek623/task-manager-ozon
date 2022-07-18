package main

import (
	"TaskAlertBot/internal/commander"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func runBot() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if len(token) == 0 {
		fmt.Print("Token missing in .env! Enter token: ")
		fmt.Scan(&token)
	}

	cmdr, err := commander.Init(token)
	if err != nil {
		log.Panic(err)
	}

	err = cmdr.Run()
	if err != nil {
		log.Panic(err)
	}
}

func testFeatures() {
	s, e := commander.AddCommand{}.Execute(`"a"`)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(s)
}

func main() {
	//testFeatures()
	runBot()
}
