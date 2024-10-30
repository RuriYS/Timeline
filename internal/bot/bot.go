// bot.go
package bot

import (
	"context"
	"fmt"
	"os"

	"Timeline/internal/database"
	"Timeline/internal/listeners"
	"Timeline/internal/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Client struct {
	Session  *discordgo.Session
	Database *database.Database
	Logger   *logger.Logger
}

func (c *Client) GetDatabase() *database.Database {
	return c.Database
}

func Initialize() (*Client, error) {
	l := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	mongoURI := os.Getenv("MONGO_URI")
	l.Debug("Connecting to MongoDB: " + mongoURI)
	// mongoURI := "mongodb://timeline:timeline@localhost:27017"

	db, err := database.Initialize(mongoURI)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Client{
		Session:  session,
		Database: db,
		Logger:   logger.GetLogger(),
	}

	listeners.RegisterListeners(bot, session)

	return bot, nil
}

func (c *Client) Open() error {
	logger.GetLogger().Debug("Connecting to Discord")
	err := c.Session.Open()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	logger.GetLogger().Debug("Closing database")
	ctx := context.Background()
	if err := c.Database.Close(ctx); err != nil {
		c.Logger.Error("Failed to close database: %v", err)
	}
	return c.Session.Close()
}
