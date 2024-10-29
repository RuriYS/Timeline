package listeners

import (
	"Timeline/internal/commands"
	"Timeline/internal/database"
	"Timeline/internal/logger"
	"context"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, c Client) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, os.Getenv("PREFIX")) {
		return
	}

	ctx := context.Background()
	event := &database.Event{
		GuildID:   m.GuildID,
		ChannelID: m.ChannelID,
		UserID:    m.Author.ID,
		EventType: "message_create",
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

	if err := c.GetDatabase().StoreEvent(ctx, event); err != nil {
		logger.GetLogger().Error("Failed to store message: %v", err)
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
