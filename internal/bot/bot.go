package bot

import (
	"fmt"
	"log"
	"os"

	"Timeline/internal/listeners"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Initialize() (*discordgo.Session, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ERROR: Can't load .env file: %w", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("ERROR: Missing token")
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Failed to create bot: %w", err)
	}

	listeners.RegisterListeners(bot)

	return bot, nil
}

func OpenBot(bot *discordgo.Session) error {
	err := bot.Open()
	if err != nil {
		return fmt.Errorf("ERROR: Failed to connect: %w", err)
	}

	log.Println("INFO: Connected")
	return nil
}
