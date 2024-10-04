package main

import (
	"log"

	"EventHorizon/internal/bot"
)

func main() {
	botSession, err := bot.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	err = bot.OpenBot(botSession)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
