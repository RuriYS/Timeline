package bot

import (
	"fmt"
	"os"

	"Timeline/internal/listeners"
	"Timeline/internal/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Initialize() (*discordgo.Session, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	listeners.RegisterListeners(bot)

	return bot, nil
}

func OpenBot(bot *discordgo.Session) error {
	l := logger.GetLogger()
	err := bot.Open()
	if err != nil {
		return err
	}

	l.Debug("Connected")
	return nil
}
