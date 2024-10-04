package main

import (
	"log"
	"os"
	"strings"

	"LoggerBot/commands"

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

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		log.Fatal("ERROR: Failed to connect", err)
	}

	log.Println("INFO: Connected as", bot.State.User.Username, "#", bot.State.User.Discriminator)
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if (m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, os.Getenv("PREFIX"))) {
		return
	}

	cmd := m.Content[1:]

	switch {
	case strings.HasPrefix(cmd, "ping"):
		commands.Ping(s, m)
	case strings.HasPrefix(cmd, "echo"):
		commands.Echo(s, m)
	}
}
