package main

import (
	"log"
	"os"

	"LoggerBot/listeners"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("ERROR: Can't load .env file")
    }
	
	token := os.Getenv("DISCORD_TOKEN")
    if token == "" {
        log.Fatalf("ERROR: Missing token")
    }

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("ERROR: Failed to create bot", err)
		return
	}

	listeners.RegisterListeners(bot)

	err = bot.Open()
	if err != nil {
		log.Fatal("ERROR: Failed to connect", err)
	}

	log.Println("INFO: Connected as", bot.State.User.Username, "#", bot.State.User.Discriminator)
	select {}
}
