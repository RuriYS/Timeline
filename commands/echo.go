package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Echo(s *discordgo.Session, m *discordgo.MessageCreate) {
    content := strings.SplitN(m.Content, " ", 2)
    if len(content) < 2 {
        s.ChannelMessageSend(m.ChannelID, "Please provide a message to echo!")
        return
    }

    s.ChannelMessageSend(m.ChannelID, content[1])
}
