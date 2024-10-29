package listeners

import (
	"Timeline/internal/commands"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, os.Getenv("PREFIX")) {
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
