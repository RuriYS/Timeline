// listeners/messageCreate.go
package listeners

import (
	"Timeline/internal/commands"
	"Timeline/internal/logger"
	"Timeline/internal/models"
	"context"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, c Client) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Handle message event
	ctx := context.Background()
	event := &models.MessageEvent{
		MessageID: m.ID,
		GuildID:   m.GuildID,
		ChannelID: m.ChannelID,
		UserID:    m.Author.ID,
		Content:   m.Content,
		Metadata: map[string]interface{}{
			"author": map[string]interface{}{
				"username":      m.Author.Username,
				"discriminator": m.Author.Discriminator,
			},
			"attachments": m.Attachments,
			"embeds":      m.Embeds,
			"mentions":    m.Mentions,
		},
		Timestamp: time.Now(),
	}

	jsonE, _ := json.Marshal(event)
	logger.GetLogger().Debug("messageCreate: " + string(jsonE))

	if err := c.GetDatabase().Messages.InsertOne(ctx, event); err != nil {
		logger.GetLogger().Error("Failed to store message: %v", err)
		return
	}

	// Handle commands
	if !strings.HasPrefix(m.Content, os.Getenv("PREFIX")) {
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
