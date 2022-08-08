package main

import (
	"flag"
	"log"
)

func main() {
	loadToken()
}

func loadToken() string {
	token := flag.String("bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("Token not found")
	}

	return *token
}
