package main

import (
	"flag"
	"log"

	"github.com/dcwk/linksaver/src"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	src.New(tgBotHost, loadToken())
}

func loadToken() string {
	token := flag.String("bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("Token not found")
	}

	return *token
}
