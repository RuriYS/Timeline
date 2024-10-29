package main

import (
	"log"

	"Timeline/internal/bot"
	"Timeline/internal/logger"
)

func main() {
	l := logger.NewLogger()
	if l == nil {
		log.Fatal("Failed to initialize logger")
	}
	defer l.Close()

	l.SetLogLevel(logger.DEBUG)
	l.SetUseColors(true)

	botSession, err := bot.Initialize()
	if err != nil {
		l.Error("Failed to initialize bot: " + err.Error())
	}

	err = bot.OpenBot(botSession)
	if err != nil {
		l.Error("Failed to create bot session: " + err.Error())
	}

	select {}
}
