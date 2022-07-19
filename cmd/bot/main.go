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

//func parseTest() {
//	parseFunc := func(s string) {
//		if r, e := commander.ExtractArgs(s); e != nil {
//			fmt.Println(e.Error())
//		} else {
//			fmt.Println(s, ":", r, "(", len(r), ")")
//		}
//	}
//
//	parseFunc("")
//	parseFunc("bob")
//	parseFunc("bob lol")
//	parseFunc(`"bob" lol`)
//	parseFunc(`lol "bob"`)
//	parseFunc(`"bob" "fof" lel`)
//	parseFunc(`bob "lol kek" mob`)
//}

func testFeatures() {

}

func main() {
	//testFeatures()
	runBot()
}
