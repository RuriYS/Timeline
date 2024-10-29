// main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Timeline/internal/bot"
	"Timeline/internal/logger"
)

func main() {
	l := logger.NewLogger(logger.DEBUG, true)
	if l == nil {
		log.Fatal("Failed to initialize logger")
	}
	defer l.Close()

	botInstance, err := bot.Initialize()
	if err != nil {
		l.Error("Failed to initialize bot: %v", err)
		return
	}
	defer botInstance.Close()

	err = botInstance.Open()
	if err != nil {
		l.Error("Failed to open bot connection: %v", err)
		return
	}

	// Wait for interrupt signal to gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	l.Info("Bot is now running.")
	<-stop
}
